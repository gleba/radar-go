package wings

type Operation int

const (
	Greater  Operation = 1
	Above    Operation = 0
	Corridor Operation = 3
)
const (
	FlowPrice          int = 1
	FlowVolume         int = 0
	FlowCapitalization int = 3
)

type FlowFilter struct {
	Type int
	Raw  RawFilter
}
type RawFilter struct {
	Id    int     `json:"id,omitempty"`
	Lv    string  `json:"lv,omitempty"`
	Flow  string  `json:"flow,omitempty"`
	Value float64 `json:"value,omitempty"`
	Label string  `json:"label,omitempty"`
}
type RawFlowRecord struct {
	Uuid  string `json:"uuid,omitempty"`
	Label string `json:"label,omitempty"`
	Code  int    `json:"code,omitempty"`
}

type MarketFiler struct {
	Id          int
	Markets     map[string]bool
	IncludeOnly bool
}
type RawMarketFilter struct {
	Id    int    `json:"int"`
	Label string `json:"label,omitempty"`
	//Exclude     bool `json:"exclude,omitempty"`
	IncludeOnly bool `json:"include_only,omitempty"`
}

type CoinMarkets struct {
	CoinId  uint32
	Markets map[string]bool
}

type AlertActiveRuleValue struct {
	Value      float64 `json:"value,omitempty"`
	StartTime  int64   `json:"start_time,omitempty"`
	UpdateTime int64   `json:"update_time,omitempty"`
	PeakValue  float64 `json:"peak_value,omitempty"`
	PeakTime   int64   `json:"peak_time,omitempty"`
}

type AlertValue struct {
	Value  float64 `json:"value,omitempty"`
	RuleId int     `json:"rule_id,omitempty"`
}

type Alert struct {
	Raw           RawAlert ``
	IsActive      bool
	Rules         []*Rule
	Filters       []*FlowFilter
	FilterMarkets []*MarketFiler
}

type Rule struct {
	Raw       RawRule
	Si        int
	Period    int
	Operation Operation
}

type RawRule struct {
	Id          int     `json:"id,omitempty"`
	Coefficient string  `json:"coefficient,omitempty"`
	Value       float64 `json:"value,omitempty"`
	Operation   string  `json:"operation,omitempty"`
	Label       string  `json:"label"`
	Period      string  `json:"period"`
	Lv          string  `json:"lv"`
}

type RawEnum struct {
	Uuid  string `json:"uuid,omitempty"`
	Label string `json:"label,omitempty"`
	Index int    `json:"index,omitempty"`
}

type RawOperation struct {
	Uuid  string `json:"uuid,omitempty"`
	Label string `json:"label,omitempty"`
	Code  int    `json:"index,omitempty"`
}

type RawAlert struct {
	Id     int    `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	Label  string `json:"label,omitempty"`
}
type RawAlertRules struct {
	AlertsId int `json:"alerts_id,omitempty"`
	RulesId  int `json:"rules_id,omitempty"`
}

type RawAlertFilters struct {
	AlertsId  int `json:"alerts_id,omitempty"`
	FiltersId int `json:"filters_id,omitempty"`
}

type RawAlertFilterMarket struct {
	AlertsId       int `json:"alerts_id,omitempty"`
	FilterMarketId int `json:"filter_market_id,omitempty"`
}
type RawAlertFilterMarketMarkets struct {
	AlertsId       int `json:"alerts_id,omitempty"`
	FilterMarketId int `json:"filter_market_id,omitempty"`
}
type RawFilterMarketMarkets struct {
	FilterMarketId int    `json:"filter_market_id,omitempty"`
	MarketsSlug    string `json:"markets_slug,omitempty"`
}
