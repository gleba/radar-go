package wings

import "time"

type ActiveAlert struct {
	Id          int
	CoinId      uint32
	Alerts      map[int]map[int]*AlertActiveRuleValue
	TimeFirst   int64
	TimeLast    int64
	PulseTime   time.Time `pg:"-"`
	IsNew       bool      `pg:"-"`
	MarkToWrite bool      `pg:"-"`
}

type Market struct {
	Slug string
	Name string
}
