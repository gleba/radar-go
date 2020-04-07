package cmc

import (
	"encoding/json"
	"net/http"
	"radar.cash/core/sol"
	"sync"

	"radar.cash/core/hand"
	"time"
)

var client = &http.Client{}
var pulseWriter sol.PulseWriter
var latestSync sync.Map // map[uint32]sol.CCoin

func MineLatest() {
	var latestQuery *sol.CmcListingQuery
	bytes := request("https://web-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?convert=USD,BTC&cryptocurrency_type=all&limit=5000")
	hand.Safe(json.Unmarshal(bytes, &latestQuery))
	pulseWriter = sol.CreatePulseWriter()
	for _, coin := range latestQuery.Data {
		_, isStable := stableCons[coin.Symbol]
		if isStable {
			continue
		}
		prev, haveAprev := latestSync.Load(coin.ID)
		if haveAprev {
			prevSince := time.Since(prev.(sol.CCoin).LastUpdated)
			lastSince := time.Since(coin.LastUpdated)
			diff := prevSince.Seconds() - lastSince.Seconds()
			if diff > 0 {
				updatePulse(coin)
			}
		} else {
			updatePulse(coin)
		}
	}
	pulseWriter.Commit()
}

func updatePulse(coin sol.CCoin) {
	latestSync.Store(coin.ID, coin)
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

	pulseWriter.Add(pulse, coin)
}
