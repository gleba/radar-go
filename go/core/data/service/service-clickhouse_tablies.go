package service

import "fmt"

const (
	TableCmcPulse = "CmcPulse"
	//TableCmcDaysPulse = "CmcDaysPulse"
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
	`CREATE TABLE IF NOT EXISTS markets (
	id UInt32,
	data BLOB
) ENGINE=EmbeddedRocksDB
PRIMARY KEY id;`,
	`CREATE TABLE IF NOT EXISTS quotes (
	id UInt32,
	data BLOB
) ENGINE=EmbeddedRocksDB
PRIMARY KEY id;`,
	`CREATE TABLE IF NOT EXISTS quotes (
	id UInt32,
	data BLOB
) ENGINE=EmbeddedRocksDB
PRIMARY KEY id;`,
}
