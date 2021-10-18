package paxos

import (
	"time"

	pb "github.com/Gavin-Lau/hi-level/protocol"
)

type Paxos struct {
	proposer chan pb.StateMsg
	acceptor chan pb.StateMsg
	learner  chan pb.StateMsg
	outTime  time.Duration
}

func newPaxosServer() (*Paxos, error) {

	return &Paxos{
		proposer: make(chan pb.StateMsg, 1024),
		acceptor: make(chan pb.StateMsg, 1024),
		learner:  make(chan pb.StateMsg, 1024),
	}, nil
}

func (p *Paxos) handleReq() error {
	msg := pb.StateMsg{
		Mtype: pb.MsgType_PAXOS_Prepare,
	}
	p.proposer <- msg
	return nil
}

func (p *Paxos) handleAccept() error {

	return nil
}
