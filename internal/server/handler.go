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
  current_term := int32(s.State.PersistentState.CurrentTerm)

  if in.Term >= current_term{
    s.State.Role = state.Follower
    go func(){
      s.QuitLeadingChan <- true
    }()
  }else{
    return &pb.EntryResult{Term: current_term, Success: false}, nil
  }
  return &pb.EntryResult{Term: current_term, Success: false}, nil
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
      s.State.PersistentState.CurrentTerm = int(in.Term)
      s.State.Role = state.Follower
      go func(){
        s.QuitLeadingChan <- true
      }()
      s.State.PersistentState.VotedFor = 0
      //TODO: log check
      return &pb.VoteResult{Term: int32(s.State.PersistentState.CurrentTerm), VoteGranted: true}, nil
    }
  }
} 
