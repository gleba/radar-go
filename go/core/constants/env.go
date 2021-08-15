package constants

var PostgresAddr = ReadEnv("POSTGRES_ADDR", "localhost:5432")
var PostgresUser = ReadEnv("POSTGRES_USER", "pegase")
var PostgresPassword = ReadEnv("POSTGRES_PASSWORD", "pegase")
var PostgresDbName = ReadEnv("POSTGRES_DB_NAME", "pegase")

var PostgresTempTables = ReadEnv("POSTGRES_TEMPORARY_TABLES", "false")

var PORT = ReadEnv("LLS_PORT", "2080")

var Version string
