package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ashupednekar/raft-go/internal/server/raft"
	"google.golang.org/grpc"
)

type server struct {
  name string
  pb.UnimplementedRaftServiceServer
}

func (s *server) AppendEntries(ctx context.Context, in *pb.EntryInput) (*pb.EntryResult, error){
  return &pb.EntryResult{}, nil
}

func StartServer(name string, port int){
  ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
  if err != nil{
    log.Fatalf("failed to listen at port 8001: %v", err)
  }

  s := grpc.NewServer()
  pb.RegisterRaftServiceServer(s, &server{name: name})

  log.Printf("gRPC server %s listening at %v", name, ln.Addr())
  if err := s.Serve(ln); err != nil{
    log.Fatalf("failed to start gRPC server: %v", err)
  }
}
