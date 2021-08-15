package pegas

import (
	"fmt"
	"radar.cash/core/data/service"
	"radar.cash/core/sol"
	"radar.cash/pharos/pegas/magic"
	"radar.cash/pharos/pegas/wings"
	"radar.cash/pharos/vars"
	"time"
)

func RestoreAlerts() {
	aa := []wings.ActiveAlert{}
	service.DB.Query(&aa, `SELECT * FROM active_alerts s1
WHERE time_last = (SELECT MAX(time_last) FROM active_alerts s2 WHERE s1.coin_id = s2.coin_id)
ORDER BY coin_id, time_last`)
	for _, alert := range aa {
		vars.ActiveAlerts.Store(alert.CoinId, alert)
	}
}

func WriteAlerts() {
	insert := []wings.ActiveAlert{}
	update := []wings.ActiveAlert{}
	all := []*wings.ActiveAlert{}
	vars.ActiveAlerts.Range(func(key uint32, aa wings.ActiveAlert) bool {

		if aa.IsNew {
			aa.IsNew = false
			aa.MarkToWrite = false
			insert = append(insert, aa)
			fmt.Println("insert", aa.CoinId)
			all = append(all, &aa)
			vars.ActiveAlerts.Store(key, aa)
		} else if aa.MarkToWrite {
			aa.MarkToWrite = false
			update = append(update, aa)
			all = append(all, &aa)
			vars.ActiveAlerts.Store(key, aa)
		}
		return true
	})
	if len(all) > 0 {
		magic.BroadCastFrontImpulses(all)
	}
	if len(insert) > 0 {
		r1, err := service.DB.Model(&insert).Insert()
		if err != nil {
			return
		}
		fmt.Println("insert", r1.RowsAffected())
	}
	if len(update) > 0 {
		r2, err := service.DB.Model(&update).Update()
		if err != nil {
			return
		}
		fmt.Println("update", r2.RowsAffected())
	}

}
func makeAA(pulse *sol.CoinPulse) wings.ActiveAlert {
	return wings.ActiveAlert{
		CoinId:      pulse.ID,
		Alerts:      map[int]map[int]*wings.AlertActiveRuleValue{},
		PulseTime:   pulse.Time,
		TimeFirst:   time.Now().Unix(),
		TimeLast:    time.Now().Unix(),
		IsNew:       true,
		MarkToWrite: true,
	}
}

func updateAlert(nowValues []*wings.AlertValue, alert *wings.Alert, pulse *sol.CoinPulse, dailyConst sol.DailyConst) {
	//isNew := false
	nowAlert, found := vars.ActiveAlerts.Load(pulse.ID)
	if !found {
		nowAlert = makeAA(pulse)
	} else {
		if nowAlert.PulseTime.Unix() == pulse.Time.Unix() {
			fmt.Println("same pulse")
			return
		}
	}

	timeNow := time.Now()
	unixTimeNow := timeNow.Unix()
	alertTimeLast := time.Unix(nowAlert.TimeLast, 0)
	diff := alertTimeLast.Sub(timeNow)

	if diff.Hours() > 24 {
		nowAlert = makeAA(pulse)
	}

	//if nowAlert.Alerts == nil {
	//	return
	//}
	alerts := nowAlert.Alerts[alert.Raw.Id]
	if alerts == nil {
		alerts = map[int]*wings.AlertActiveRuleValue{}
	}

	for _, v := range nowValues {
		a := alerts[v.RuleId]
		if a == nil {
			alerts[v.RuleId] = &wings.AlertActiveRuleValue{
				Value:     v.Value,
				StartTime: unixTimeNow,
				PeakValue: v.Value,
				PeakTime:  unixTimeNow,
			}

		} else {
			if v.Value > a.PeakValue {
				a.PeakValue = v.Value
				a.PeakTime = unixTimeNow
			}
			//if a.Value != v.Value {
			a.Value = v.Value
			a.UpdateTime = unixTimeNow
			//}
		}
	}

	nowAlert.Alerts[alert.Raw.Id] = alerts
	nowAlert.TimeLast = unixTimeNow
	nowAlert.MarkToWrite = true
	nowAlert.PulseTime = pulse.Time
	vars.ActiveAlerts.Store(pulse.ID, nowAlert)
}
