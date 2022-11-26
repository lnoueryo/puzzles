package models

type ActivityContent struct {
	ID      int    `gorm:"AUTO_INCREMENT"json:"id"`
	Content string `json:"content"`
}
