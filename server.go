package main

import (
	"fmt"
	"log"
	"os"

	"github.com/4geeks/pex-awardwallet/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := routes.SetupRouter()

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
