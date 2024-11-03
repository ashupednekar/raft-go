package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/ashupednekar/raft-go/internal"
	"github.com/ashupednekar/raft-go/internal/client"
	"github.com/ashupednekar/raft-go/internal/server"
	"github.com/ashupednekar/raft-go/internal/state"
)


type VoteResult struct{
  Addr string
  VoteGranted bool
  Term int
  Err error
}


func InitiateElection(s *server.Server) error {
  s.State.PersistentState.CurrentTerm = s.State.PersistentState.CurrentTerm + 1
  s.State.PersistentState.VotedFor = s.State.Id
  s.State.Role = state.Candidate
  s.State.SavePersistentState()

  server_count, err := strconv.Atoi(os.Getenv("SERVER_COUNT"))
  if err != nil{
    fmt.Println("error reading server count env")
  }
  majority := server_count / 2 + 1

  var wg sync.WaitGroup
  results := make(chan VoteResult)

  servers := strings.Split(os.Getenv("SERVERS"), ",")
  for _, addr := range servers{
    wg.Add(1)
    go func(addr string){
      defer wg.Done()
      fmt.Printf("requesting vote from server at: %s\n", addr)
      c, err := client.Connect(addr)
      if err != nil{
        results <- VoteResult{Addr: addr, Err: err}
      }
      term, voteResult, err := client.RequestVote(c, s)
      if err != nil{
        fmt.Printf("error requesting vote from server at: %s - %v\n", addr, err)
        results <- VoteResult{Addr: addr, Err: err}
      }
      results <- VoteResult{Addr: addr, VoteGranted: voteResult, Term: term, Err: nil} 
    }(addr)
  }

  go func(){
    wg.Wait()
    close(results)
  }()

  var count int
  for result := range results{
    fmt.Printf("Vote result from server at: %s - %v with term %d\n", result.Addr, result.VoteGranted, result.Term)
    if result.VoteGranted{
      count++
    }
  }

  fmt.Printf("Total votes received for server %d: %d, majority needed: %d\n", s.State.Id, count, majority)
  if count >= majority{
    fmt.Printf("election complete, winner: server: %d", s.State.Id)
    internal.StartLeading()
  }else{
    fmt.Printf("server %d countn't obtain majority votes, election failed\n", s.State.Id)
  }

  return nil
}
