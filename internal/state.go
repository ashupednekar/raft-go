package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type PersistentState struct{
  CurrentTerm int `json:"current_term"`
  VotedFor string `json:"voted_for"`
}

type LeaderState struct{
  nextIndex map[string]int
  matchIndex map[string]int
}

type State struct{
  Name string
  CommitIndex int
  LastAppliedIndex int
  PersistentState PersistentState
  Log []string
}


func (s *State) AppendLog(entry string) error {
  s.Log = append(s.Log, entry)
  file, err := os.OpenFile(fmt.Sprintf("%s.log", s.name), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil{
    fmt.Printf("error opening log file: %v\n", err)
    return err
  }
  defer file.Close()

  _, err = file.WriteString(fmt.Sprintf("%s\n", entry))
  if err != nil{
    fmt.Printf("error appending to log file: %v\n", err)
    return err
  }

  return nil
}

func (s *State) SavePersistentState() error {
  file, err := os.Create(fmt.Sprintf("%s.json", s.Name))
  if err != nil{
    fmt.Printf("error creating file: %v\n", err)
    return err
  }
  defer file.Close()

  data, err := json.Marshal(s.PersistentState)
  if err != nil{
    fmt.Printf("error marshalling persistent state: %v\n", err)
    return err
  }

  _, err = file.Write(data)
  if err != nil{
    fmt.Printf("error writing to state file: %v\n", err)
    return err
  }

  return nil
}

func (s *State) LoadPersistentState() error {
  file, err := os.Open(fmt.Sprintf("%s.json", s.Name))
  defer file.Close()
  if err != nil{
    fmt.Printf("error reading persistent state: %v\n", err)
    return err
  }
  buffer := make([]byte, 1024)
  bytesRead, err := file.Read(buffer)
  if err != nil{
    fmt.Printf("error reading file: %v\n", err)
    return err
  }
  err = json.Unmarshal(buffer[:bytesRead], &s.PersistentState)
  if err != nil{
    fmt.Printf("error unmarshaling persistent state")
    return err
  }
  return nil
}
