package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	Version := os.Getenv("VERSION")
	NodeName := os.Getenv("NODE_NAME")
	r := mux.NewRouter()
	r.HandleFunc("/v1/version", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, Version+"-"+NodeName)
	}).Methods("GET")

	log.Println("listening at :9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}
