package seeders

import (
	"example.com/exam-api/configs"
	"example.com/exam-api/models"
)

func SeedQuestions() {
	var count int64

	configs.DB.Model(&models.Question{}).Count(&count)

	if count > 0 {
		return
	}

	questions := []models.Question{
		{
			Question: "2 + 2 = ?",
			Answer:   "B",
			Choices: []models.Choice{
				{Key: "A", Text: "3"},
				{Key: "B", Text: "4"},
				{Key: "C", Text: "5"},
				{Key: "D", Text: "6"},
			},
		},
		{
			Question: "Capital of Thailand?",
			Answer:   "C",
			Choices: []models.Choice{
				{Key: "A", Text: "Tokyo"},
				{Key: "B", Text: "Seoul"},
				{Key: "C", Text: "Bangkok"},
				{Key: "D", Text: "Beijing"},
			},
		},
	}

	for _, question := range questions {
		configs.DB.Create(&question)
	}
}