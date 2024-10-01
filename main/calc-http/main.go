package main

import (
	"log"
	"net/http"

	"github.com/mdwhatcott/calc-apps/handlers"
)

func main() {
	address := ":8080"
	log.Println("Serving requests at:", address)
	err := http.ListenAndServe(address, handlers.NewHTTPRouter())
	if err != nil {
		log.Fatalln(err)
	}
}
