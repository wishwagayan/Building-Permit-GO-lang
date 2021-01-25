package main

import (
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Ping..")
	w.Write([]byte("ping"))
}

func runServer() {
	addr := ":7171"
	http.HandleFunc("/ping", ping)
}
