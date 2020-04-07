package scifi

import (
	"radar.cash/core/sol"
	"radar.cash/core/sol/si"
)

func insides(id uint32, multiTails []sol.MultiTail) {
	//for period, tail := range multiTails {
	//	for LV, con := range tail {
	//		for k, v := range con {
	//
	//		}
	//	}
	//}
}

var tactic = []Tactic{{
	LV:     BTC,
	K:      si.VolumeMean,
	Period: si.Period30,
	Change: 700,
},{
	LV:     BTC,
	K:      si.VolumeMean,
	Period: si.Period60,
	Change: 500,
},{
	LV:     BTC,
	K:      si.VolumeMean,
	Period: si.Period90,
	Change: 300,
}}

type Tactic struct {
	Period int
	LV     string
	K      int
	Change float64
}
