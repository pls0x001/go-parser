package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	. "planScrapper/controllers"
	"time"
)

func main() {
	router := gin.New()
	router.LoadHTMLGlob("image.html")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/", GetImage)
	router.GET("/mfz", GetMFZ12All)
	router.GET("/mfz-week", GetMFZ12PerWeek)
	router.GET("/nur", GetS2MPAll)
	router.GET("/nur-week", GetS2MPPerWeek)
	err := router.Run()

	if err != nil {
		return
	}
}
