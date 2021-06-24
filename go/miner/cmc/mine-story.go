package cmc

import (
	"encoding/json"
	"fmt"
	"log"
	"radar.cash/core/heat"
	"radar.cash/miner/end"

	//"github.com/syndtr/goleveldb/leveldb"
	"radar.cash/core/hand"
	"radar.cash/core/sol"
	"strconv"
	"time"
)

var storyPool chan *sol.CoinQuote

//var ldb *leveldb.DB

func init() {
	storyPool = make(chan *sol.CoinQuote)
	var err error
	hand.Safe(err)
	for range [24]int{} {
		go storyMinePool()
	}
}

var timeStartStr string
var timeEndStr string
var forDay string

func MineStory() {
	log.Println("MineStory")
	t := time.Now()
	timeNow := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	timeEnd := timeNow.Add(time.Hour * 24 * -1)
	forDay = timeEnd.Format("20060102")
	timeStartStr = strconv.FormatInt(timeNow.Add(time.Hour*24*91*-1).Unix(), 10)
	timeEndStr = strconv.FormatInt(timeEnd.Unix(), 10)
	fmt.Println(len(storyPool))
	if len(storyPool) < 5 {
		heat.Quotas.Range(func(key uint32, value sol.CoinQuote) bool {
			storyPool <- &value
			return true
		})
	}
}

func storyMinePool() {
	for {
		coin := <-storyPool
		if coin.ID != 1 {
			getStory(coin)
		}
	}
}

func storyAddPrice(a []float64, v sol.PriceDay) []float64 {
	if v.Close != 0 {
		return append(a, v.Open, v.Low, v.High, v.Close)
	}
	return a
}
func storyAdd(a []float64, v float64) []float64 {
	if v != 0 {
		return append(a, v)
	}
	return a
}

func getStory(coin *sol.CoinQuote) {
	dc, have := heat.DailyConst.Load(coin.ID)
	if !have || coin.Quote.BTC.Volume24H > 5 && coin.Quote.BTC.MarketCap > 100 {
		if dc.Day != forDay {
			url := "https://web-api.coinmarketcap.com/v1/cryptocurrency/ohlcv/historical?convert=BTC,USD&slug=" + coin.Slug + "&time_end=" + timeEndStr + "&time_start=" + timeStartStr
			bytes := request(url)
			if bytes == nil {
				return
			}
			var storyQuery sol.CmcStoryQuery
			hand.Safe(json.Unmarshal(bytes, &storyQuery))
			size := len(storyQuery.Data.Quotes)
			if size < 42 {
				return
			}
			last := storyQuery.Data.Quotes[size-1]
			newDay := last.TimeClose.Format("20060102")
			//fmt.Print("-")
			dailyConst := makeDailyConst(coin.ID, newDay, &storyQuery)
			end.UpdateDailyConst(&dailyConst)
			_, ok := end.Markets.Load(coin.ID)
			if !ok {
				marketPool <- coin
			}
		}
	}
}
