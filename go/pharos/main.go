package main

import (
	"fmt"
	"radar.cash/core/data/eSpace"
	"radar.cash/core/data/service"
	"radar.cash/core/heat"
	"radar.cash/core/sol"
	"radar.cash/pharos/pegas"
	"time"
)

func main() {
	service.OpenNATS()
	service.OpenClickHose()

	//rain.Market.Sub(pegas.ReceiveMarkets)
	//rain.Market.Sub(func(cm *sol.CoinMarkets) {
	//	fmt.Println("market", cm.ID)
	//})

	heat.RestorePulse()
	heat.Pulses.Range(func(key uint32, value *sol.CoinPulse) bool {
		pegas.AcceptPulse(value)
		return true
	})
	eSpace.Pulses.Sub(func(pulses []*sol.CoinPulse) {
		fmt.Println(pulses)
	})
	time.Sleep(time.Duration(10000) * time.Hour)
}
