package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/ashupednekar/raft-go/internal/server/raft"
	"github.com/ashupednekar/raft-go/internal/state"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *Server) AppendEntries(ctx context.Context, in *pb.EntryInput) (*pb.EntryResult, error){
  s.LastHeartBeat = time.Now()
  if s.State.Role == state.Leader && in.Term >= int32(s.State.PersistentState.CurrentTerm){
    fmt.Printf("server %d stepping down as leader\n", s.State.Id)
    s.QuitLeadingChan <- true
  }
  s.State.Role = state.Follower
  s.State.PersistentState.LeaderId = int(in.LeaderId) 
  s.State.SavePersistentState()
  return &pb.EntryResult{Term: int32(s.State.PersistentState.CurrentTerm), Success: false}, nil
}

func (s *Server) RequestVote(ctx context.Context, in *pb.VoteInput) (*pb.VoteResult, error){
  if in.Term < int32(s.State.PersistentState.CurrentTerm){
    return &pb.VoteResult{Term: int32(s.State.PersistentState.CurrentTerm), VoteGranted: false}, nil
  }else{
    //same term
    if in.Term == int32(s.State.PersistentState.CurrentTerm){
      if s.State.PersistentState.VotedFor != int(in.CandidateId) && s.State.PersistentState.VotedFor != 0{
        // TODO: add log check
        s.State.PersistentState.CurrentTerm = int(in.Term)
        s.State.PersistentState.VotedFor = int(in.CandidateId)
        s.State.SavePersistentState()
        return &pb.VoteResult{Term: int32(s.State.PersistentState.CurrentTerm), VoteGranted: true}, nil
      }else{
        return &pb.VoteResult{Term: int32(s.State.PersistentState.CurrentTerm), VoteGranted: false}, nil
      }
    }else{
      //new term
      s.State.PersistentState.VotedFor = int(in.CandidateId)
      s.State.PersistentState.CurrentTerm = int(in.Term)
      s.State.SavePersistentState()
      //TODO: add log check
      return &pb.VoteResult{Term: in.Term, VoteGranted: true}, nil
    }
  }
}

func (s *Server) Write(ctx context.Context, in *pb.File) (*pb.WriteResult, error){
  if s.State.Role == state.Leader{
    log.Printf("leader: server%d received request, processing", s.State.Id)
    s.State.Log = append(s.State.Log, "write~~%s~~%s", in.Name, in.Content)
    s.State.CommitIndex++;
    results := SpawnAppends(s)
    server_count, err := strconv.Atoi(os.Getenv("SERVER_COUNT"))
    if err != nil{
      fmt.Println("error reading server count env")
    }
    majority := server_count / 2 + 1
    count := 0
    for res := range(results){
      if res.Success{
        count++
      }
    }
    if (count >= majority){
      fmt.Println("quorum achieved, proceeding to commit")
      file, err := os.Create(in.Name)
      if err != nil{
        fmt.Printf("error creating file: %v\n", err)
        return &pb.WriteResult{}, err
      }
      _, err = file.WriteString(in.Content)
      if err != nil{
        fmt.Printf("error writing to file: %v\n", err)
      }
      return &pb.WriteResult{Ok: true}, nil
    } else{
      fmt.Printf("quorum not reached, no. of successful appends: %d\n", count) 
      return &pb.WriteResult{}, fmt.Errorf("no quorum") 
    }
  }else{
    log.Printf("received, request... server%d not a leader, redirecting to server %d", s.State.Id, s.State.PersistentState.LeaderId)
    leaderAddr := fmt.Sprintf("server%d:800%d", s.State.PersistentState.LeaderId, s.State.PersistentState.LeaderId)
    conn, err := grpc.Dial(leaderAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil{
      return &pb.WriteResult{}, err
    }
    client := pb.NewFileServiceClient(conn)
    return client.Write(ctx, in)
  }
}

func (s *Server) Read(ctx context.Context, in *pb.Name) (*pb.ReadResult, error){
  return &pb.ReadResult{Content: ""}, nil
}











