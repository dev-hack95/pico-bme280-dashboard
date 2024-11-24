package routes

import (
	"github.com/dev-hack95/pico-bme280-dashboard/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/pico/dashboard/login", handlers.UserLogin)
	r.POST("/pico/dashboard/user/create", handlers.CreateUserAccount)
	r.GET("/pico/dashboard/logout", handlers.UserLogout)
	r.GET("/pico/dashboard/chart/humidity", handlers.GetHumidityChartDetails)
	r.GET("/pico/dashboard/chart/temperature", handlers.GetTempreatureChartDetails)
	r.GET("/pico/dashboard/chart/pressure", handlers.GetPressureChartDetails)
}
