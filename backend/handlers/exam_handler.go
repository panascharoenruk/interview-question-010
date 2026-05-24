package handlers

import (
	"net/http"
	"strconv"

	"example.com/exam-api/configs"
	"example.com/exam-api/models"
	"github.com/gin-gonic/gin"
)

type SubmitRequest struct {
	Name    string            `json:"name"`
	Answers map[string]string `json:"answers"`
}

func GetQuestions(c *gin.Context) {
	var questions []models.Question

	configs.DB.Preload("Choices").Find(&questions)

	c.JSON(http.StatusOK, questions)
}

func SubmitExam(c *gin.Context) {
	var request SubmitRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	var questions []models.Question

	configs.DB.Find(&questions)

	score := 0

	for _, question := range questions {
		questionID := strconv.Itoa(int(question.ID))

		selected := request.Answers[questionID]

		if selected == question.Answer {
			score++
		}
	}

	result := models.ExamResult{
		Name:  request.Name,
		Score: score,
		Total: len(questions),
	}

	configs.DB.Create(&result)

	c.JSON(http.StatusOK, gin.H{
		"score": score,
		"total": len(questions),
	})
}