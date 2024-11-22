package server 

import (
	"context"
	"fmt"
	"time"

	pb "github.com/ashupednekar/raft-go/internal/server/raft"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(addr string) (pb.RaftServiceClient, error) {
  conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil{
    fmt.Printf("error connecting to server at: %s - %v\n", addr, err)
  }
  return pb.NewRaftServiceClient(conn), nil
}

func RequestVote(client pb.RaftServiceClient, s *Server) (int, bool, error){
  ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 100)
  defer cancel()
  r, err := client.RequestVote(ctx, &pb.VoteInput{
    Term: int32(s.State.PersistentState.CurrentTerm),
    CandidateId: int32(s.State.Id), // TODO: fill rest of the fields
  })
  if err != nil{
    fmt.Printf("requestVote rpc call failed: %v\n", err)
    return 0, false, err
  }
  return int(r.Term), r.VoteGranted, nil
}

func AppendEntries(client pb.RaftServiceClient, s *Server) (int, bool, error) {
  ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 100)
  defer cancel()
  entries := make([]*pb.Entry, len(s.State.Log))
  if len(s.State.Log) != 0{
    entries = append(entries, &pb.Entry{Name: s.State.Log[s.State.CommitIndex]})
  }

  prevLogIndex := len(s.State.Log) - 1
  prevLogTerm := 0
  if len(s.State.Log) != 0{
    prevLogTerm = s.State.Log[prevLogIndex].Term
  }

  r, err := client.AppendEntries(ctx, &pb.EntryInput{
    Term: int32(s.State.PersistentState.CurrentTerm),
    LeaderId: int32(s.State.Id),
    Entries: entries,
    PrevLogIndex: int32(prevLogIndex),
    PrevLogTerm: int32(prevLogTerm),

    // TODO: fill rest of the fields
  })
  if err != nil{
    fmt.Printf("appendEntries rpc call failed: %v\n", err)
    return 0, false, err
  }
  return int(r.Term), r.Success, nil
}
