package paxos

import (
	"context"

	"github.com/Gavin-Lau/hi-level/protocol/PAXOS"
)

type Paxosserver struct {
	PAXOS.UnimplementedSessionalServer
}

func (s *Paxosserver) Put(ctx context.Context, in *PAXOS.Request) (*PAXOS.Learn, error) {
	//create
	return &PAXOS.Learn{Tid: 100}, nil
}

func (s *Paxosserver) Prepare(ctx context.Context, in *PAXOS.Request) (*PAXOS.Learn, error) {

	return &PAXOS.Learn{Tid: 100}, nil
}
