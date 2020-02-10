package restful

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/joaoN1x/minilru/src/cache"
	"github.com/joaoN1x/minilru/src/db"
	"github.com/joaoN1x/minilru/src/debugger"
	t "github.com/joaoN1x/minilru/src/types"
)

// GetUrl support interface to get an Url to redirect
func GetUrl(w http.ResponseWriter, r *http.Request) {

	//start := time.Now()

	params := mux.Vars(r)
	short := params["short"]
	urlLong := cache.GetUrl(short)

	if urlLong == "" {
		url := db.GetUrl(short)
		if url.Long != "" {
			urlLong = url.Long
			cache.SetUrl(url.Short, url.Long)
		} else {
			var data t.MessageOutData
			data.Detail = "Inexistant"
			debugger.Log("warning", "404 Inexistant", nil)

			writeOut(w, r, 404, t.MessageOut{Status: "Resource not found", Code: 404, Data: data})
		}
	}

	defer duration(short, time.Now())

	if urlLong != "" {
		http.Redirect(w, r, urlLong, http.StatusTemporaryRedirect)
		go func() {
			resultDb, resultMessage := db.SumUrlStats(short)
			fmt.Println(resultDb, resultMessage)
		}()

	}

}

func duration(short string, start time.Time) {
	duration := time.Since(start)
	debugger.Log("info", "ShortURL "+short+" done in "+duration.String(), nil)
}
