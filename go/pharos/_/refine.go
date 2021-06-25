package _

import (
	"math"
	"radar.cash/core/heat/box"
	"radar.cash/core/sol"
	"radar.cash/core/sol/si"
)

//var pulseBox box.CoinPulse

var percentBox box.DailyConsts
var USD = "USD"
var BTC = "BTC"

func refine(pulse sol.CoinPulse) []sol.MultiTail {
	daily, have := dailyBox.Load(pulse.ID)
	if have {
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

func PercentChange(target float64, now float64) float64 {
	return math.Floor((1-(target/now))*1000) / 10
}

func calcChange(con []float64, nowPrice float64, nowVolume float64) []float64 {
	changes := make([]float64, 10)
	for k, kValue := range con {
		switch k {
		case si.PriceGeometric, si.PriceHarmonic, si.PriceMean, si.PriceMedian, si.PriceVariance:
			changes[k] = PercentChange(kValue, nowPrice)
		case si.VolumeGeometric, si.VolumeHarmonic, si.VolumeMean, si.VolumeMedian, si.VolumeVariance:
			changes[k] = PercentChange(kValue, nowVolume)
		}
	}
	return changes
}
