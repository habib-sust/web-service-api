package main

import (
	"fmt"
	"habib/web-service-api/app/router"
	"habib/web-service-api/config"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error in loading .env file, Error:", err)
	}
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")
	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
