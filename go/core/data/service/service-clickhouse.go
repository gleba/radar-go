package service

import (
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"radar.cash/core/hand"
)

var SqlX *sqlx.DB

func OpenClickHose() {
	var err error
	//SqlX, err = sqlx.Open("clickhouse", "tcp://clickhouse:9000?username=default&password=a3dd9e844b7a43af4b0b5966016103f1d5ee97dd")
	SqlX, err = sqlx.Open("clickhouse", "tcp://"+os.Getenv("CLICKHOUSE")+":9000?username=default&password=a3dd9e844b7a43af4b0b5966016103f1d5ee97dd")
	hand.Safe(err)
	if err := SqlX.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			log.Println("ClickHouse: sql ", ok)
		}
		return
	}

	for _, table := range tablies {
		_, err = SqlX.Exec(table)
		hand.Safe(err)
	}
	log.Println("ClickHouse: ok")
}

//
//type ClickHouseWriter struct {
//	X     *sql.Tx
//	stmt  *sql.Stmt
//	xx    *sql.DB
//	Count int
//}
//
//func (chi *ClickHouseInstance) Writer() *ClickHouseWriter  {
//	writer := ClickHouseWriter{
//		Count: 0,
//	}
//	var err error
//	writer.X, err = chi.X.Begin()
//	SafeError(err)
//	return &writer
//}
//
//func (w *ClickHouseWriter) Statement(stmt *sql.Stmt, err error) {
//	if SafeError(err) {
//		w.stmt = stmt
//	}
//}
//
//func (w *ClickHouseWriter) Add(data ...interface{}) {
//	_, err := w.stmt.Exec(data...)
//	SafeError(err)
//	w.Count = 1 + w.Count
//}
//func (w *ClickHouseWriter) Commit() {
//	if w.Count >= 1 {
//		SafeError(w.X.Commit())
//		w.Count = 0
//	}
//}
