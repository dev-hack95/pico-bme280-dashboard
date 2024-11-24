package main

import (
	"log"
	"os"

	"github.com/dev-hack95/pico-bme280-dashboard/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())

	routes.RegisterRenderRoutes(r)
	routes.RegisterRoutes(r)

	r.Run(os.Getenv("HOST"))
}
