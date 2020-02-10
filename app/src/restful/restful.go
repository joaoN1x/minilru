package restful

import (
	"net/http"

	t "github.com/joaoN1x/minilru/src/types"
)

// HeartBeat, the rest answer for health check calls
func HeartBeat(w http.ResponseWriter, r *http.Request) {
	writeOut(w, r, 200, t.MessageOut{Status: "OK", Code: 200})
}
