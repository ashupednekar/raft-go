package server

import (
	"context"
	"fmt"
	"time"

	pb "github.com/ashupednekar/raft-go/internal/server/raft"
	"github.com/ashupednekar/raft-go/internal/state"
)

func (s *Server) AppendEntries(ctx context.Context, in *pb.EntryInput) (*pb.EntryResult, error){
  s.LastHeartBeat = time.Now()
  fmt.Printf("received heartbeat from %d\n", in.LeaderId)
  if s.State.Role == state.Leader && in.Term >= int32(s.State.PersistentState.CurrentTerm){
    fmt.Printf("server %d stepping down as leader\n", s.State.Id)
    s.QuitLeadingChan <- true
    s.State.Role = state.Follower
  }
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
