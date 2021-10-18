package paxos

import (
	"context"

	pb "github.com/Gavin-Lau/hi-level/protocol"
)

type server struct {
	pb.UnimplementedSessionalServer
}

func (s *server) Put(ctx context.Context, in *pb.PAXOS.Request) (*pb.Learn, error) {
	//create
	return &pb.Learn{Tid: 100}, nil
}
