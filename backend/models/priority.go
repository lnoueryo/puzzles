package models



type Priority struct {
	ID   int    `gorm:"AUTO_INCREMENT"json:"id"`
	Name string `json:"name"`
}