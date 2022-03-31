package models

import (
	"errors"
	"time"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)



type Organization struct {
	ID					string		`gorm:"primaryKey;type:varchar(100);unique_index"json:"id"`
	Name				string		`json:"name"`
	Address				string		`json:"address"`
	Number				string		`json:"number"`
	Founded				string		`json:"founded"`
	Image				string		`json:"image"`
	Description			string		`json:"description"`
	Plan				string		`json:"plan"`
	CreditCard			string		`json:"creditcard"`
	Expiry				string		`json:"expiry"`
	Projects			[]Project	`json:"projects"`
	Users				[]User		`gorm:"many2many:organization_authorities;"json:"users"`
	CreatedAt			time.Time	`gorm:"autoCreateTime;"json:"created_at"`
	UpdatedAt			time.Time	`gorm:"autoUpdateTime;"json:"updated_at"`
}

func (o *Organization)GetOrganization(id string) error {
	result := DB.Preload("Projects", func(db *gorm.DB) *gorm.DB {
		return DB.Preload(clause.Associations)
	  }).Preload(clause.Associations).First(&o, "id = ?", id); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}
