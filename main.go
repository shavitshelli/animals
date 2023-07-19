package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/shavitshelli/puppy"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	r := mux.NewRouter()
	r.HandleFunc("/", servehome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))

}

func servehome(w http.ResponseWriter, r *http.Request) {
	fmt.Println()
	w.Write([]byte("welcome to golang"))

}
