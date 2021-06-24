package _

import (
	"radar.cash/core/sol"
	"radar.cash/pharos/box"
)

var dailyBox box.DailyConst

//var pulseBox box.CoinPulse
var tunesBox box.UserTune
var marketMapBox box.MarketMap

func ReceiveDailyConst(values []sol.DailyConst) {
	for _, v := range values {
		dailyBox.Store(v.ID, v)
	}
}

func ReceiveCoin(values []sol.CoinQuote) {
	for _, coin := range values {
		//pulseBox.Store(pulse.ID, pulse.Pulse())
		pulse := coin.Pulse()
		changes := refine(pulse)
		if changes != nil && len(changes) > 2 {
			upFloor(coin, insides(&pulse, changes))
		}
	}
}
func ReceiveTunes(values []sol.UserTune) {
	for _, v := range values {
		tunesBox.Store(v.ID, v)
	}
}
func ReceiveMarkets(values []sol.CoinMarkets) {
	for _, c := range values {
		mmap := map[uint32]bool{}
		for _, market := range c.Markets {
			mmap[market.ID] = true
		}
		marketMapBox.Store(c.ID, mmap)
	}
}
