package models

type Choice struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	QuestionID uint   `json:"questionId"`
	Key        string `json:"key"`
	Text       string `json:"text"`
}