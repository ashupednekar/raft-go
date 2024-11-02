package internal

import "github.com/ashupednekar/raft-go/internal/server"

func InitiateElection(s *server.Server){
  s.State.PersistentState.CurrentTerm = s.State.PersistentState.CurrentTerm + 1
  s.State.SavePersistentState()
    
}
