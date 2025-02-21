package cmd

import (
	"ecommerce-api/cmd/api"
	"log"
)

func main() {
	server := api.NewAPIServer(":3000", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
