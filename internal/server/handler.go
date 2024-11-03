package server

import (
	"context"
	"time"

	pb "github.com/ashupednekar/raft-go/internal/server/raft"
	"github.com/ashupednekar/raft-go/internal/state"
)

func (s *Server) AppendEntries(ctx context.Context, in *pb.EntryInput) (*pb.EntryResult, error){
  s.LastHeartBeat = time.Now()
  //fmt.Printf("received heartbeat from %d\n", in.LeaderId)
  current_term := int32(s.State.PersistentState.CurrentTerm)

  if in.Term >= current_term{
    s.State.Role = state.Follower
    s.QuitLeadingChan <- true
  }else{
    return &pb.EntryResult{Term: current_term, Success: false}, nil
  }
  return &pb.EntryResult{Term: current_term, Success: false}, nil
}


func (s *Server) RequestVote(ctx context.Context, in *pb.VoteInput) (*pb.VoteResult, error) {
    currentTerm := int32(s.State.PersistentState.CurrentTerm)
    reject :=  &pb.VoteResult{Term: currentTerm, VoteGranted: false}
    vote :=  &pb.VoteResult{Term: currentTerm, VoteGranted: true}
    if in.Term < currentTerm {
        return reject, nil 
    } else {
      s.State.PersistentState.CurrentTerm = int(in.Term)
      s.State.SavePersistentState()
      logUpToDate := in.LastLogIndex >= int32(s.State.CommitIndex)
      alreadyVotedOther := s.State.PersistentState.VotedFor != 0 && int32(s.State.PersistentState.VotedFor) != in.CandidateId 
      if in.Term == currentTerm{
        if alreadyVotedOther{
          return reject, nil
        } else {
          if logUpToDate{
            s.State.PersistentState.VotedFor = int(in.CandidateId)
            s.State.PersistentState.CurrentTerm = int(in.Term)
            s.State.SavePersistentState()
            return vote, nil
          } else {
            return reject, nil
          }
        }
      } else {
        if logUpToDate{
          s.State.PersistentState.VotedFor = int(in.CandidateId)
          s.State.SavePersistentState()
          return vote, nil
        } else {
          return reject, nil
        }
      }
    }
}
