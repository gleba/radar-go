package data

//
//import (
//	"encoding/json"
//	"log"
//	"radar.cash/core/data/service"
//	"reflect"
//	"time"
//)
//
//type PutMessage struct {
//	Key   string
//	Value interface{}
//}
//
////func Publish(channel string, v interface{}) {
////	//data, e := json.Marshal(v)
////	//if e != nil {
////	//	hand.Safe(e)
////	//}
////	//hand.Safe(e)
////	service.NatsEncoded.Publish(channel, v)
////}
//
////func Share(namespace string, v interface{}) {
////	service.NatsEncoded.Publish(namespace, v)
////	service.Nats.QueueSubscribe(namespace, "", func(msg *nats.Msg) {
////		b, err := msgpack.Marshal(&v)
////		hand.Safe(err)
////		msg.Respond(b)
////	})
////}
//
//type Handler interface{}
//
//func Subscribe(namespace string, cb Handler) {
//	service.NatsEncoded.Subscribe(namespace, cb)
//}
//
//
//
//
//func Request(path string, cb Handler) {
//	log.Println("req."+path)
//	msg, err := service.Nats.Request("req."+path, nil, time.Second*10)
//	if err != nil {
//		log.Println("timeout NATS request: req."+path)
//	} else if msg.Data != nil {
//		cbValue := reflect.ValueOf(cb)
//		aValue := cbValue.Type().In(0)
//		v := reflect.New(aValue)
//		e := v.Interface()
//		json.Unmarshal(msg.Data, e)
//		in := make([]reflect.Value, 1)
//		in[0] = v.Elem()
//		cbValue.Call(in)
//	}
//}
