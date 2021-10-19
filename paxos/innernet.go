package paxos

import (
	"time"

	"github.com/Gavin-Lau/hi-level/protocol/PAXOS"
)

type Paxos struct {
	proposer chan PAXOS.StateMsg
	acceptor chan PAXOS.StateMsg
	learner  chan PAXOS.StateMsg
	outTime  time.Duration
}

func newPaxosServer() (*Paxos, error) {

	return &Paxos{
		proposer: make(chan PAXOS.StateMsg, 1024),
		acceptor: make(chan PAXOS.StateMsg, 1024),
		learner:  make(chan PAXOS.StateMsg, 1024),
	}, nil
}

func (p *Paxos) handleReq() error {
	msg := PAXOS.StateMsg{
		Mtype: PAXOS.MsgType_PAXOS_Prepare,
	}
	//broadcast Prepare to all acceptor
	p.proposer <- msg
	return nil
}

func (p *Paxos) handleAccept() error {

	return nil
}
