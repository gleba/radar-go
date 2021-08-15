package heat

import (
	"fmt"
	"log"
	"radar.cash/core/data/service"
	"radar.cash/core/heat/box"
	"radar.cash/core/sol"
)

var Pulses box.Pulses

func RestorePulse() []*sol.CoinPulse {
	fmt.Println("RestorePulse")

	var pulses []*sol.CoinPulse
	rows, err := service.SqlX.Queryx(`SELECT *
FROM CmcPulse
WHERE Time IN (SELECT Time
               FROM CmcPulse AS T2
               WHERE T2.ID = CmcPulse.ID
               ORDER BY Time DESC
               LIMIT 1)`)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		pulse := &sol.CoinPulse{}
		err := rows.StructScan(pulse)
		if err != nil {
			log.Fatalln(err)
		}
		pulses = append(pulses, pulse)
		Pulses.Store(pulse.ID, *pulse)
	}

	fmt.Println("RestorePulse done")
	//eSpace.Pulses.Sub(func(pulses []*sol.CoinPulse) {
	//	for _, pulse := range pulses {
	//		Pulses.Store(pulse.ID, *pulse)
	//	}
	//})
	return pulses
}
