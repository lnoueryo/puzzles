package models

import "time"



type Field struct {
	ID			int			`gorm:"AUTO_INCREMENT"json:"id"`
	ProjectID	int			`json:"project_id"`
	Name		string		`json:"name"`
	CreatedAt	time.Time	`gorm:"autoCreateTime;"json:"created_at"`
	UpdatedAt	time.Time	`gorm:"autoUpdateTime;"json:"updated_at"`
}

func (f *Field)DeleteAll(id int) {
	DB.Delete(&f, "project_id = ?", id)
}