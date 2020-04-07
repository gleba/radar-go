package sol

import (
	"database/sql"
	"fmt"

	"log"
	"radar.cash/core/hand"
	"radar.cash/core/intel"
	"radar.cash/core/intel/service"
	"time"
)

type PriceVol struct {
	VolumeUSD float64 `db:"VolumeUSD"`
	VolumeBTC float64 `db:"VolumeBTC"`
	PriceUSD  float64 `db:"PriceUSD"`
	PriceBTC  float64 `db:"PriceBTC"`
}
type MarketCap struct {
	MarketCapUSD float64 `db:"MarketCapUSD"`
	MarketCapBTC float64 `db:"MarketCapBTC"`
}



type CoinPulse struct {
	PriceVol
	MarketCap
	ID   uint32    `db:"ID" json:"id"`
	Time time.Time `db:"Time"`
}


type PulseWriter struct {
	tx        *sql.Tx
	stmt      *sql.Stmt
	hasUpdate bool
	Count     int
	queue []CoinPulse
	qcc []CCoin
}

func CreatePulseWriter() PulseWriter {
	fmt.Printf("☉")
	writer := PulseWriter{
		Count: 0,
	}
	var err error
	writer.tx, err = service.SqlX.Begin()
	hand.Safe(err)
	writer.stmt, err = writer.tx.Prepare("INSERT INTO CmcPulse (ID, Time, VolumeUSD, VolumeBTC, MarketCapUSD, MarketCapBTC,PriceUSD,PriceBTC) VALUES (?,?,?,?,?,?,?,?)")
	hand.Safe(err)
	writer.Commit()
	return writer
}

func (self *PulseWriter) Add(pulse CoinPulse, coin CCoin) {
	_, err := self.stmt.Exec(
		pulse.ID, pulse.Time,
		pulse.VolumeUSD, pulse.VolumeBTC, pulse.MarketCapUSD, pulse.MarketCapBTC, pulse.PriceUSD, pulse.PriceBTC)
	hand.Safe(err)
	self.Count = 1 + self.Count
	self.queue = append(self.queue, pulse)
	self.qcc = append(self.qcc, coin)
}

func (w *PulseWriter) Commit() {
	if w.Count >= 1 {
		hand.Safe(w.tx.Commit())
		fmt.Print("\n")
		log.Println("pulse change ", w.Count, "elements")
		w.Count = 0
		intel.Publish("flow.pulse.updates", w.queue)
		intel.Publish("flow.coin.updates", w.qcc)
		w.queue = w.queue[:0]
	}
}