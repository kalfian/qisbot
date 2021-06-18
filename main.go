package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kalfian/qisbot/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := ":" + os.Getenv("APP_PORT")

	r := gin.Default()

	routes.InitRoute(r)

	if err := r.Run(port); err != nil {
		log.Fatalf("Error running server at %v: %+v", port, err)
	}
}
