package handlers

import (
	f "github.com/ambelovsky/gosf"
)

func init() {
	plug(handlerMap{
		"echo": func(client *f.Client, request *f.Request) *f.Message {
			return f.NewSuccessMessage("pong")
		},
	})
}
