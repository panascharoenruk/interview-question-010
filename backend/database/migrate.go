package database

import (
	"example.com/exam-api/configs"
	"example.com/exam-api/models"
)

func MigrateDatabase() {
	configs.DB.AutoMigrate(
		&models.Question{},
		&models.Choice{},
		&models.ExamResult{},
	)
}