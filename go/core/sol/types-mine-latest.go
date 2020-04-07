package sol

import "time"

type CCoin struct {
	ID                uint32      `json:"id"`
	Name              string      `json:"name"`
	Symbol            string      `json:"symbol"`
	Slug              string      `json:"slug"`
	NumMarketPairs    int         `json:"num_market_pairs"`
	DateAdded         time.Time   `json:"date_added"`
	Tags              []string    `json:"tags"`
	MaxSupply         float64     `json:"max_supply"`
	CirculatingSupply float64     `json:"circulating_supply"`
	TotalSupply       float64     `json:"total_supply"`
	Platform          interface{} `json:"platform"`
	CmcRank           float64     `json:"cmc_rank"`
	LastUpdated       time.Time   `json:"last_updated"`
	Quote             struct {
		//BCH ListingPrice `json:"BCH"`
		BTC ListingPrice `json:"BTC"`
		//ETH ListingPrice `json:"ETH"`
		//LTC ListingPrice `json:"LTC"`
		USD ListingPrice `json:"USD"`
	} `json:"quote"`
}

type ListingPrice struct {
	Price            float64   `json:"price"`
	Volume24H        float64   `json:"volume_24h"`
	PercentChange1H  float64   `json:"percent_change_1h"`
	PercentChange24H float64   `json:"percent_change_24h"`
	PercentChange7D  float64   `json:"percent_change_7d"`
	MarketCap        float64   `json:"market_cap"`
	LastUpdated      time.Time `json:"last_updated"`
}

type CmcListingQuery struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
	} `json:"status"`
	Data []CCoin `json:"data"`
}
