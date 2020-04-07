package scifi

import (
	"radar.cash/core/sol"
	"radar.cash/pharos/box"
)


var dailyBox box.DailyConst
var pulseBox box.CoinPulse


func ReceiveDailyConst(values []sol.DailyConst)  {
	for _, v := range values {
		dailyBox.Store(v.ID, v)
	}
}

func ReceivePulse(values []sol.CoinPulse)  {
	for _, v := range values {
		pulseBox.Store(v.ID, v)
		changes := reuse(v)
		if changes != nil {
			insides(v.ID, changes)
		}
	}
}