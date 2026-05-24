package main

import (
	"example.com/exam-api/configs"
	"example.com/exam-api/database"
	"example.com/exam-api/routes"
	"example.com/exam-api/seeders"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDatabase()

	database.MigrateDatabase()

	seeders.SeedQuestions()

	router := gin.Default()

	router.Use(cors.Default())

	routes.SetupRoutes(router)

	router.Run(":8080")
}