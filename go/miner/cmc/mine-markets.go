package cmc

import (
	"encoding/json"
	"fmt"
	"radar.cash/core/hand"
	"radar.cash/core/intel/df"
	"radar.cash/core/sol"
	"strconv"
)

var marketPool chan sol.CCoin

func init() {
	marketPool = make(chan sol.CCoin)
	var err error
	hand.Safe(err)
	for range [3]int{} {
		go marketMinePool()
	}
}

func marketMinePool() {
	for {
		coin := <-marketPool
		_, found := coinMarketSync.Load(coin.ID)
		if !found {
			getMarket(coin)
		}
	}
}

func getMarket(coin sol.CCoin) {
	url := "https://web-api.coinmarketcap.com/v1/cryptocurrency/market-pairs/latest?aux=market_url&limit=4000&id=" + strconv.FormatInt(int64(coin.ID), 10)
	bytes := request(url)
	if bytes == nil {
		return
	}
	var query *sol.MarketQuery
	hand.Safe(json.Unmarshal(bytes, &query))

	var marketMap = map[int]sol.Market{}
	for _, mp := range query.Data.MarketPairs {
		newPair := sol.Pair{
			Pair: mp.MarketPair,
			URL:  mp.MarketURL,
		}
		xid := mp.Exchange.ID
		exchange := sol.Exchange{
			ID:   xid,
			Name: mp.Exchange.Name,
			Slug: mp.Exchange.Slug,
		}
		_, haveEx := exchangesMapSync.Load(xid)
		if !haveEx {
			exchangesMapSync.Store(xid, exchange)
		}
		market, have := marketMap[xid]
		if !have {
			market = sol.Market{
				Exchange: exchange,
				Pairs:    []sol.Pair{},
			}
		}
		market.Pairs = append(market.Pairs, newPair)
		marketMap[xid] = market
	}
	coinMarkets := sol.CoinMarkets{
		ID:      coin.ID,
		Markets: []sol.Market{},
	}
	for _, m := range marketMap {
		coinMarkets.Markets = append(coinMarkets.Markets, m)
	}
	coinMarketSync.Store(coin.ID, coinMarkets)
	df.Market.UpdateItem(coin.ID, coinMarkets)
	fmt.Print("★")
}
