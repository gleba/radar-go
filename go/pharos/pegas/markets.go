package pegas

import (
	"encoding/json"
	"fmt"
	"radar.cash/core/data/rocks"
	"radar.cash/core/data/service"
	"radar.cash/core/hand"
	"radar.cash/core/sol"
	"radar.cash/pharos/models"
)

var IndexMarket map[uint32]map[string]bool

func SyncMarkets() {
	fmt.Println("SyncMarkets start")
	rr := rocks.Markets.LoadAll()
	IndexMarket = map[uint32]map[string]bool{}
	marketKeys := map[string]string{}
	marketPSK := map[string]bool{}
	for _, r := range rr {
		am := []*sol.Market{}
		hand.Safe(json.Unmarshal(r.Data, &am))
		ex := map[string]bool{}
		for _, market := range am {
			ex[market.Slug] = true
			marketKeys[market.Slug] = market.Name
		}
		IndexMarket[r.Id] = ex
	}

	markets := []*models.Market{}
	err := service.DB.Model(&markets).Select(&markets)
	if err != nil {
		fmt.Println("+")
	}
	//fmt.Println(markets)
	for _, market := range markets {
		marketPSK[market.Slug] = true
	}
	for key, name := range marketKeys {
		fmt.Println(key, name)
		_, found := marketPSK[key]
		if !found {
			mmm := &models.Market{
				Slug: key,
				Name: name,
			}
			insert, err := service.DB.Model(mmm).Insert()
			fmt.Println(insert, err)
			if err != nil {
				return
			}
		}
	}
	fmt.Println("SyncMarkets end")
}
