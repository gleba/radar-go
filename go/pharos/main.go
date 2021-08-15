package main

import (
	"net/http"
	"radar.cash/core/data/eSpace"
	"radar.cash/core/data/rocks"
	"radar.cash/core/data/service"
	"radar.cash/core/heat"
	"radar.cash/core/sol"
	"radar.cash/core/tool"
	"radar.cash/pharos/pegas"
	"radar.cash/pharos/pegas/sio"
	"radar.cash/pharos/pegas/wings"
	"radar.cash/pharos/vars"
	"time"

	_ "radar.cash/pharos/pegas/sio/handlers"
)

func main() {
	service.OpenNATS()
	service.OpenClickHose()
	service.SetupPostgres(
		(*wings.ActiveAlert)(nil),
		(*wings.Market)(nil),
	)

	//pegas.RestoreAlerts()
	sio.Start()
	rocks.Init()
	pegas.SyncMarkets()
	pegas.SyncPresets()

	heat.RestoreQuota()
	heat.RestoreDaily()
	heat.RestorePulse()

	heat.Pulses.Range(func(key uint32, value sol.CoinPulse) bool {
		pegas.AcceptPulse([]*sol.CoinPulse{&value})
		return true
	})

	eSpace.Pulses.Sub(pegas.AcceptPulse)
	go sio.StartHttp(func(w http.ResponseWriter, r *http.Request) {
		pegas.SyncPresets()
		sio.UpdateDict()
		vars.ActiveAlerts.Range(func(key uint32, value wings.ActiveAlert) bool {
			vars.ActiveAlerts.Delete(key)
			return true
		})
		heat.Pulses.Range(func(key uint32, value sol.CoinPulse) bool {
			pegas.AcceptPulse([]*sol.CoinPulse{&value})
			return true
		})
	})

	writeTicker := time.NewTicker(time.Second * 5)

	for {
		select {
		case _ = <-writeTicker.C:
			tool.MemState()
			pegas.WriteAlerts()
		}
	}

}
