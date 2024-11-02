package internal

import (
	"context"
	"log"
	"net"

	pb "github.com/ashupednekar/raft-go/internal/raft"
	"google.golang.org/grpc"
)

type server struct {
  pb.UnimplementedRaftServiceServer
}

func (s *server) AppendEntries(ctx context.Context, in *pb.EntryInput) (*pb.EntryResult, error){
  return &pb.EntryResult{}, nil
}

func Serve(){
  ln, err := net.Listen("tcp", ":8001")
  if err != nil{
    log.Fatalf("failed to listen at port 8001: %v", err)
  }

  s := grpc.NewServer()
  pb.RegisterRaftServiceServer(s, &server{})

  log.Printf("gRPC server listening at %v", ln.Addr())
  if err := s.Serve(ln); err != nil{
    log.Fatalf("failed to start gRPC server: %v", err)
  }
}
