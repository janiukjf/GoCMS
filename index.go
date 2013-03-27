package main

import (
	"./gophers/controllers"
	"./gophers/plate"
	"net/http"
	"flag"
	"log"
)

var (
	listenAddr = flag.String("http", ":8080", "http listen address")

	CorsHandler = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		return
	}
)

const (
	port = "80"
)

func main() {
	flag.Parse()
	server := plate.NewServer("doughboy")

	server.AddFilter(CorsHandler)

	server.Get("/", controllers.Index)

	http.Handle("/", server)
	http.ListenAndServe(*listenAddr, nil)

	log.Println("Server running on port " + *listenAddr)
}
