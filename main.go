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

	r.HandleFunc("/", helloHandler)

	log.Println("Starting server on port 5000")

	log.Fatal(http.ListenAndServe(":5000", r))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!\n"))
}

func deleterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str := vars["str"]
	toRemove := vars["toRemove"]

	result := stringstuff.RemoveChars(str, toRemove)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(result))
}
