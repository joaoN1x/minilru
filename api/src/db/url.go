package db

import (
	"fmt"

	"github.com/joaoN1x/minilru/src/debugger"
	"github.com/joaoN1x/minilru/src/process"
)

func AddUrl(url Url) (bool, string) {

	var (
		returnStatus  bool   = true
		returnMessage string = ""
	)

	if dbConnection == nil {
		dbConnection = initPostgresql()
	}

	//check url func returns proper format
	if url.Long == "" {
		returnMessage = returnMessage + "Need an Url to process"
		returnStatus = false
	} else {
		query := `
				INSERT INTO url (long, short) VALUES($1, '')
				RETURNING id`
		statement, err := dbConnection.Prepare(query)
		if err != nil {
			returnMessage = returnMessage + string(err.Error())
			returnStatus = false
			debugger.Log("error", "Query Insert string didn't work.", err)
		}

		lid := int64(0)
		err = statement.QueryRow(url.Long).Scan(&lid)
		if err != nil {
			returnMessage = returnMessage + string(err.Error())
			returnStatus = false
			debugger.Log("error", "Couldn't write to database.", err)
		}

		if lid > 0 {
			url.Id = lid
			//process urlshortener
			urlShort, ok := process.GenerateShortUrl(lid)
			if !ok {
				returnMessage = returnMessage + " Couldn't generate short."
				returnStatus = false
			} else {
				url.Short = urlShort
				updateUrlShort(url)
				returnMessage = url.Short
			}
		} else {
			returnMessage = returnMessage + " Couldn't add url to database."
			returnStatus = false
			debugger.Log("error", "Couldn't add url to database.", nil)
		}
	}
	return returnStatus, returnMessage
}

// updateUrl updates an existant Device record at the database
func updateUrlShort(url Url) (bool, string) {
	var (
		returnStatus  bool   = true
		returnMessage string = ""
	)

	if dbConnection == nil {
		dbConnection = initPostgresql()
	}

	if url.Id == 0 {
		returnMessage = returnMessage + " No Url ID found for update"
		returnStatus = false
		debugger.Log("warning", "No Url ID found for update.", nil)
		return returnStatus, returnMessage
	}

	stmt := fmt.Sprintf(`UPDATE url
		SET short = $1
		WHERE id = $2`)

	statement, err := dbConnection.Prepare(stmt)
	if err != nil {
		returnMessage = returnMessage + string(err.Error())
		returnStatus = false
		debugger.Log("error", "Query Update string didn't work.", err)
	}

	_, err = statement.Exec(url.Short, url.Id)
	if err != nil {
		returnMessage = returnMessage + string(err.Error())
		returnStatus = false
		debugger.Log("error", "Couldn't write to database.", err)
	}

	return returnStatus, returnMessage
}
