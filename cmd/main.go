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
    electionTimeout, err := time.ParseDuration(fmt.Sprintf("%dms", rand.Intn(151)+ 150))
    if err != nil{
      log.Fatalf("error calculating election timeout: %v", err)
    }
    for {
      if s.LastHeartBeat.Before(time.Now().Add(-electionTimeout-time.Millisecond * 100)){
        fmt.Println("No viable leader found, initiating election")
        internal.InitiateElection(s)
      } 
      time.Sleep(electionTimeout)
    }
  }(&s)

  select{}
}
