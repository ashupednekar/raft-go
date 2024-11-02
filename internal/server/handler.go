package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ashupednekar/raft-go/internal"
	pb "github.com/ashupednekar/raft-go/internal/server/raft"
	"google.golang.org/grpc"
)

type server struct {
  name string
  state internal.State
  pb.UnimplementedRaftServiceServer
}

func (s *server) AppendEntries(ctx context.Context, in *pb.EntryInput) (*pb.EntryResult, error){
  return &pb.EntryResult{}, nil
}

func (s *server) RequestVote(ctx context.Context, in *pb.VoteInput) (*pb.VoteResult, error){
  current_term := int32(s.state.PersistentState.CurrentTerm)
  if in.Term < current_term{
    return &pb.VoteResult{Term: current_term, VoteGranted: false}, nil
  }else{
    if s.state.PersistentState.VotedFor == "" || s.state.PersistentState.VotedFor == in.CandidateId{
      if in.LastLogIndex >= int32(s.state.CommitIndex){
        return &pb.VoteResult{Term: current_term, VoteGranted: true}, nil
      }
    }
    return &pb.VoteResult{}, nil
  }
} 

func StartServer(name string, port int){
  ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
  if err != nil{
    log.Fatalf("failed to listen at port 8001: %v", err)
  }

  s := grpc.NewServer()
  pb.RegisterRaftServiceServer(s, &server{name: name})

  log.Printf("gRPC server %s listening at %v", name, ln.Addr())
  if err := s.Serve(ln); err != nil{
    log.Fatalf("failed to start gRPC server: %v", err)
  }
}
