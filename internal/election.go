package internal

import "github.com/ashupednekar/raft-go/internal/server"

func InitiateElection(s *server.Server){
  s.State.PersistentState.CurrentTerm = s.State.PersistentState.CurrentTerm + 1
  s.State.PersistentState.VotedFor = s.Id
  s.State.SavePersistentState()
   
}
