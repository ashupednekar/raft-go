package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/ashupednekar/raft-go/internal/server/raft"
	"github.com/ashupednekar/raft-go/internal/state"
	"google.golang.org/grpc"
)

type Server struct {
  Id int 
  State state.State
  LastHeartBeat time.Time
  pb.UnimplementedRaftServiceServer
}

func (s *Server) AppendEntries(ctx context.Context, in *pb.EntryInput) (*pb.EntryResult, error){
  s.LastHeartBeat = time.Now()
  current_term := int32(s.State.PersistentState.CurrentTerm)
  if in.Term < current_term{
    return &pb.EntryResult{Term: current_term, Success: false}, nil
  }else{
    
  }
  return &pb.EntryResult{}, nil
}

func (s *Server) RequestVote(ctx context.Context, in *pb.VoteInput) (*pb.VoteResult, error){
  current_term := int32(s.State.PersistentState.CurrentTerm)
  if in.Term < current_term{
    return &pb.VoteResult{Term: current_term, VoteGranted: false}, nil
  }else{
    if s.State.PersistentState.VotedFor == 0 || int32(s.State.PersistentState.VotedFor) == in.CandidateId{
      if in.LastLogIndex >= int32(s.State.CommitIndex){
        return &pb.VoteResult{Term: current_term, VoteGranted: true}, nil
      }
    }
    return &pb.VoteResult{}, nil
  }
} 

func (s *Server) Start (name string, port int){
  ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
  if err != nil{
    log.Fatalf("failed to listen at port 8001: %v", err)
  }

  grpcServer := grpc.NewServer()
  pb.RegisterRaftServiceServer(grpcServer, s)

  log.Printf("gRPC server %s listening at %v", name, ln.Addr())
  if err := grpcServer.Serve(ln); err != nil{
    log.Fatalf("failed to start gRPC server: %v", err)
  }
}
