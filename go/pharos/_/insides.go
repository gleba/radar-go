package _

import (
	"radar.cash/core/sol"
)

type AliveTune = struct {
	owner       int
	tuneId      string
	coinId      uint32
	impulse     float64
	priceChange float64
}

func insides(pulse *sol.CoinPulse, tails sol.Tails) []AliveTune {
	var alives []AliveTune
	tunesBox.Range(func(Id string, tune sol.UserTune) bool {
		if inCoverage(pulse, &tune.Coverage) {
			haveI, impulseR := tails.Qerry(tune.Impulse.Period, tune.Impulse.LV, tune.Impulse.K)
			if haveI && impulseR < tune.Impulse.Rate*10 {
				haveI = false
			}
			if haveI {
				haveC, priceC := tails.Qerry(tune.Corridor.Period, tune.Corridor.LV, tune.Corridor.K)
				if haveC && priceC <= tune.Corridor.Up && priceC >= (tune.Corridor.Down*-1) {
					alives = append(alives, AliveTune{
						owner:       tune.Author,
						tuneId:      tune.ID,
						coinId:      pulse.ID,
						impulse:     impulseR,
						priceChange: priceC,
					})
				}
			}
		}
		return true
	})
	return alives
}

//func isTargeted(pulse *sol.CoinPulse, tail []sol.MultiTail, tune *sol.UserTune) bool {
//	k := tail[tune.Impulse.K]
//	return false
//}

func inCoverage(pulse *sol.CoinPulse, cov *sol.CoverageTune) bool {
	if pulse.VolByLV(cov.VolLV) < cov.VolValue && pulse.CapByLV(cov.CapLV) < cov.CapValue {
		return false
	}
	marketMap, haveMarkets := marketMapBox.Load(pulse.ID)
	if haveMarkets {
		for _, marketId := range cov.Markets {
			if marketMap[marketId] {
				return true
			}
		}
	}
	return false
}
