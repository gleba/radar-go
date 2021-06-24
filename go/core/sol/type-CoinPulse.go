package sol

import (
	"time"
)

type PriceVol struct {
	VolumeUSD float64 `db:"VolumeUSD"`
	VolumeBTC float64 `db:"VolumeBTC"`
	PriceUSD  float64 `db:"PriceUSD"`
	PriceBTC  float64 `db:"PriceBTC"`
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
	MarketCapUSD float64 `db:"MarketCapUSD"`
	MarketCapBTC float64 `db:"MarketCapBTC"`
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
	Time time.Time `db:"Time"`
}
