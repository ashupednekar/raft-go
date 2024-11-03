package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/ashupednekar/raft-go/internal"
	"github.com/ashupednekar/raft-go/internal/server"
)

func main(){

  s := server.NewServer()
  s.LastHeartBeat = time.Now()

  go s.Start()

  go func(s *server.Server){
    electionTimeout, err := time.ParseDuration(fmt.Sprintf("%dms", rand.Intn(2001)+ 2000))
    if err != nil{
      log.Fatalf("error calculating election timeout: %v", err)
    }
    for {
      fmt.Printf("last: %v", s.LastHeartBeat)
      if s.LastHeartBeat.Before(time.Now().Add(-electionTimeout)) {
          fmt.Printf("Condition met: Last heartbeat (%v) is older than current time (%v) minus election timeout (%v)\n", 
              s.LastHeartBeat, time.Now(), electionTimeout)
        internal.InitiateElection(s)
      } else {
          fmt.Printf("Condition not met: Last heartbeat (%v) is NOT older than current time (%v) minus election timeout (%v)\n", 
              s.LastHeartBeat, time.Now(), electionTimeout)
      }

      time.Sleep(electionTimeout)
    }
  }(&s)

  select{}
}
