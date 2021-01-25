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
	addr := ":7171"
	http.HandleFunc("/ping", ping)
	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func main() {
	spider.Phase()
}
