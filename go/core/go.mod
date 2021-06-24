module radar.cash/core

go 1.16

require (
	github.com/ClickHouse/clickhouse-go v1.3.14
	github.com/a8m/syncmap v0.0.0-20210309180435-03e1ab403ea9 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/nats-io/nats.go v1.11.0
	github.com/vmihailenco/msgpack/v4 v4.3.10 // indirect
	golang.org/x/sys v0.0.0-20210514084401-e8d321eab015 // indirect
	golang.org/x/tools v0.1.1 // indirect

)

replace radar.cash/core v0.0.0 => ./
