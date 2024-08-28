package main

import (
	"log"
	"os"

	"github.com/aevea/mock-jwt-server/api"
	"github.com/aevea/mock-jwt-server/generator"
)

func main() {
	gen, err := generator.InitGenerator("./tmp")

	if err != nil {
		os.Exit(1)
	}

	certAPI := api.CreateServer(8090, gen.GenerateToken, gen.GenerateJWK)

	err = certAPI.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
