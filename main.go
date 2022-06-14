package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	now := time.Now().UTC().Format(http.TimeFormat)
	w.Header().Set("Date", now)
	w.Header().Set("Last-Modified", now)
	w.Header().Set("Vary", "Accept")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Happy learning, keep learning"))
}
func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", home)
	log.Println("Server starting on port 3214")
	log.Fatal(http.ListenAndServe(":3214", myRouter))
}
