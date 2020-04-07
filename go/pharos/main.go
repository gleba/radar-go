package main

import (
	"math/rand"
	"radar.cash/core/intel/df"
	"radar.cash/core/intel/service"
	"radar.cash/pharos/scifi"
)

func main() {
	service.OpenNATS()
	df.DailyConst.Up(scifi.ReceiveDailyConst)
	df.Pulse.Up(scifi.ReceivePulse)
	service.Nats.Drain()
}

func randomID() string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 16)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
