package main

import (
	"radar.cash/core/data/rocks"
	"radar.cash/core/data/service"
	"radar.cash/miner/cmc"
	"radar.cash/miner/end"
	"time"
)

func init() {
	service.OpenClickHose()
	service.OpenNATS()
	rocks.Init()
}

func main() {

	end.WarmCaches()

	latestTicker := time.NewTicker(time.Second * 24)
	storyTicker := time.NewTicker(time.Hour * 1)

	cmc.MineLatest()
	go cmc.MineStory()

	for {
		select {
		//case _ = <-infoTicker.C:
		case _ = <-storyTicker.C:
			go cmc.MineStory()
		case _ = <-latestTicker.C:
			cmc.MineLatest()
		}
	}
}
