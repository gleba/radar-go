package end

import (
	"radar.cash/core/data/eSpace"
	"radar.cash/core/data/rocks"
	"radar.cash/core/heat"
	"radar.cash/core/sol"
)

func UpdateDailyConst(dailyConst *sol.DailyConst) {
	heat.DailyConst.Store(dailyConst.ID, *dailyConst)
	rocks.DailyConst.Store(dailyConst.ID, dailyConst)
	eSpace.DailyConst.Publish(dailyConst)
}
