package client

import (
	"context"
	"fmt"
	"time"

	"github.com/ashupednekar/raft-go/internal/server"
	pb "github.com/ashupednekar/raft-go/internal/server/raft"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(addr string) (pb.RaftServiceClient, error) {
  conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
  if err != nil{
    fmt.Printf("error connecting to server at: %s - %v\n", addr, err)
  }
  return pb.NewRaftServiceClient(conn), nil
}

func RequestVote(client pb.RaftServiceClient, s *server.Server) (int, bool, error){
  ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 100)
  defer cancel()
  r, err := client.RequestVote(ctx, &pb.VoteInput{
    Term: int32(s.State.PersistentState.CurrentTerm),
    CandidateId: int32(s.State.Id), // TODO: fill rest of the fields
  })
  if err != nil{
    fmt.Printf("requestVote rpc call failed: %v\n", err)
    return 0, false, err
  }
  return int(r.Term), r.VoteGranted, nil
}

func AppendEntries(client pb.RaftServiceClient) (int, bool, error) {
  ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 100)
  defer cancel()
  r, err := client.AppendEntries(ctx, &pb.EntryInput{
    // TODO: fill rest of the fields
  })
  if err != nil{
    fmt.Printf("appendEntries rpc call failed: %v\n", err)
    return 0, false, err
  }
  fmt.Printf("appendEntries result: %v", r)
  
  return int(r.Term), r.Success, nil
}
