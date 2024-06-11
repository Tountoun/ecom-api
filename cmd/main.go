package main

import (
	"log"

	"github.com/Tountoun/ecom-api/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}