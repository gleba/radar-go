package cmc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"radar.cash/core/heat"
	"radar.cash/core/sol"

	"radar.cash/core/hand"
)

var client = &http.Client{}

func MineLatest() {
	var latestQuery sol.CmcListingQuery
	var pulseWriter = CreatePulseWriter()
	bytes := request("https://web-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?convert=USD,BTC&cryptocurrency_type=all&limit=5000")
	hand.Safe(json.Unmarshal(bytes, &latestQuery))

	updatePulse := func(coin sol.CoinQuote) {
		pulse := sol.CoinPulse{
			PriceVol: sol.PriceVol{
				PriceBTC:  coin.Quote.BTC.Price,
				PriceUSD:  coin.Quote.USD.Price,
				VolumeUSD: coin.Quote.USD.Volume24H,
				VolumeBTC: coin.Quote.BTC.Volume24H,
			},
			MarketCap: sol.MarketCap{
				MarketCapBTC: coin.Quote.BTC.MarketCap,
				MarketCapUSD: coin.Quote.USD.MarketCap,
			},
			ID:   coin.ID,
			Time: coin.LastUpdated,
		}
		pulseWriter.Add(&pulse, &coin)
		heat.Quotas.Store(coin.ID, coin)
	}

	for _, coin := range latestQuery.Data {

		_, isStable := stableCons[coin.Symbol]
		if isStable {
			continue
		}
		prev, havePrevSyncTime := heat.Quotas.Load(coin.ID)
		if havePrevSyncTime {
			//prevSince := time.Since(prev.LastUpdated)
			//lastSince := time.Since(coin.LastUpdated)
			//coin.LastUpdated.Sub(prev.LastUpdated)
			diff := coin.LastUpdated.Sub(prev.LastUpdated)

			if diff.Seconds() > 0 {
				updatePulse(coin)
				if coin.ID == 4233 {
					fmt.Println(diff.Seconds(), coin.LastUpdated)
					fmt.Println(coin.LastUpdated)
					fmt.Println(prev.LastUpdated)
				}
			} else {
				//fmt.Println("no changes")
				//fmt.Println(coin.LastUpdated)
				//fmt.Println(prev.LastUpdated)
			}
		} else {
			updatePulse(coin)
		}
	}
	pulseWriter.Commit()
}
