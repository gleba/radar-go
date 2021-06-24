package service

import (
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"radar.cash/core/hand"
)

var Nats *nats.Conn
var NatsEncoded *nats.EncodedConn

func OpenNATS() {
	var err error

	Nats, err = nats.Connect("nats://"+os.Getenv("NATS")+":4222", nats.Token("2yKnjkfXCtA8ik2yKnjkfXCtA8ik"))
	hand.Safe(err)
	NatsEncoded, _ = nats.NewEncodedConn(Nats, nats.GOB_ENCODER)
	log.Print("nats: ok")
}
