package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/ashupednekar/raft-go/internal"
	"github.com/ashupednekar/raft-go/internal/server"
)

func main(){
  port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
  if err != nil{
    log.Fatalf("port env not set: %v", err)
  }

  s := server.Server{}
  s.LastHeartBeat = time.Now()

  server_id, err := strconv.Atoi(os.Getenv("SERVER_ID"))
  if err != nil{
    log.Fatalf("SERVER_ID env missing: %v\n", err)
  }
  go s.Start(server_id, port)

  go func(s *server.Server){
    electionTimeout, err := time.ParseDuration(fmt.Sprintf("%dms", rand.Intn(6000)+ 5850))
    if err != nil{
      log.Fatalf("error calculating election timeout: %v", err)
    }
    for {
      fmt.Printf("last: %v | now: %v| timeout: %v\n", s.LastHeartBeat, time.Now(), electionTimeout)
      if s.LastHeartBeat.Before(time.Now().Add(-electionTimeout)){
        fmt.Println("No viable leader found, initiating election")
        internal.InitiateElection(s)
      } 
      time.Sleep(electionTimeout)
    }
  }(&s)

  select{}
}
