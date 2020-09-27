package main

import (
    "log"
    "net/http"

	"github.com/brandonforster/coolstuff-go/stringstuff"

    "github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	stringsAPI := r.PathPrefix("/api/strings").Subrouter()
	stringsAPI.HandleFunc("deleter", deleterHandler).Methods(http.MethodPost)


	r.HandleFunc("/", hello)
	
    log.Fatal(http.ListenAndServe(":8080", r))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func deleterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str := vars["str"]
	toRemove := vars["toRemove"]

	result := removeChars(str, toRemove)

	w.Header().Set("Content-Type", "application/json")
	
	w.Write([]byte(result))
}