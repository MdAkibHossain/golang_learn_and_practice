package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello!")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":500", r))
}
func greeter() {
	fmt.Println("Greeter")
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Hello there! </h1>"))
}
