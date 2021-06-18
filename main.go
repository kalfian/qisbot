package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	sdk "github.com/kalfian/qisbot/qiscus_sdk"
	"github.com/kalfian/qisbot/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := ":" + os.Getenv("APP_PORT")

	r := gin.Default()
	sdk := sdk.NewQiscusSDK(sdk.ParamQiscusSDK{
		AppID:    os.Getenv("QISCUS_APP_ID"),
		SecretID: os.Getenv("QISCUS_SECRET"),
		BaseUrl:  os.Getenv("QISCUS_BASE_URL"),
	})

	routes.InitRoute(r, sdk)

	if err := r.Run(port); err != nil {
		log.Fatalf("Error running server at %v: %+v", port, err)
	}
}
