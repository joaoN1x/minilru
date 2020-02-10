package db

import (
	"database/sql"
	"fmt"

	"github.com/joaoN1x/minilru/src/debugger"
)

func GetUrl(short string) Url {
	var (
		urlRecord = Url{}
	)

	if dbConnection == nil {
		dbConnection = initPostgresql()
	}

	stmt := `SELECT id AS url_id,
		long AS url_long,
		short AS url_short
	FROM url
	WHERE short = $1
	`
	row := dbConnection.QueryRow(stmt, short)

	switch err := row.Scan(&urlRecord.Id, &urlRecord.Long, &urlRecord.Short); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(urlRecord.Id, urlRecord.Long, urlRecord.Short)
	default:
		debugger.Log("error", "Select from DB url", err)
	}

	return urlRecord
}
