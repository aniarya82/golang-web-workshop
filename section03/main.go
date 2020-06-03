package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func paramHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		name = "John Doe"
	}
	fmt.Fprintf(w, "Hello my freind %s !!", name)
}
func bodyHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "could not read body : %v", err)
		return
	}
	name := string(b)
	if name == "" {
		name = "friend"
	}
	fmt.Fprintf(w, "Hello,, %s!", name)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", bodyHandler)
	http.Handle("/", r)
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal(err)
	}
}
