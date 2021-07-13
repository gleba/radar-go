package pegas

import (
	"fmt"
	"math"
	"radar.cash/core/heat"
	"radar.cash/core/sol"
	"radar.cash/core/tool"
)

func AcceptPulse(pulses []*sol.CoinPulse) {
	for _, pulse := range pulses {
		dconst, _ := heat.DailyConst.Load(pulse.ID)
		//fmt.Println(dconst)
		imarket := IndexMarket[pulse.ID]
		for _, alert := range Alerts {
			inMarket := len(alert.FilterMarkets) == 0
			for _, m := range alert.FilterMarkets {
				for mn, _ := range m.Markets {
					if imarket[mn] {
						inMarket = true
						break
					}
				}
				if inMarket {
					break
				}
			}
			if !inMarket {
				continue
			}


			inFilter := []bool{}
			checkF := func(v float64, vv float64){
				if vv > v {
					inFilter = append(inFilter, true)
				}
			}
			for _, f := range alert.Filters {
				switch f.Type {
				case FlowCapitalization:
					checkF(f.Raw.Value, pulse.MarketCap.CapByLV(f.Raw.Lv))
				case FlowPrice:
					checkF(f.Raw.Value, pulse.PriceByLV(f.Raw.Lv))
				case FlowVolume:
					checkF(f.Raw.Value, pulse.VolByLV(f.Raw.Lv))
				}
			}
			if len(inFilter) != len(alert.Filters) {
				continue
			}

			fine := []bool{}

			for _, rule := range alert.Rules {
				tV , tD := targetValue(pulse, &dconst, rule.Raw.Lv, rule.Si, rule.Period)
				if tV == -1 && tD == -1 {
					continue
				}
				pc := tool.PercentageChange(tD, tV)
				xV := rule.Raw.Value
				//q, _ := heat.Quotas.Load(pulse.ID)
				//fmt.Println(rule.Raw)
				//fmt.Println(pc)
				//fmt.Println("https://coinmarketcap.com/currencies/" +q.Slug)
				switch rule.Operation {
				case Above:
					if pc < xV {
						fine = append(fine, true)
					}
				case Greater:
					if pc > xV {
						fine = append(fine, true)
					}
				case Corridor:
					if math.Abs(pc) < xV {
						fine = append(fine, true)
					}
				}
			}
			isActive := len(fine) == len(alert.Rules)
			if isActive {
				q, _ := heat.Quotas.Load(pulse.ID)
				fmt.Println("https://coinmarketcap.com/currencies/" +q.Slug)
			}
		}
	}
}

func targetValue(pulse *sol.CoinPulse, dialy *sol.DailyConst, lv string, s int, period int) (float64, float64) {
	if dialy.Tails == nil {
		return -1, -1
	}
	if len(dialy.Tails) <= period+1 {
		return -1, -1
	}
	if dialy. [period][lv] == nil {
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
