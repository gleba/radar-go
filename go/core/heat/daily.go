package heat

import (
	"encoding/json"
	"fmt"
	"radar.cash/core/data/rocks"
	"radar.cash/core/hand"
	"radar.cash/core/heat/box"
	"radar.cash/core/sol"
)

var DailyConst box.DailyConsts

func RestoreDaily() {
	fmt.Println("RestoreDaily")
	r := rocks.DailyConst.LoadAll()
	fmt.Println("RestoreDaily", len(r))
	for _, raw := range r {
		var q sol.DailyConst
		hand.Safe(json.Unmarshal(raw.Data, &q))
		DailyConst.Store(q.ID, q)
	}
}
