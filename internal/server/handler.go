package server

import (
	"context"
	"fmt"
	"time"

	pb "github.com/ashupednekar/raft-go/internal/server/raft"
)

func (s *Server) AppendEntries(ctx context.Context, in *pb.EntryInput) (*pb.EntryResult, error){
  s.LastHeartBeat = time.Now()
  fmt.Printf("received heartbeat from leader: %d at %v\n", in.LeaderId, s.LastHeartBeat)
  current_term := int32(s.State.PersistentState.CurrentTerm)
  if in.Term < current_term{
    return &pb.EntryResult{Term: current_term, Success: false}, nil
  }
  return &pb.EntryResult{}, nil
}

func (s *Server) RequestVote(ctx context.Context, in *pb.VoteInput) (*pb.VoteResult, error){
  current_term := int32(s.State.PersistentState.CurrentTerm)
  if in.Term < current_term{
    return &pb.VoteResult{Term: current_term, VoteGranted: false}, nil
  }else{
    fmt.Printf("%d has voted for %d\n", s.State.Id, s.State.PersistentState.VotedFor)
    if s.State.PersistentState.VotedFor == 0 || int32(s.State.PersistentState.VotedFor) == in.CandidateId{
      if in.LastLogIndex >= int32(s.State.CommitIndex){
        s.State.PersistentState.VotedFor = int(in.CandidateId)
        s.State.SavePersistentState()
        return &pb.VoteResult{Term: current_term, VoteGranted: true}, nil
      }
    }
    return &pb.VoteResult{}, nil
  }
} 


