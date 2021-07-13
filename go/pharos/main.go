package main

import (
	"radar.cash/core/data/service"
	"radar.cash/pharos/models"
	"time"
)

func main() {
	service.OpenNATS()
	service.OpenClickHose()
	service.SetupPostgres(
		(*models.ActiveAlert)(nil),
		(*models.Market)(nil),
	)

	//aa := models.ActiveAlert{
	//	Id:             0,
	//	CoinName:       "",
	//	AlertId:        0,
	//	DetectTime:     time.Time{},
	//	LastActiveTime: time.Time{},
	//}
	//rocks.Init()
	//pegas.SyncMarkets()
	//pegas.SyncPresets()
	//heat.RestoreQuota()
	//heat.RestoreDaily()
	//heat.RestorePulse()
	//heat.Pulses.Range(func(key uint32, value sol.CoinPulse) bool {
	//	pegas.AcceptPulse([]*sol.CoinPulse{&value})
	//	return true
	//})
	//eSpace.Pulses.Sub(pegas.AcceptPulse)
	time.Sleep(time.Duration(10000) * time.Hour)
}
