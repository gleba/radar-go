package cmc

import (
	"encoding/json"
	"fmt"
	"log"

	//"github.com/syndtr/goleveldb/leveldb"
	"radar.cash/core/hand"
	"radar.cash/core/intel/df"
	"radar.cash/core/sol"
	"strconv"
	"time"
)

var storyPool chan sol.CCoin

//var ldb *leveldb.DB

func init() {
	storyPool = make(chan sol.CCoin)
	var err error
	hand.Safe(err)
	for range [3]int{} {
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
	latestSync.Range(func(key, coin interface{}) bool {
		storyPool <- coin.(sol.CCoin)
		return true
	})
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

func getStory(coin sol.CCoin) {
	lastDay, have := lastDaily.Load(coin.ID)
	if have || coin.Quote.BTC.Volume24H > 5 && coin.Quote.BTC.MarketCap > 100 {
		if lastDay != forDay {
			url := "https://web-api.coinmarketcap.com/v1/cryptocurrency/ohlcv/historical?convert=BTC,USD&slug=" + coin.Slug + "&time_end=" + timeEndStr + "&time_start=" + timeStartStr
			bytes := request(url)
			if bytes == nil {
				return
			}
			var storyQuery *sol.CmcStoryQuery
			hand.Safe(json.Unmarshal(bytes, &storyQuery))
			size := len(storyQuery.Data.Quotes)
			last := storyQuery.Data.Quotes[size-1]
			newDay := last.TimeClose.Format("20060102")
			lastDaily.Store(coin.ID, newDay)
			if size < 42 {
				return
			}
			fmt.Print("+")
			dailyConst := makeDailyConst(coin.ID, newDay, storyQuery)
			df.DailyConst.UpdateItem(coin.ID, dailyConst)
			marketPool <- coin
		}
	}
}
