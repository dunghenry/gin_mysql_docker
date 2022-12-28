package main

import (
	"gin/mysql/routes"

	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.SetTrustedProxies(nil)
	routes.Run(r)
}
