package cmc

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"radar.cash/core/data/eSpace"
	"radar.cash/core/data/service"
	"radar.cash/core/hand"
	"radar.cash/core/sol"
	"radar.cash/core/tool"
)

type queue struct {
	tx   *sql.Tx
	stmt *sql.Stmt
}

func (q queue) JustClose() {
	hand.Safe(q.stmt.Close())
	q.tx = nil
	q.stmt = nil

}
func (q queue) CommitAndClose() {
	hand.Safe(q.tx.Commit())
	q.JustClose()
}

type PulseWriter struct {
	queueCoinPulse []*sol.CoinPulse
	queueQuotes    []*sol.CoinQuote
}

func CreatePulseWriter() PulseWriter {
	//fmt.Printf("☉")
	writer := PulseWriter{
		queueCoinPulse: []*sol.CoinPulse{},
		queueQuotes:    []*sol.CoinQuote{},
	}
	return writer
}

func (w *PulseWriter) Add(pulse *sol.CoinPulse, coin *sol.CoinQuote) {
	w.queueCoinPulse = append(w.queueCoinPulse, pulse)
	w.queueQuotes = append(w.queueQuotes, coin)
}
func (w *PulseWriter) Commit() {
	if len(w.queueCoinPulse) > 0 {
		tx, _ := service.SqlX.Begin()
		stmt, _ := tx.Prepare("INSERT INTO CmcPulse (ID, Time, VolumeUSD, VolumeBTC, MarketCapUSD, MarketCapBTC,PriceUSD,PriceBTC) VALUES (?,?,?,?,?,?,?,?)")
		for _, pulse := range w.queueCoinPulse {
			_, _ = stmt.Exec(
				pulse.ID, pulse.Time,
				pulse.VolumeUSD, pulse.VolumeBTC, pulse.MarketCapUSD, pulse.MarketCapBTC, pulse.PriceUSD, pulse.PriceBTC)

		}
		hand.Safe(tx.Commit())
		eSpace.Pulses.Publish(w.queueCoinPulse)
	}
	if len(w.queueQuotes) > 0 {
		tx, _ := service.SqlX.Begin()
		stmt, _ := tx.Prepare("INSERT INTO CoinQuote (id, data) VALUES (?,?)")
		for _, quote := range w.queueQuotes {
			bytes, _ := json.Marshal(quote)
			_, _ = stmt.Exec(quote.ID, bytes)
		}
		hand.Safe(tx.Commit())

		fmt.Println("*", len(w.queueCoinPulse))
	}
	tool.MemState()
}
