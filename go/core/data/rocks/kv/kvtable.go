package kv

import (
	"encoding/json"
	"fmt"
	"radar.cash/core/data/service"
	"radar.cash/core/hand"
	"time"
)

type RocksKv struct {
	NameSpace string
	queue     chan RockRaw
}
type RockRaw struct {
	Id   uint32 `db:"id"`
	Data []byte `db:"data"`
}

func MakeRocks(name string) *RocksKv {
	r := RocksKv{NameSpace: name}
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
	id UInt32,
	data String
) ENGINE=EmbeddedRocksDB
PRIMARY KEY id;`, name)
	_, err := service.SqlX.Exec(q)
	hand.Safe(err)
	r.queue = make(chan RockRaw, 10000)
	writers.Store(name, r)
	return &r
}

func (t RocksKv) Load(id uint32, i interface{}) {
	q := fmt.Sprintf(`SELECT data from %s where id == ?`, t.NameSpace)
	err := service.SqlX.Select(i, q, id)
	hand.Safe(err)
}

func (t RocksKv) LoadAll() []*RockRaw {
	var rr []*RockRaw
	q := fmt.Sprintf(`SELECT * from %s`, t.NameSpace)
	err := service.SqlX.Select(&rr, q)
	hand.Safe(err)
	return rr
}
func (t RocksKv) Store(id uint32, data interface{}) {
	bytes, _ := json.Marshal(data)
	if len(t.queue) > 5000 {
		time.Sleep(time.Duration(5) * time.Second)
	}
	t.queue <- RockRaw{
		Id:   id,
		Data: bytes,
	}
}

//var writers []*RocksKv
var writers SyncRocksKv
var writeTicker = time.NewTicker(time.Second * 1)

func init() {
	go weather()
}

func weather() {
	for {
		select {
		case _ = <-writeTicker.C:
			writers.Range(func(key string, writer RocksKv) bool {
				count := len(writer.queue)
				if count > 0 {
					q := fmt.Sprintf(`INSERT INTO %s (id, data) VALUES (?, ?)`, writer.NameSpace)
					tx, err := service.SqlX.Begin()
					hand.Safe(err)
					stmt, err := tx.Prepare(q)
					hand.Safe(err)
					for i := 0; i < count; i++ {
						r := <-writer.queue
						_, err = stmt.Exec(r.Id, r.Data)
						hand.Safe(err)
					}
					hand.Safe(tx.Commit())
					tx = nil
					stmt = nil
					fmt.Println(writer.NameSpace, "+", count)
				}
				return true
			})
		}
	}
}
