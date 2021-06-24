module radar.cash/pharos

require (
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/nats-io/nats-server/v2 v2.1.6 // indirect
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
	radar.cash/core v0.0.0
)

replace radar.cash/core v0.0.0 => ../core

go 1.14
