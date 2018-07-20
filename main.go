package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//PingHandler - check if server is steel alive
func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "steel alive")
}

//AllProxiesHandler - get all proxies
func AllProxiesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingHandler).Methods("GET")
	r.HandleFunc("/proxy", AllProxiesHandler).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
