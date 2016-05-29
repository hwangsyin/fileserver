package main

import (
	"log"
	"net/http"
	"handler"
)

func main() {
	http.HandleFunc("/files", handler.HandleFile)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}