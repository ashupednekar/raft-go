package internal

import (
	"context"
	"log"
	"time"

	pb "github.com/ashupednekar/raft-go/internal/raft"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AppendEntries(){
  conn, err := grpc.Dial("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil{
    log.Fatalf("failed to connect to gRPC server: %v", err)
  }
  c := pb.NewRaftServiceClient(conn)

  ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 100)
  defer cancel()

  r, err := c.AppendEntries(ctx, &pb.EntryInput{})
  if err != nil{
    log.Fatalf("error calling rpc, AppendEntries: %v", err)
  }

  log.Printf("Response - term: %d, success: %v", r.Term, r.Success)
}
