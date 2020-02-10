package db

import (
	"database/sql"
	"fmt"

	"github.com/joaoN1x/minilru/src/debugger"
	_ "github.com/lib/pq"
)

var (
	dbConnection *sql.DB
	config       = GetConfig()
)

func initPostgresql() *sql.DB {

	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?application_name=postgres&sslmode=disable", config[0].PostgresqlList[0].Databases[0].Username, config[0].PostgresqlList[0].Databases[0].Password, config[0].PostgresqlList[0].Server, config[0].PostgresqlList[0].Port, config[0].PostgresqlList[0].Databases[0].Name)
	var postgresqlConnection, postgresqlConnectionErr = sql.Open("postgres", connectionString)

	if postgresqlConnectionErr != nil {
		debugger.Log("error", "error connecting to the database", postgresqlConnectionErr)
	}

	return postgresqlConnection

}
