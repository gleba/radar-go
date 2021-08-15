module radar.cash/pharos

require (
	github.com/ambelovsky/go-structs v1.1.0 // indirect
	github.com/ambelovsky/gosf v0.0.0-20201109201340-237aea4d6109
	github.com/ambelovsky/gosf-socketio v0.0.0-20201109193639-add9d32f8b19 // indirect
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/googollee/go-socket.io v1.6.0
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	radar.cash/core v0.0.0
)

replace radar.cash/core v0.0.0 => ../core

go 1.16
