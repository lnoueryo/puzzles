package models



type Status struct {
	ID   int    `gorm:"AUTO_INCREMENT"json:"id"`
	Name string `json:"name"`
}