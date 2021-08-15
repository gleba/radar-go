package handlers

import (
	f "github.com/ambelovsky/gosf"
	"radar.cash/pharos/pegas/magic"
	"radar.cash/pharos/vars"
)

func init() {
	plug(handlerMap{
		"dict": func(client *f.Client, request *f.Request) *f.Message {
			return bodyData(vars.FrontDict)
		},
		"activeAlerts": func(client *f.Client, request *f.Request) *f.Message {
			return bodyData(magic.MakeFrontImpulses())
		},
	})
}
