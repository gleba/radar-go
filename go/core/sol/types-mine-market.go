package sol

import "time"

type CoinMarkets struct {
	ID      uint32   `json:"id"`
	Markets []Market `json:"markets"`
}

type Exchange struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
type Market struct {
	Exchange
	Pairs []Pair `json:"pairs"`
}

type Pair struct {
	Pair string `json:"pair"`
	URL  string `json:"url"`
}

type MarketPair struct {
	Exchange        Exchange    `json:"exchange"`
	OutlierDetected int         `json:"outlier_detected"`
	Exclusions      interface{} `json:"exclusions"`
	MarketPairBase  struct {
		ExchangeSymbol string `json:"exchange_symbol"`
		CurrencyID     int    `json:"currency_id"`
		CurrencySymbol string `json:"currency_symbol"`
		CurrencyType   string `json:"currency_type"`
	} `json:"market_pair_base"`
	MarketPairQuote struct {
		ExchangeSymbol string `json:"exchange_symbol"`
		CurrencyID     int    `json:"currency_id"`
		CurrencySymbol string `json:"currency_symbol"`
		CurrencyType   string `json:"currency_type"`
	} `json:"market_pair_quote"`
	Quote struct {
		ExchangeReported struct {
			Price          float64   `json:"price"`
			Volume24HBase  float64   `json:"volume_24h_base"`
			Volume24HQuote float64   `json:"volume_24h_quote"`
			LastUpdated    time.Time `json:"last_updated"`
		} `json:"exchange_reported"`
		USD struct {
			Price       float64   `json:"price"`
			Volume24H   float64   `json:"volume_24h"`
			LastUpdated time.Time `json:"last_updated"`
		} `json:"USD"`
	} `json:"quote"`
	MarketID   int    `json:"market_id"`
	MarketPair string `json:"market_pair"`
	MarketURL  string `json:"market_url"`
}

type MarketData struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Symbol      string       `json:"symbol"`
	MarketPairs []MarketPair `json:"market_pairs"`
}

type MarketQuery struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
	} `json:"status"`
	Data MarketData `json:"data"`
}
