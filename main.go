package main

import (
	"log"
	"post-api/config"
	"post-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.InitDB()

	server := gin.New()

	routes.SetupRoute(server)

	log.Fatal(server.Run(":9090"))
}
