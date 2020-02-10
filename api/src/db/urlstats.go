package db

import (
	"github.com/joaoN1x/minilru/src/debugger"
)

func GetUrlStats24Count(short string) (bool, int64) {

	var (
		returnStatus  bool  = false
		returnMessage int64 = 0
	)

	if dbConnection == nil {
		dbConnection = initPostgresql()
	}

	returnUrl := GetUrl(short)
	if returnUrl.Id > 0 {
		returnStatus = true
	} else {
		debugger.Log("info", "Stats from Inexistant Short URL", nil)
	}

	stmt := `SELECT SUM(count) AS urlstats_total
				FROM urlstats
				WHERE url_short = $1
					AND today >= NOW() - INTERVAL '24 HOURS'`
	rows, err := dbConnection.Query(stmt, short)
	if err != nil {
		debugger.Log("error", "SUM from DB GetUrlStats24Count", err)
	}

	rows.Next()
	rows.Scan(&returnMessage)

	return returnStatus, returnMessage
}

func GetUrlStatsWeekCount(short string) (bool, int64) {

	var (
		returnStatus  bool  = false
		returnMessage int64 = 0
	)

	if dbConnection == nil {
		dbConnection = initPostgresql()
	}

	returnUrl := GetUrl(short)
	if returnUrl.Id > 0 {
		returnStatus = true
	} else {
		debugger.Log("info", "Stats from Inexistant Short URL", nil)
	}

	stmt := `SELECT SUM(count) AS urlstats_total
				FROM urlstats
				WHERE url_short = $1
					AND today >= NOW() - INTERVAL '1 WEEK'`
	rows, err := dbConnection.Query(stmt, short)
	if err != nil {
		debugger.Log("error", "SUM from DB GetUrlStatsWeekCount", err)
	}

	rows.Next()
	rows.Scan(&returnMessage)

	return returnStatus, returnMessage
}

func GetUrlStatsAllCount(short string) (bool, int64) {

	var (
		returnStatus  bool  = false
		returnMessage int64 = 0
	)

	if dbConnection == nil {
		dbConnection = initPostgresql()
	}

	returnUrl := GetUrl(short)
	if returnUrl.Id > 0 {
		returnStatus = true
	} else {
		debugger.Log("info", "Stats from Inexistant Short URL", nil)
	}

	stmt := `SELECT SUM(count) AS urlstats_total
				FROM urlstats
				WHERE url_short = $1`
	rows, err := dbConnection.Query(stmt, short)
	if err != nil {
		debugger.Log("error", "SUM from DB GetUrlStatsAllCount", err)
	}

	rows.Next()
	rows.Scan(&returnMessage)

	return returnStatus, returnMessage
}
