package service

//import (
//	"github.com/restream/reindexer"
//	"reflect"
//)
//
//var ReIndexer *reindexer.Reindexer
//
//func OpenReIndexer() {
//	ReIndexer = reindexer.NewReindex("cproto://localhost:6534/mydb")
//	ReIndexer.OpenNamespace(NsDailyConst, reindexer.DefaultNamespaceOptions(), DailyConst{})
//	//ReIndexer.OpenNamespace("story", reindexer.DefaultNamespaceOptions(), StoryCoin{})
//	//ReIndexer.OpenNamespace("cc", reindexer.DefaultNamespaceOptions(), CCoin{})
//}
//
//func AddReIndexerIndex(namespace string, s interface{}) {
//	ReIndexer.OpenNamespace(namespace, reindexer.DefaultNamespaceOptions(), s)
//}
//
//type Handler interface{}
//
//func argInfo(cb Handler) (reflect.Type, int) {
//	cbType := reflect.TypeOf(cb)
//	if cbType.Kind() != reflect.Func {
//		panic("nats: Handler needs to be a func")
//	}
//	numArgs := cbType.NumIn()
//	if numArgs == 0 {
//		return nil, numArgs
//	}
//	return cbType.In(numArgs - 1), numArgs
//}
//func GetAll(namespace string, cb Handler) {
//	query := ReIndexer.Query(namespace).ReqTotal()
//	iterator := query.Exec()
//	defer iterator.Close()
//	var items reflect.Value
//	var elementsType reflect.Type
//	for iterator.Next() {
//		o := iterator.Object()
//		if elementsType == nil {
//			elementsType = reflect.TypeOf(o)
//			items = reflect.MakeSlice(reflect.SliceOf(elementsType), 0, iterator.Count())
//		}
//		items = reflect.Append(items, reflect.ValueOf(o))
//	}
//	if elementsType != nil  {
//		cbValue := reflect.ValueOf(cb)
//		in := make([]reflect.Value, 1)
//		in[0] = items
//		cbValue.Call(in)
//	}
//}
