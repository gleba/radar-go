package handlers

import f "github.com/ambelovsky/gosf"

type SioHandler func(client *f.Client, request *f.Request) *f.Message
type handlerMap map[string]SioHandler

func plug(shm handlerMap) {
	for s, handler := range shm {
		f.Listen(s, handler)
	}
}

func bodyData(data interface{}) *f.Message {
	return &f.Message{
		Success: true,
		Body: map[string]interface{}{
			"data": data,
		},
	}
}
