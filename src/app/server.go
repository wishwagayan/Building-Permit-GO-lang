package main

import (
	"go-scraper/src/spider"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Ping..")
	w.Write([]byte("ping sucessfully"))
}

func runServer() {
	addr := ":7181"
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/search", spider.GetData)
	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func main() {
	runServer()
}
