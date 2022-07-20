package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID			int				`gorm:"AUTO_INCREMENT"json:"id"`
	UserID		int				`json:"user_id"`
	ProjectID	int				`json:"project_id"`
	ContentID	int				`json:"content_id"`
	User		User			`gorm:"references:ID;"json:"user"`
	Content		ActivityContent	`gorm:"references:ID;"json:"content"`
	CreatedAt	time.Time 		`gorm:"autoCreateTime;"json:"created_at"`
	UpdatedAt	time.Time 		`gorm:"autoUpdateTime;"json:"updated_at"`
}


func (a *Activity)Create(DB *gorm.DB) error {
	result := DB.Create(&a); if result.Error != nil {
		return result.Error
	}
	return nil
}