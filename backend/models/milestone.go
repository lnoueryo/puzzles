package models

import "time"



type Milestone struct {
	ID   	  int		`gorm:"AUTO_INCREMENT"json:"id"`
	ProjectID int		`json:"project_id"`
	Name 	  string	`json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime;"json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"json:"-"`
}

func (m *Milestone)DeleteAll(id int) {
	DB.Delete(&m, "project_id = ?", id)
}