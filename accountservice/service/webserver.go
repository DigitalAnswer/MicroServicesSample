package service

import (
	"log"
	"net/http"
)

// StartWebServer start new webServer
func StartWebServer(port string) {
	log.Println("Starting HTTP service at " + port)

	r := NewRouter()
	http.Handle("/", r)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Panicln("Error: " + err.Error())
	}
}
