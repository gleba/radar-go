package sol

import (
	"time"
)

type PriceVol struct {
	VolumeUSD float64 `db:"VolumeUSD" json:"volume_usd"`
	VolumeBTC float64 `db:"VolumeBTC" json:"volume_btc"`
	PriceUSD  float64 `db:"PriceUSD" json:"price_usd"`
	PriceBTC  float64 `db:"PriceBTC" json:"price_btc"`
}

func (self *PriceVol) PriceByLV(lv string) float64 {
	if lv == "USD" {
		return self.PriceUSD
	} else {
		return self.PriceBTC
	}
}

func (self *PriceVol) VolByLV(lv string) float64 {
	if lv == "USD" {
		return self.VolumeUSD
	} else {
		return self.VolumeBTC
	}
}

type MarketCap struct {
	MarketCapUSD float64 `db:"MarketCapUSD" json:"market_cap_usd"`
	MarketCapBTC float64 `db:"MarketCapBTC" json:"market_cap_btc"`
}

func (self *MarketCap) CapByLV(lv string) float64 {
	if lv == "USD" {
		return self.MarketCapUSD
	} else {
		return self.MarketCapBTC
	}
}

type CoinPulse struct {
	PriceVol
	MarketCap
	ID   uint32    `db:"ID" json:"id"`
	Time time.Time `db:"Time" json:"time"`
}
