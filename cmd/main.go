package main

import (
	"log"
	"os"
	"strconv"

	"github.com/ashupednekar/raft-go/internal/server"
)

func main(){
  port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
  if err != nil{
    log.Fatalf("port env not set: %v", err)
  }
  go server.StartServer(os.Getenv("SERVER_NAME"), port)

  select{}
}
