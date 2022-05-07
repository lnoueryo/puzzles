package models

import "time"



type Version struct {
	ID   	  int		`gorm:"AUTO_INCREMENT"json:"id"`
	ProjectID int		`json:"project_id"`
	Name 	  string	`json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime;"json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"json:"-"`
}
