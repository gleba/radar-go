package heat

import (
	"log"
	"radar.cash/core/data/service"
	"radar.cash/core/heat/box"
	"radar.cash/core/sol"
)

var Pulses box.Pulses

func RestorePulse() []*sol.CoinPulse {
	pulse := sol.CoinPulse{}
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
		err := rows.StructScan(&pulse)
		if err != nil {
			log.Fatalln(err)
		}
		pulses = append(pulses, &pulse)
		Pulses.Store(pulse.ID, pulse)
	}
	return pulses
}
