package magic

import (
	f "github.com/ambelovsky/gosf"
	"radar.cash/core/heat"
	"radar.cash/pharos/pegas/sio"
	"radar.cash/pharos/pegas/wings"
	"radar.cash/pharos/vars"
)

//func addFront(fi []*wings.FrontImpulse, a *wings.ActiveAlert) []*wings.FrontImpulse {
//	last := time.Unix(a.TimeLast,0)
//	if
//}

func BroadCastFrontImpulses(activeAlerts []*wings.ActiveAlert) {
	fi := []*wings.FrontImpulse{}
	for _, a := range activeAlerts {
		fi = append(fi, makeFrontAlert(a))
	}
	f.Broadcast("main", "impulses", sio.BodyData(fi))
}
func MakeFrontImpulses() []*wings.FrontImpulse {
	aa := []*wings.FrontImpulse{}
	//aa := []*wings.ActiveAlert{}
	vars.ActiveAlerts.Range(func(key uint32, value wings.ActiveAlert) bool {
		aa = append(aa, makeFrontAlert(&value))
		//aa = append(aa, &value)
		return true
	})
	return aa
}
func makeFrontAlert(alert *wings.ActiveAlert) *wings.FrontImpulse {
	//pulse, _ := heat.Pulses.Load(alert.CoinId)
	quote, _ := heat.Quotas.Load(alert.CoinId)
	return &wings.FrontImpulse{
		Id: alert.Id,
		Coin: wings.FrontCoin{
			ID:     quote.ID,
			Name:   quote.Name,
			Symbol: quote.Symbol,
			Slug:   quote.Slug,
		},
		//Pulse:     pulse,
		Alerts:    alert.Alerts,
		TimeFirst: alert.TimeFirst,
		TimeLast:  alert.TimeLast,
	}
}
