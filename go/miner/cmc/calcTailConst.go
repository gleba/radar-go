package cmc

import (
	"github.com/montanaflynn/stats"
	"radar.cash/core/hand"
	"radar.cash/core/sol"
	"radar.cash/core/sol/si"
)

func calcTailConst(price []float64, vol []float64, isBTC bool) sol.TailCoinConst {
	rate := make(sol.TailCoinConst, 20)
	meanPrice := hand.SafeFloat64(stats.Mean(price))
	rate[si.PriceMean] = meanPrice
	rate[si.PriceMedian] = hand.SafeFloat64(stats.Median(price))
	rate[si.PriceGeometric] = hand.SafeFloat64(stats.GeometricMean(price))
	rate[si.PriceHarmonic] = hand.SafeFloat64(stats.HarmonicMean(price))
	rate[si.PriceVariance] = hand.SafeFloat64(stats.SampleVariance(price))
	if len(vol) > 1 {
		rate[si.VolumeVariance] = hand.SafeFloat64(stats.SampleVariance(vol))
	} else {
		rate[si.VolumeVariance] = -1
	}
	rate[si.VolumeMean] = hand.SafeFloat64(stats.Mean(vol))
	rate[si.VolumeMedian] = hand.SafeFloat64(stats.Median(vol))
	rate[si.VolumeGeometric] = hand.SafeFloat64(stats.GeometricMean(vol))
	rate[si.VolumeHarmonic] = hand.SafeFloat64(stats.HarmonicMean(vol))

	rate[si.Percentile30] = hand.SafeFloat64(stats.Percentile(price, 30))
	rate[si.Percentile60] = hand.SafeFloat64(stats.Percentile(price, 60))
	rate[si.Percentile90] = hand.SafeFloat64(stats.Percentile(price, 90))

	rate[si.PercentileNearestRank30] = hand.SafeFloat64(stats.PercentileNearestRank(price, 30))
	rate[si.PercentileNearestRank60] = hand.SafeFloat64(stats.PercentileNearestRank(price, 60))
	rate[si.PercentileNearestRank90] = hand.SafeFloat64(stats.PercentileNearestRank(price, 90))

	if meanPrice != 1 {
		rate[si.Volatility] = calcVolatility(isBTC, price, meanPrice, stats.Mean)
		rate[si.VolatilityMedian] = calcVolatility(isBTC, price, rate[si.PriceMedian], stats.Median)
		rate[si.VolatilityHarmonic] = calcVolatility(isBTC, price, rate[si.PriceHarmonic], stats.HarmonicMean)
		rate[si.VolatilityGeometric] = calcVolatility(isBTC, price, rate[si.PriceGeometric], stats.GeometricMean)
	}

	return rate
}

func makeMultiTail(usdPrice []float64, usdVol []float64, btcPrice []float64, btcVol []float64) sol.MultiTail {
	return sol.MultiTail{
		USD: calcTailConst(usdPrice, usdVol, false),
		BTC: calcTailConst(btcPrice, btcVol, true),
	}
}

type AvgFn func(data stats.Float64Data) (float64, error)

func calcVolatility(fixFloat bool, arr []float64, avg float64, avgFn AvgFn) float64 {
	var pav []float64
	var k = 1.0
	if fixFloat {
		k = 1.0//00000
	}
	for _, v := range arr {
		v = v * k
		p1 := avg * k - v
		pav = append(pav, p1*p1)
	}
	cv := hand.FiniteFloat(float64(len(arr)) / ( hand.SafeFloat64(avgFn(pav)) * k))
	//vv := (cv/avg)*100
	//fmt.Println("avg", vv)
	return cv
}
