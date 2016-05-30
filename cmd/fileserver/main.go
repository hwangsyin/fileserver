package main

import (
	"log"
	"net/http"
	"handler"
	"config"
)

func main() {
	err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}
	
	http.HandleFunc("/files", handler.HandleFile)
	
	err = http.ListenAndServe(config.Configs["port"], nil)
	if err != nil {
		log.Fatal(err)
	}
}