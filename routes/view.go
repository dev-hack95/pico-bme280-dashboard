package routes

import (
	"github.com/dev-hack95/pico-bme280-dashboard/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRenderRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/", handlers.HomePage)
	r.GET("/signup-form", handlers.SignupForm)
	r.GET("/dashboard", handlers.Dashboard)
}
