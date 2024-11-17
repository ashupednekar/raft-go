package server 

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ashupednekar/raft-go/internal/state"
)

type AppendResult struct{
  Addr string
  Err error
  Success bool
  Term int
}

func SpawnAppends(s *Server) chan AppendResult {
  results := make(chan AppendResult)
  var wg sync.WaitGroup
  servers := strings.Split(os.Getenv("SERVERS"), ",")
  for _, addr := range servers{
    wg.Add(1)
    go func(addr string){
      defer wg.Done()
      c, err := Connect(addr)
      if err != nil{
        fmt.Printf("couldnt's connect to follower at: %s\n", addr)
        results <- AppendResult{Addr: addr, Err: err} 
      }
      term, success, err := AppendEntries(c, s)
      if err != nil{
        fmt.Printf("error appending entries at: %s\n", addr)
        results <- AppendResult{Addr: addr, Err: err} 
      }
      results <- AppendResult{Addr: addr, Success: success, Term: term, Err: nil} 
    }(addr)
  }

  go func(){
    wg.Wait()
    close(results)
  }()

  return results
}

func StartLeading(s *Server) {
  s.State.Role = state.Leader
  for{

    SpawnAppends(s)

    select {
    case <- s.QuitLeadingChan:
      fmt.Printf("Leader %d stepping down, is now a follower\n", s.State.Id)
      break
    default:
    }
    time.Sleep(time.Millisecond * 100)
  }
}
