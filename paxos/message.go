package paxos

type MsgType uint8

const (
	Prepare MsgType = iota
	Promise
	Propose
	Accept
)

type message struct {
	tp    MsgType
	id    int
	value string
}
