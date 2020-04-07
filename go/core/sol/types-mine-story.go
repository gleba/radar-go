package sol

import "time"

type PriceDay = struct {
	Open      float64   `json:"open"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	Close     float64   `json:"close"`
	Volume    float64   `json:"volume"`
	MarketCap float64   `json:"market_cap"`
	Timestamp time.Time `json:"timestamp"`
}
type CmcStoryQuery struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
	} `json:"status"`
	Data StoryCoin `json:"data"`
}

type Quotes = struct {
	TimeOpen  time.Time `json:"time_open"`
	TimeClose time.Time `json:"time_close"`
	TimeHigh  time.Time `json:"time_high"`
	TimeLow   time.Time `json:"time_low"`
	Quote     struct {
		BTC PriceDay `json:"BTC"`
		USD PriceDay `json:"USD"`
	} `json:"quote"`
}
type StoryCoin = struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Quotes []Quotes `json:"quotes"`
}
