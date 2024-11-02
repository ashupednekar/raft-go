package state

type PersistentState struct{
  CurrentTerm int `json:"current_term"`
  VotedFor int `json:"voted_for"`
}

type LeaderState struct{
  nextIndex map[string]int
  matchIndex map[string]int
}

type State struct{
  Id int 
  CommitIndex int
  LastAppliedIndex int
  PersistentState PersistentState
  Log []string
}


