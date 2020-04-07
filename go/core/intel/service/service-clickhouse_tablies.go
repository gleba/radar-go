package service

import "fmt"

const (
	TableCmcPulse    = "CmcPulse"
	TableCmcDaysPulse = "CmcDaysPulse"
)

var tablePulseVol = `	
	VolumeUSD  Float64,
	VolumeBTC  Float64,
	PriceUSD   Float64,
	PriceBTC   Float64,`

var tableMarketCap = `
	MarketCapUSD Float64,
	MarketCapBTC Float64,`

var tablies = []string{
	fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	%s
	%s	
	ID      UInt32,
	Time  	  	DateTime
)
ENGINE=MergeTree()
ORDER BY (Time)`, TableCmcPulse, tablePulseVol, tableMarketCap),

fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	Open      Float64,
	Close     Float64,
	High      Float64,
	Low       Float64,
	Volume    Float64,
	MarketCap Float64,
	ID    UInt32,
	Date  		Date
)
ENGINE=MergeTree()
ORDER BY (ID, Date)
`, TableCmcDaysPulse),
}
