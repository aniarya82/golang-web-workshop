package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Heelow web")
}
func byeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bye, we")
}
func listProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "list all products")
}
func addProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "add a product")
}
func getProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["productID"]
	log.Printf("fetching product with ID %q", id)
	// get a specific product
	fmt.Fprintln(w, "get a specific product")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/product/", listProducts).Methods("GET")
	r.HandleFunc("/product/", addProduct).Methods("POST")
	r.HandleFunc("/product/{productID}", getProduct)
	http.Handle("/", r)
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatal(err)
	}
}
