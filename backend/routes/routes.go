package routes

import (
	"example.com/exam-api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	{
		api.GET("/questions", handlers.GetQuestions)
		api.POST("/submit", handlers.SubmitExam)
	}
}