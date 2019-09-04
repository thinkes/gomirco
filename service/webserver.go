package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	r := NewRoute()
	http.Handle("/", r)
	log.Println("Starting Http server at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
