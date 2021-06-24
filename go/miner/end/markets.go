package end

import (
	"radar.cash/core/data/eSpace"
	"radar.cash/core/data/rocks"
	"radar.cash/core/sol"
	"sync"
)

var Markets sync.Map

func UpdateMarkets(cm *sol.CoinMarkets) {
	Markets.Store(cm.ID, true)
	eSpace.Markets.Publish(cm)
	rocks.Markets.Store(cm.ID, cm.Markets)
}

func restoreMarkets() {
	for _, rockRaw := range rocks.Markets.LoadAll() {
		Markets.Store(rockRaw.Id, true)
	}
}
