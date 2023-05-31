package models

import (
	"time"

	"gorm.io/gorm"
)



type Field struct {
	ID			int			`gorm:"AUTO_INCREMENT"json:"id"`
	ProjectID	int			`json:"project_id"`
	Name		string		`json:"name"`
	CreatedAt	time.Time	`gorm:"autoCreateTime;"json:"created_at"`
	UpdatedAt	time.Time	`gorm:"autoUpdateTime;"json:"updated_at"`
}

func (f *Field)DeleteAll(DB *gorm.DB, id int) error {
	if err := DB.Delete(&f, "project_id = ?", id).Error;err != nil {
		return err
	}
	return nil
}

func (f *Field)DeleteByProjectID(DB *gorm.DB, projectID int) error {
	if err := DB.Delete(&f, "project_id = ?", projectID).Error;err != nil {
		return err
	}
	return nil
}