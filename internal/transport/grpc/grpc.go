package grpc

import (
	"context"

	raftpb "github.com/potix2/raft/proto/raft/v1"
)

type RaftServiceServer struct {
	raftpb.UnimplementedRaftServiceServer
}

func (*RaftServiceServer) RequestVote(context.Context, *raftpb.RequestVoteRequest) (*raftpb.RequestVoteResponse, error) {
	return &raftpb.RequestVoteResponse{
		Term:        1,
		VoteGranted: true,
	}, nil
}
