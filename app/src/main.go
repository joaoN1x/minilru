package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joaoN1x/minilru/src/debugger"
	rest "github.com/joaoN1x/minilru/src/restful"
)

func main() {

	porta, ok := os.LookupEnv("PORT")
	if !ok {
		porta = "8081"
		debugger.Log("warning", "Can't find incoming port, using default:"+porta, nil)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", rest.HeartBeat).Methods("GET")
	r.HandleFunc("/{short}", rest.GetUrl).Methods("GET")

	log.Println("I'm Listenin on :", porta, "...")

	if err := http.ListenAndServe(string(":"+porta), r); err != nil {
		debugger.Log("error", "Cant Listen on port:"+porta, err)
	}

}
