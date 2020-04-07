package cmc

import (
	"fmt"
	"radar.cash/core/intel/df"
	"radar.cash/core/sol"
	"sync"
)

var lastDaily sync.Map //map[uint32]string{}

//var marketMap = map[int]sol.Market{}
//var coinMarkets = make(map[uint32][]sol.Market)

//type CoinMarketsMap sync.Map
var coinMarketSync sync.Map
var exchangesMapSync sync.Map

func WarmCaches() {
	df.DailyConst.Init(func(dailyList []sol.DailyConst) {
		for _, dc := range dailyList {
			lastDaily.Store(dc.ID,dc.Day)
		}
		fmt.Println("warmed DailyConst")
	})
	df.Market.Init(func(mc []sol.CoinMarkets) {
		for _, c := range mc {
			coinMarketSync.Store(c.ID, c.Markets)
			for _, m := range c.Markets {
				exchangesMapSync.Store(m.Exchange.ID, m.Exchange)
			}
		}
		fmt.Println("warmed CoinMarkets")
	})
}