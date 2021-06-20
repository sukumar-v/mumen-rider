package main

import (
	"log"
	"net/http"

	"github.com/sukumar-varma/mumen-rider/crawler"
)

func main() {
	addr := ":9001"

	http.HandleFunc("/ping", crawler.Ping)
	http.HandleFunc("/get-url", crawler.GetUrls)

	log.Printf("Listening on %q", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
