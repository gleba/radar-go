package pegas

import (
	"fmt"
	"radar.cash/core/data/service"
)

var Alerts map[int]*Alert

func SyncPresets() {
	fmt.Println("SyncPresets start")
	period := []RawEnum{}
	mapPeriod := map[string]int{}
	service.DB.Query(&period, `select * from period;`)
	for _, enum := range period {
		mapPeriod[enum.Uuid] = enum.Index
	}

	mapCoeff := map[string]int{}
	coefficients := []RawEnum{}
	service.DB.Query(&coefficients, `select * from coefficients;`)
	for _, coefficient := range coefficients {
		mapCoeff[coefficient.Uuid] = coefficient.Index
	}

	mapOperations := map[string]int{}
	operations := []RawOperation{}
	service.DB.Query(&operations, `select * from operations;`)
	for _, operation := range operations {
		mapOperations[operation.Uuid] = operation.Code
	}

	rawFlows := []RawFlowRecord{}
	service.DB.Query(&rawFlows, `select * from flow;`)
	mapFlows := map[string]int{}
	for _, flow := range rawFlows {
		mapFlows[flow.Uuid] = flow.Code
	}

	rawFilters := []RawFilter{}
	service.DB.Query(&rawFilters, `select * from filters;`)

	mapFilters := map[int]*FlowFilter{}
	for _, r := range rawFilters {
		mapFilters[r.Id] = &FlowFilter{
			Type: mapFlows[r.Flow],
			Raw:  r,
		}
	}

	filterMarkets := []RawMarketFilter{}
	service.DB.Query(&filterMarkets, `select * from filter_market;`)

	mapFilterMarkets := map[int]*MarketFiler{}
	for _, market := range filterMarkets {
		mapFilterMarkets[market.Id] = &MarketFiler{
			Id:      market.Id,
			Markets: map[string]bool{},
			Exclude: market.Exclude,
		}
	}

	filterMarketMarkets := []RawFilterMarketMarkets{}
	service.DB.Query(&filterMarketMarkets, `select * from filter_market_markets;`)
	for _, m := range filterMarketMarkets {
		mapFilterMarkets[m.FilterMarketId].Markets[m.MarketsSlug] = true
	}

	rules := []RawRule{}
	service.DB.Query(&rules, `select * from rules;`)

	mapRule := map[int]*Rule{}
	for _, rawRule := range rules {
		var o Operation
		switch mapOperations[rawRule.Operation] {
		case 0:
			o = Above
		case 1:
			o = Greater
		case 3:
			o = Corridor
		}
		rule := &Rule{
			Raw:       rawRule,
			Si:        mapCoeff[rawRule.Coefficient],
			Period:    mapPeriod[rawRule.Period],
			Operation: o,
		}
		mapRule[rawRule.Id] = rule
	}

	Alerts = map[int]*Alert{}
	alerts := []RawAlert{}
	service.DB.Query(&alerts, `select * from alerts where status = 'published'`)
	for _, alert := range alerts {
		Alerts[alert.Id] = &Alert{
			Raw:           alert,
			Rules:         []*Rule{},
			Filters:       []*FlowFilter{},
			FilterMarkets: []*MarketFiler{},
		}
	}

	alertRules := []RawAlertRules{}
	service.DB.Query(&alertRules, `select * from alerts_rules`)

	for _, ar := range alertRules {
		alert := Alerts[ar.AlertsId]
		if alert == nil {
			continue
		}
		rule := mapRule[ar.RulesId]
		alert.Rules = append(alert.Rules, rule)
	}

	alertsFilters := []RawAlertFilters{}
	service.DB.Query(&alertsFilters, `select * from alerts_filters;`)
	for _, af := range alertsFilters {
		alert := Alerts[af.AlertsId]
		if alert == nil {
			continue
		}
		alert.Filters = append(alert.Filters, mapFilters[af.FiltersId])
	}

	alertsFilterMarkets := []RawAlertFilterMarket{}
	service.DB.Query(&alertsFilterMarkets, `select * from alerts_filter_market;`)
	for _, af := range alertsFilterMarkets {
		alert := Alerts[af.AlertsId]
		if alert == nil {
			continue
		}
		m := mapFilterMarkets[af.FilterMarketId]
		fmt.Println(m)
		alert.FilterMarkets = append(alert.FilterMarkets, m)
	}
	fmt.Println("SyncPresets done")
}
