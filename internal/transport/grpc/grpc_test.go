package grpc

import (
	"context"
	"net"
	"testing"

	"google.golang.org/grpc"

	raftpb "github.com/potix2/raft/proto/raft/v1"
)

func setupTestGRPCServerAndClient(t *testing.T) (*grpc.Server, raftpb.RaftServiceClient, func()) {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		t.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	raftpb.RegisterRaftServiceServer(server, &RaftServiceServer{})
	go func() {
		server.Serve(lis)
	}()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial: %v", err)
	}

	client := raftpb.NewRaftServiceClient(conn)

	teardown := func() {
		server.GracefulStop()
		conn.Close()
		lis.Close()
	}
	return server, client, teardown
}

func TestRequestVote(t *testing.T) {
	_, client, teardown := setupTestGRPCServerAndClient(t)
	defer teardown()

	request := &raftpb.RequestVoteRequest{
		Term:         1,
		CandidateId:  1,
		LastLogIndex: 0,
		LastLogTerm:  0,
	}

	response, err := client.RequestVote(context.Background(), request)

	if err != nil {
		t.Fatalf("Failed to call RequestVote: %v", err)
	}

	if response.Term != 1 {
		t.Fatalf("Term is not 1: %v", response.Term)
	}
	if response.VoteGranted != true {
		t.Fatalf("VoteGranted is not true: %v", response.VoteGranted)
	}
}
