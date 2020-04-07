package cmc

import (
	"radar.cash/core/sol"
)

const USD = "USD"
const BTC = "BTC"

func reverse(ss []sol.Quotes) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func makeDailyConst(coinID uint32, day string, storyQuery *sol.CmcStoryQuery) sol.DailyConst {
	var btcPrice []float64
	var usdPrice []float64
	var btcVol []float64
	var usdVol []float64
	var tails []sol.MultiTail
	dailyConst := sol.DailyConst{
		ID:    coinID,
		Day:   day,
		Fine:  false,
		Tails: nil,
	}

	reverse(storyQuery.Data.Quotes)

	for day, i := range storyQuery.Data.Quotes {
		btcPrice = storyAddPrice(btcPrice, i.Quote.BTC)
		usdPrice = storyAddPrice(usdPrice, i.Quote.USD)
		btcVol = storyAdd(btcVol, i.Quote.BTC.Volume)
		usdVol = storyAdd(usdVol, i.Quote.USD.Volume)
		switch day {
		case 0:
			if btcPrice == nil || btcVol == nil {
				return dailyConst
			}
			tails = append(tails, makeMultiTail(usdPrice, usdVol, btcPrice, btcVol))
		case 7:
			if len(btcVol) < 5 {
				return dailyConst
			}
			if btcPrice == nil || btcVol == nil {
				return dailyConst
			}
			tails = append(tails, makeMultiTail(usdPrice, usdVol, btcPrice, btcVol))
		case 30:
			if len(btcVol) < 25 {
				return dailyConst
			}
			tails = append(tails, makeMultiTail(usdPrice, usdVol, btcPrice, btcVol))
		case 60:
			if len(btcVol) < 50 {
				return dailyConst
			}
			tails = append(tails, makeMultiTail(usdPrice, usdVol, btcPrice, btcVol))
		case 45:
			tails = append(tails, makeMultiTail(usdPrice, usdVol, btcPrice, btcVol))
		case 89:
			if len(btcVol) < 70 {
				return dailyConst
			}
			tails = append(tails, makeMultiTail(usdPrice, usdVol, btcPrice, btcVol))
		}
	}
	dailyConst.Fine = true
	dailyConst.Tails = tails
	return dailyConst
}
