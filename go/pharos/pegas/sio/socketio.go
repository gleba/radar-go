package sio

import (
	f "github.com/ambelovsky/gosf"
	"radar.cash/pharos/vars"
)

func BodyData(data interface{}) *f.Message {
	return &f.Message{
		Success: true,
		Body: map[string]interface{}{
			"data": data,
		},
	}
}

//func leaveAll(client *f.Client)  {
//for _, r := range client.Rooms {
//
//	client.LeaveAll()
//}
//}
func UpdateDict() {
	f.Broadcast("main", "dict", BodyData(vars.FrontDict))
}

//
//var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
//
//func RandStringRunes(n int) string {
//	b := make([]rune, n)
//	for i := range b {
//		b[i] = letterRunes[rand.Intn(len(letterRunes))]
//	}
//	return string(b)
//}
func Start() {
	//rand.Seed(time.Now().UnixNano())
	f.OnConnect(func(client *f.Client, request *f.Request) {
		client.Join("main")
	})
	f.OnDisconnect(func(client *f.Client, request *f.Request) {
		client.LeaveAll()
	})
	go f.Startup(map[string]interface{}{
		"enableCORS": "*",
		"port":       8180})
}
