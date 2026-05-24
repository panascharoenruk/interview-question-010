package models

type Question struct {
	ID       uint     `json:"id" gorm:"primaryKey"`
	Question string   `json:"question"`
	Answer   string   `json:"-"`
	Choices  []Choice `json:"choices"`
}