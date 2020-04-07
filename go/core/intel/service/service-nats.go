package service

import (
	"github.com/nats-io/nats.go"
	"log"
	"radar.cash/core/hand"
)

var Nats *nats.Conn
var NatsEncoded *nats.EncodedConn

func OpenNATS() {
	var err error
	Nats, err = nats.Connect(nats.DefaultURL, nats.Token("2yKnjkfXCtA8ik2yKnjkfXCtA8ik"))
	hand.Safe(err)
	log.Print("nats: ok")
	NatsEncoded, _ = nats.NewEncodedConn(Nats, nats.GOB_ENCODER)
}


