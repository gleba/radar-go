package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/syndtr/goleveldb/leveldb"
	"radar.cash/core/intel/service"
	"time"
)


var db *leveldb.DB

func init() {
	//var err error
	//db, err = leveldb.OpenFile("db", nil)
	//hand.Safe(err)
	service.OpenNATS()
	//core.On(ns.IntelPut, onPutMessage)
}


func main() {
	//service.Nats.Subscribe("put.sis", func(msg *nats.Msg) {
	//	fmt.Println(msg.Subject)
	//})
	//service.Nats.Subscribe("put.>", func(msg *nats.Msg) {
	//	fmt.Println(msg.Subject)
	//})
	service.Nats.Subscribe("xxx", func(msg *nats.Msg) {
		fmt.Println(msg)
		fmt.Println(msg)
	})

	service.Nats.Publish("xxx", []byte("sd"))
	fmt.Println("?")
	//var t = TT{
	//	Ok:"s",
	//	On:2,
	//}
	//core.Share("test", &t)
	//
	//b, _ := msgpack.Marshal(t)
	//fmt.Println(string(b))
	//fmt.Println("b is", string(b))
	//var v  interface{}
	//msgpack.Unmarshal(b, &v)
	//fmt.Println(v)
	time.Sleep(time.Duration(1000000)*time.Hour)
}

