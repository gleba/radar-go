package pegas

import (
	"fmt"
	"radar.cash/core/data/service"
	"radar.cash/pharos/pegas/wings"
	"radar.cash/pharos/vars"
)

func SyncPresets() {
	fmt.Println("SyncPresets start")
	period := []wings.RawEnum{}
	mapPeriod := map[string]int{}
	service.DB.Query(&period, `select * from period;`)
	for _, enum := range period {
		mapPeriod[enum.Uuid] = enum.Index
	}

	mapCoeff := map[string]int{}
	coefficients := []wings.RawEnum{}
	service.DB.Query(&coefficients, `select * from coefficients;`)
	for _, coefficient := range coefficients {
		mapCoeff[coefficient.Uuid] = coefficient.Index
	}

	mapOperations := map[string]int{}
	mapOperationsInstances := map[string]wings.RawOperation{}
	operations := []wings.RawOperation{}
	service.DB.Query(&operations, `select * from operations;`)
	for _, operation := range operations {
		mapOperations[operation.Uuid] = operation.Code
		mapOperationsInstances[operation.Uuid] = operation
	}

	rawFlows := []wings.RawFlowRecord{}
	service.DB.Query(&rawFlows, `select * from flow;`)
	mapFlows := map[string]int{}
	for _, flow := range rawFlows {
		mapFlows[flow.Uuid] = flow.Code
	}

	rawFilters := []wings.RawFilter{}
	service.DB.Query(&rawFilters, `select * from filters;`)

	mapFilters := map[int]*wings.FlowFilter{}
	for _, r := range rawFilters {
		mapFilters[r.Id] = &wings.FlowFilter{
			Type: mapFlows[r.Flow],
			Raw:  r,
		}
	}

	filterMarkets := []wings.RawMarketFilter{}
	service.DB.Query(&filterMarkets, `select * from filter_market;`)

	mapFilterMarkets := map[int]*wings.MarketFiler{}
	for _, market := range filterMarkets {
		mapFilterMarkets[market.Id] = &wings.MarketFiler{
			Id:          market.Id,
			Markets:     map[string]bool{},
			IncludeOnly: market.IncludeOnly,
		}
	}

	filterMarketMarkets := []wings.RawFilterMarketMarkets{}
	service.DB.Query(&filterMarketMarkets, `select * from filter_market_markets;`)
	for _, m := range filterMarketMarkets {
		mapFilterMarkets[m.FilterMarketId].Markets[m.MarketsSlug] = true
	}

	rules := []wings.RawRule{}
	service.DB.Query(&rules, `select * from rules;`)

	mapRule := map[int]*wings.Rule{}
	for _, rawRule := range rules {
		var o wings.Operation
		switch mapOperations[rawRule.Operation] {
		case 0:
			o = wings.Above
		case 1:
			o = wings.Greater
		case 3:
			o = wings.Corridor
		}
		rule := &wings.Rule{
			Raw:       rawRule,
			Si:        mapCoeff[rawRule.Coefficient],
			Period:    mapPeriod[rawRule.Period],
			Operation: o,
		}
		mapRule[rawRule.Id] = rule
		op := mapOperationsInstances[rawRule.Operation]
		vars.FrontDict.Rules[rawRule.Id] = &wings.FrontRule{
			Label:     rule.Raw.Label,
			Code:      op.Code,
			Operation: op.Label,
		}
	}

	//vars.Alerts = map[int]*wings.Alert{}
	alerts := []wings.RawAlert{}
	service.DB.Query(&alerts, `select * from alerts`)
	for _, alert := range alerts {
		vars.Alerts.Store(alert.Id, wings.Alert{
			Raw:           alert,
			IsActive:      alert.Status == "published",
			Rules:         []*wings.Rule{},
			Filters:       []*wings.FlowFilter{},
			FilterMarkets: []*wings.MarketFiler{},
		})
		vars.FrontDict.Alert[alert.Id] = &wings.FrontAlert{
			Label:    alert.Label,
			IsActive: alert.Status == "published",
			Rules:    []int{},
		}
	}

	alertRules := []wings.RawAlertRules{}
	service.DB.Query(&alertRules, `select * from alerts_rules`)

	for _, ar := range alertRules {
		alert, found := vars.Alerts.Load(ar.AlertsId)
		if !found {
			continue
		}
		rule := mapRule[ar.RulesId]
		alert.Rules = append(alert.Rules, rule)
		vars.FrontDict.Alert[ar.AlertsId].Rules = append(vars.FrontDict.Alert[ar.AlertsId].Rules, rule.Raw.Id)
		vars.Alerts.Store(alert.Raw.Id, alert)
	}

	alertsFilters := []wings.RawAlertFilters{}
	service.DB.Query(&alertsFilters, `select * from alerts_filters;`)
	for _, af := range alertsFilters {
		alert, found := vars.Alerts.Load(af.AlertsId)
		if !found {
			continue
		}
		alert.Filters = append(alert.Filters, mapFilters[af.FiltersId])
		vars.Alerts.Store(alert.Raw.Id, alert)
	}

	alertsFilterMarkets := []wings.RawAlertFilterMarket{}
	service.DB.Query(&alertsFilterMarkets, `select * from alerts_filter_market;`)
	for _, af := range alertsFilterMarkets {
		alert, found := vars.Alerts.Load(af.AlertsId)
		if !found {
			continue
		}
		m := mapFilterMarkets[af.FilterMarketId]
		alert.FilterMarkets = append(alert.FilterMarkets, m)
		vars.Alerts.Store(alert.Raw.Id, alert)
	}
	fmt.Println("SyncPresets done")
}
