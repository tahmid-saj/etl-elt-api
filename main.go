package main

import (
	"os"
	"tahmid-saj/etl-elt-api/db/mongodb"
	"tahmid-saj/etl-elt-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db.InitMongoDB()

	server := gin.Default()

	routes.RegisterRoutes(server)
	
	server.Run(os.Getenv("PORT"))

	db.DisconnectMongoDB()
}