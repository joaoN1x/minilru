package db

import (
	"fmt"

	"github.com/joaoN1x/minilru/src/debugger"
)

func SumUrlStats(short string) (bool, string) {

	var (
		returnStatus  bool   = true
		returnMessage string = ""
	)

	if dbConnection == nil {
		dbConnection = initPostgresql()
	}

	stmt := fmt.Sprintf(`WITH upsert AS (
		UPDATE urlstats SET count=count+1 
		WHERE url_short=$1 
			AND today = 'today' RETURNING *
		)
		INSERT INTO urlstats 
		(url_short, today, count) 
		SELECT $1, 'today', 1
			WHERE NOT EXISTS (SELECT * FROM upsert);`)

	statement, err := dbConnection.Prepare(stmt)
	if err != nil {
		returnMessage = returnMessage + string(err.Error())
		returnStatus = false
		debugger.Log("error", "Query Upsert string didn't work.", err)
	}

	_, err = statement.Exec(short)
	if err != nil {
		returnMessage = returnMessage + string(err.Error())
		returnStatus = false
		debugger.Log("error", "Couldn't write to database.", err)
	}

	return returnStatus, returnMessage
}
