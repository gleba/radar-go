package scifi

import (
	"math"
	"radar.cash/core/sol"
	"radar.cash/core/sol/si"
	"radar.cash/pharos/box"
)

//var pulseBox box.CoinPulse

var percentBox box.DailyConst
var USD = "USD"
var BTC = "BTC"




func reuse(pulse sol.CoinPulse) []sol.MultiTail {
	daily, have := dailyBox.Load(pulse.ID)
	if have && daily.Fine {
		percentPeriods := make([]sol.MultiTail, len(daily.Tails))
		for period, tail := range daily.Tails {
			multiTail := make(sol.MultiTail, 2)
			for LV, con := range tail {
				switch LV {
				case USD:
					multiTail[LV] = calcChange(con, pulse.PriceUSD, pulse.VolumeUSD)
				case BTC:
					multiTail[LV] = calcChange(con, pulse.PriceBTC, pulse.VolumeBTC)
				}
			}
			percentPeriods[period] = multiTail
		}
		return percentPeriods
	}
	return nil
}

func PercentChange(past float64, now float64) float64 {
	return 100 - math.Round((past/now)*100)
}

func calcChange(con []float64, price float64, volume float64) []float64 {
	changes := make([]float64, 10)
	for k, value := range con {
		switch k {
		case si.PriceGeometric:
		case si.PriceHarmonic:
		case si.PriceMean:
		case si.PriceMedian:
		case si.PriceVariance:
			changes[k] = PercentChange(value, price)
		case si.VolumeGeometric:
		case si.VolumeHarmonic:
		case si.VolumeMean:
		case si.VolumeMedian:
		case si.VolumeVariance:
			changes[k] = PercentChange(value, volume)
		}
	}
	return changes
}
