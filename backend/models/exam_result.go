package models

type ExamResult struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Score int    `json:"score"`
	Total int    `json:"total"`
}