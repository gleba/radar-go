package main

import (
	"radar.cash/core/intel/service"
	"radar.cash/miner/cmc"
	"time"
)

func init() {
	service.OpenClickHose()
	service.OpenNATS()
}

func main() {
	cmc.WarmCaches()
	latestTicker := time.NewTicker(time.Second * 12)
	storyTicker := time.NewTicker(time.Hour * 1)
	//infoTicker := time.NewTicker(time.Hour * 360)
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
