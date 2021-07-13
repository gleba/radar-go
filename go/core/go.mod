module radar.cash/core

go 1.16

require (
	github.com/ClickHouse/clickhouse-go v1.3.14
	github.com/go-pg/pg/v10 v10.10.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/nats-io/nats-server/v2 v2.3.1 // indirect
	github.com/nats-io/nats.go v1.11.1-0.20210623165838-4b75fc59ae30
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/protobuf v1.27.1 // indirect

)

replace radar.cash/core v0.0.0 => ./
