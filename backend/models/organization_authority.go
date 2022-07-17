package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
)



type OrganizationAuthority struct {
	ID				int				`gorm:"AUTO_INCREMENT"json:"id"`
	Active			bool			`json:"active"`
	Verification	string			`json:"verification"`
    AuthorityID		int				`json:"auth_id"`
	OrganizationID	string			`json:"organization_id"`
	UserID			int				`json:"user_id"`
	User 			User			`gorm:"foreignkey:UserID;"json:"user"`
	Organization 	Organization	`gorm:"foreignkey:OrganizationID;"json:"organization"`
	Type			Authority		`gorm:"foreignkey:AuthorityID;"json:"type"`
	CreatedAt		time.Time		`gorm:"<-;autoCreateTime;"json:"created_at"`
	UpdatedAt		time.Time		`gorm:"<-;autoUpdateTime;"json:"-"`
}


func NewOrganizationAuthority(r *http.Request) (OrganizationAuthority, error) {
	organizationAuthority, _ := GetOrganizationAuthorityJson(r)
	return organizationAuthority, nil
}

func (oa *OrganizationAuthority)Find(DB *gorm.DB, verification string) error {
	result := DB.Preload("User").First(&oa, "verification = ? and active = false", verification); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (oa *OrganizationAuthority)Create(DB *gorm.DB) error {
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

func (oa *OrganizationAuthority)DeleteByUserIDs(DB *gorm.DB, ids []int) error {
	result := DB.Delete(&oa, "user_id = ?", ids); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func GetOrganizationAuthorityJson(r *http.Request) (OrganizationAuthority, error) {
	var organizationAuthority OrganizationAuthority
	err := json.NewDecoder(r.Body).Decode(&organizationAuthority)
	if err != nil {
		fmt.Println(err)
	}
	return organizationAuthority, nil
}


func (oa *OrganizationAuthority)ChangeActive(DB *gorm.DB) error {
	oa.Active = true
	result := DB.Save(&oa); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (oa *OrganizationAuthority)Update(DB *gorm.DB) error {
	result := DB.Omit("User", "Organization", "Type").Save(&oa); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}