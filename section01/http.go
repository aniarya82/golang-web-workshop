package main

import (
	"fmt"
	"log"
	"net/http"
)

func doGet() {
	req, err := http.NewRequest("PUT", "https://http-methods.appspot.com/aniarya82/goisamazinghitentoo", nil)
	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("http request failed: %v", err)
	}
	fmt.Println(res.Status)
}

func main() {
	res, err := http.Get("http:/thisurldoesnotexists")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Status)
	doGet()
}
