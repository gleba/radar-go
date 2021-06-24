package heat

import (
	"encoding/json"
	"fmt"
	"radar.cash/core/data/rocks"
	"radar.cash/core/hand"
	"radar.cash/core/heat/box"
	"radar.cash/core/sol"
)

var Quotas box.Quotes

func RestoreQuota() {
	r := rocks.CoinQuote.LoadAll()
	fmt.Println("RestoreQuota", len(r))
	for _, raw := range r {
		var q sol.CoinQuote
		hand.Safe(json.Unmarshal(raw.Data, &q))
		Quotas.Store(q.ID, q)
	}
}
