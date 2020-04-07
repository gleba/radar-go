package sol




type DailyConst struct {
	ID    uint32 `json:"id"`
	Fine  bool   `json:"-"`
	Day   string `json:"day"`
	Tails []MultiTail `json:"tails"`
}


type TailCoinConst []float64
type MultiTail map[string][]float64
