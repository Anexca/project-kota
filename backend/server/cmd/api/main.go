package main

import (
	"log"
	"server/internal/server"
)

func main() {

	server := server.InitServer()

	log.Println("Starting server on address", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("cannot start server: %s", err)
	}
}
