package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	pb "github.com/ashupednekar/raft-go/internal/server/raft"
	"github.com/ashupednekar/raft-go/internal/state"
	"google.golang.org/grpc"
)

type Server struct {
  State state.State
  LastHeartBeat time.Time
  QuitLeadingChan chan bool
  pb.UnimplementedRaftServiceServer
}

func NewServer() Server {
  port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
  if err != nil{
    log.Fatalf("port env not set: %v", err)
  }

  server_id, err := strconv.Atoi(os.Getenv("SERVER_ID"))
  if err != nil{
    log.Fatalf("SERVER_ID env missing: %v\n", err)
  }

  return Server{State: state.State{Id: server_id, Port: port}}
}

func (s *Server) Start (){
  ln, err := net.Listen("tcp", fmt.Sprintf(":%d", s.State.Port))
  if err != nil{
    log.Fatalf("failed to listen at port 8001: %v", err)
  }

  grpcServer := grpc.NewServer()
  pb.RegisterRaftServiceServer(grpcServer, s)

  log.Printf("gRPC server %d listening at %v as %v", s.State.Id, ln.Addr(), s.State.Role)
  if err := grpcServer.Serve(ln); err != nil{
    log.Fatalf("failed to start gRPC server: %v", err)
  }
}
