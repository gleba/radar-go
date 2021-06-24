package rocks

import "radar.cash/core/data/rocks/kv"

var Markets *kv.RocksKv
var DailyConst *kv.RocksKv
var CoinQuote *kv.RocksKv

func Init() {
	Markets = kv.MakeRocks("Markets")
	CoinQuote = kv.MakeRocks("CoinQuote")
	DailyConst = kv.MakeRocks("DailyConst")
}
