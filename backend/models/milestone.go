package models

import (
	"time"

	"gorm.io/gorm"
)



type Milestone struct {
	ID   	  int		`gorm:"AUTO_INCREMENT"json:"id"`
	ProjectID int		`json:"project_id"`
	Name 	  string	`json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime;"json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"json:"-"`
}

func (m *Milestone)DeleteAll(DB *gorm.DB, id int) error {
	if err := DB.Delete(&m, "project_id = ?", id).Error;err != nil {
		return err
	}
	return nil
}

func (m *Milestone)DeleteByProjectID(DB *gorm.DB, projectID int) error {
	if err := DB.Delete(&m, "project_id = ?", projectID).Error;err != nil {
		return err
	}
	return nil
}