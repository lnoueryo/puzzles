package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)



type OrganizationAuthority struct {
	ID				int				`gorm:"AUTO_INCREMENT"json:"id"`
	UserID			int				`json:"user_id"`
	OrganizationID	string			`json:"organization_id"`
    AuthorityID		int				`json:"auth_id"`
	Active			bool			`json:"active"`
	Verification	string			`json:"verification"`
	User 			User			`gorm:"foreignkey:UserID;migrate"json:"user"`
	Organization 	Organization	`gorm:"foreignkey:OrganizationID;migrate"json:"organization"`
	Type			Authority		`gorm:"foreignkey:AuthorityID;migrate"json:"type"`
	CreatedAt		time.Time		`gorm:"<-;autoCreateTime;"json:"-"`
	UpdatedAt		time.Time		`gorm:"<-;autoUpdateTime;"json:"-"`
}

func (oa *OrganizationAuthority)Find(verification string) error {
	result := DB.Preload("User").First(&oa, "verification = ? and active = false", verification); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (oa *OrganizationAuthority)Create() error {
	condition := OrganizationAuthority{
		OrganizationID: oa.OrganizationID,
		UserID: oa.UserID,
		Active: oa.Active,
	}
	result := DB.Where(condition).FirstOrCreate(&oa); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (oa *OrganizationAuthority)Update() error {
	condition := OrganizationAuthority{
		OrganizationID: oa.OrganizationID,
		UserID: oa.UserID,
		Active: oa.Active,
	}
	var newOA OrganizationAuthority
	result := DB.Where(condition).First(&newOA); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result = DB.Create(&oa); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		return nil
	}
	newOA.Verification = oa.Verification
	result = DB.Save(&newOA); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (oa *OrganizationAuthority)ChangeActive() error {
	oa.Active = true
	result := DB.Save(&oa); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}