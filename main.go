package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = 8080

type application struct {
	Domain string
}

func main() {
	//set appication config
	var app application
	//read from command line
	// connect to DB
	app.Domain = "example.com"
	//start a web server
	log.Println("Starting Application on port:", PORT)

	//http.HandleFunc("/", Hello)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), app.routes())

	if err != nil {
		log.Fatal(err)
	}

}
