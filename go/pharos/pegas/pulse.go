package pegas

import (
	"math"
	"radar.cash/core/heat"
	"radar.cash/core/sol"
	"radar.cash/core/tool"
	"radar.cash/pharos/pegas/wings"
	"radar.cash/pharos/vars"
)

func AcceptPulse(pulses []*sol.CoinPulse) {
	for _, pulse := range pulses {
		heat.Pulses.Store(pulse.ID, *pulse)
		dconst, _ := heat.DailyConst.Load(pulse.ID)
		//fmt.Println(dconst)
		imarket := indexMarket[pulse.ID]
		vars.Alerts.Range(func(z int, alert wings.Alert) bool {
			if !alert.IsActive {
				return true
			}
			outOfMarket := len(alert.FilterMarkets) != 0
			for _, m := range alert.FilterMarkets {
				for mn, _ := range m.Markets {
					if imarket[mn] {
						outOfMarket = !m.IncludeOnly
						break
					}
				}
				if outOfMarket {
					break
				}
			}
			if outOfMarket {
				return true
			}

			alertFilter := []bool{}
			checkF := func(v float64, vv float64, filterType int) {
				if vv > v {
					alertFilter = append(alertFilter, true)
				}
			}
			for _, f := range alert.Filters {
				switch f.Type {
				case wings.FlowCapitalization:
					checkF(f.Raw.Value, pulse.MarketCap.CapByLV(f.Raw.Lv), f.Type)
				case wings.FlowPrice:
					checkF(f.Raw.Value, pulse.PriceByLV(f.Raw.Lv), f.Type)
				case wings.FlowVolume:
					checkF(f.Raw.Value, pulse.VolByLV(f.Raw.Lv), f.Type)
				}
			}
			if len(alertFilter) != len(alert.Filters) {
				return true
			}

			alertValues := []*wings.AlertValue{}

			for _, rule := range alert.Rules {
				tV, tD := targetValue(pulse, &dconst, rule.Raw.Lv, rule.Si, rule.Period)
				if tV == -1 && tD == -1 {
					continue
				}
				pc := tool.PercentageChange(tD, tV)
				xV := rule.Raw.Value
				av := &wings.AlertValue{
					Value:  pc,
					RuleId: rule.Raw.Id,
				}
				switch rule.Operation {
				case wings.Above:
					if pc < xV {
						alertValues = append(alertValues, av)
					}
				case wings.Greater:
					if pc > xV {
						alertValues = append(alertValues, av)
					}
				case wings.Corridor:
					if math.Abs(pc) < xV {
						alertValues = append(alertValues, av)
					}
				}
			}
			if len(alertValues) == len(alert.Rules) {
				updateAlert(alertValues, &alert, pulse, dconst)
			}
			return true
		})
	}
}

func targetValue(pulse *sol.CoinPulse, dialy *sol.DailyConst, lv string, s int, period int) (float64, float64) {
	if dialy.Tails == nil {
		return -1, -1
	}
	if len(dialy.Tails) <= period+1 {
		return -1, -1
	}
	if dialy.Tails[period][lv] == nil {
		return -1, -1
	}
	d := dialy.Tails[period][lv][s]
	switch s {
	case 100:
		return pulse.CapByLV(lv), d
	case 10, 11, 12, 13:
		return 0, d
	case 0, 1, 2, 3, 4, 14, 15, 16, 17, 18, 19:
		return pulse.PriceByLV(lv), d
	}
	return pulse.VolByLV(lv), d
}
