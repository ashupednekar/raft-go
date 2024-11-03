package internal

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ashupednekar/raft-go/internal/client"
	"github.com/ashupednekar/raft-go/internal/server"
)

type AppendResult struct{
  Addr string
  Err error
  Success bool
  Term int
}

func StartLeading(s *server.Server) {
  for{
    time.Sleep(time.Millisecond * 100)
    results := make(chan AppendResult)
    var wg sync.WaitGroup
    servers := strings.Split(os.Getenv("SERVERS"), ",")
    for _, addr := range servers{
      wg.Add(1)
      go func(addr string){
        defer wg.Done()
        c, err := client.Connect(addr)
        if err != nil{
          fmt.Printf("couldnt's connect to follower at: %s\n", addr)
          results <- AppendResult{Addr: addr, Err: err} 
        }
        term, success, err := client.AppendEntries(c, s)
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

    /*for result := range results{
      fmt.Printf("appendEntries result: %v\n", result)
    }*/

    select {
    case <- s.QuitLeadingChan:
      fmt.Printf("Leader %d stepping down, is now a follower\n", s.State.Id)
      break
    default:
    }
  }
}
