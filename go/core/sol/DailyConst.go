package sol

type DailyConst struct {
	ID    uint32 `json:"id"`
	Fine  bool   `json:"fine"`
	Day   string `json:"day"`
	Tails Tails  `json:"tails"`
}

type TailCoinConst []float64
type MultiTail map[string][]float64
type Tails []MultiTail

func (self Tails) Qerry(period int, lv string, k int) (have bool, kValue float64) {
	if len(self) > period {
		p := self[period]
		if p != nil {
			v, have := p[lv]
			if have {
				return have, v[k]
			}
		}
	}
	return false, 0
}
