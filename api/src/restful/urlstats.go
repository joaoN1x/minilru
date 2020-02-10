package restful

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/joaoN1x/minilru/src/db"
	"github.com/joaoN1x/minilru/src/debugger"
	t "github.com/joaoN1x/minilru/src/types"
)

// GetUrlStats24Count support interface to create a new Url record
func GetStats(w http.ResponseWriter, r *http.Request) {

	var (
		resultDb      bool  = false
		resultMessage int64 = 0
	)

	params := mux.Vars(r)

	doWhen := params["when"]
	urlShort := params["short"]

	switch doWhen {
	case "day":
		resultDb, resultMessage = db.GetUrlStats24Count(urlShort)
	case "week":
		resultDb, resultMessage = db.GetUrlStatsWeekCount(urlShort)
	case "all":
		resultDb, resultMessage = db.GetUrlStatsAllCount(urlShort)
	default:
		debugger.Log("info", "No valid stats range chosen", nil)
	}

	var data t.MessageOutData
	data.Detail = strconv.FormatInt(resultMessage, 10)

	if resultDb {
		writeOut(w, r, 200, t.MessageOut{Status: "OK", Code: 200, Data: data})
	} else {
		debugger.Log("warning", "204 No Content", nil)
		writeOut(w, r, 204, t.MessageOut{Status: "No Content", Code: 204, Data: data})
	}

}
