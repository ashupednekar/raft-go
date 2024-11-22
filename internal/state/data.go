package state

type Role int

const (
  Follower Role = iota
  Candidate
  Leader
)

func (w Role) String() string {
	return [...]string{"Follower", "Candidate", "Leader"}[w]
}

type PersistentState struct{
  CurrentTerm int `json:"current_term"`
  VotedFor int `json:"voted_for"`
  LeaderId int `json:"leader_id"`
}

type LeaderState struct{
  nextIndex map[string]int
  matchIndex map[string]int
}

type LogEntry struct{
  Term int
  Index int
  Command string
}

type State struct{
  Id int 
  Port int
  CommitIndex int
  LastAppliedIndex int
  PersistentState PersistentState
  Role Role
  Log []LogEntry
}


