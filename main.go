package main

import (
	"log"
	"os"

	"github.com/GitSplit-org/backend/api"
	"github.com/GitSplit-org/backend/config/dbconfig"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dbconfig.Init()
	ginApp := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	ginApp.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": 404, "message": "Invalid Endpoint Request"})
	})
	ginApp.Use(cors.New(config))
	api.ApplyRoutes(ginApp)
	err := ginApp.Run(":" + os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
