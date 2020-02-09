package restful

import (
	//"encoding/json"
	"encoding/json"
	"io/ioutil"

	//"net/http"
	//"strings"

	"net/http"

	"github.com/joaoN1x/minilru/src/db"
	"github.com/joaoN1x/minilru/src/debugger"

	t "github.com/joaoN1x/minilru/src/types"
)

// AddUrl support interface to create a new Url record
func AddUrl(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		debugger.Log("error", "Error", err)
	}
	debugger.PrettyPrint(string(reqBody))
	bytes := []byte(string(reqBody))
	var url db.Url
	json.Unmarshal(bytes, &url)

	resultDb, resultMessage := db.AddUrl(url)

	debugger.PrettyPrint(resultDb)
	debugger.PrettyPrint(resultMessage)

	var data t.MessageOutData
	data.Detail = resultMessage
	if resultDb {
		data.Affected = 1
		writeOut(w, r, 200, t.MessageOut{Status: "Added", Code: 200, Data: data})
	} else {
		writeOut(w, r, 409, t.MessageOut{Status: "Conflict", Code: 409, Data: data})
	}

}
