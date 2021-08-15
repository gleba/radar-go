package wings

type FrontImpulse struct {
	Id   int       `json:"id,omitempty"`
	Coin FrontCoin `json:"coin"`
	//Pulse  sol.CoinPulse                     `json:"pulse,omitempty"`
	Alerts map[int]map[int]*AlertActiveRuleValue `json:"alerts"`
	//FrontAlerts []*AlertActiveRuleValue `json:"front_alerts"`
	TimeFirst int64 `json:"time_first"`
	TimeLast  int64 `json:"time_last"`
}

type FrontCoin struct {
	ID     uint32 `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Slug   string `json:"slug"`
}

type FrontDict struct {
	Rules map[int]*FrontRule  `json:"rules,omitempty"`
	Alert map[int]*FrontAlert `json:"alert,omitempty"`
}
type FrontAlert struct {
	Label    string `json:"label,omitempty"`
	Rules    []int  `json:"rules,omitempty"`
	IsActive bool   `json:"isActive"`
}
type FrontRule struct {
	Label     string `json:"label,omitempty"`
	Code      int    `json:"code"`
	Operation string `json:"operation"`
}

//type FrontRule struct {
//	LastValue  float64   `json:"last_value,omitempty"`
//	PeakValue  float64   `json:"peak_value,omitempty"`
//	PeakTime   time.Time `json:"peak_time"`
//	DetectTime time.Time `json:"detect_time"`
//	UpdateTime time.Time `json:"update_time"`
//}
