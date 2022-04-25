package models

import (
	"backend/modules/crypto"
	"errors"
	"io"
	"net/http"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)



type Organization struct {
	ID					string					`gorm:"primaryKey;type:varchar(100);unique_index"json:"id"`
	Name				string					`json:"name"`
	Address				string					`json:"address"`
	Number				string					`json:"number"`
	Founded				string					`json:"founded"`
	Image				string					`json:"image"`
	Description			string					`json:"description"`
	Plan				string					`json:"plan"`
	CreditCard			string					`json:"creditcard"`
	Expiry				string					`json:"expiry"`
	Projects			[]Project				`json:"projects"`
	Users				[]OrganizationAuthority	`gorm:"foreignkey:OrganizationID;"json:"users"`
	CreatedAt			time.Time				`gorm:"autoCreateTime;"json:"created_at"`
	UpdatedAt			time.Time				`gorm:"autoUpdateTime;"json:"updated_at"`
}

func (o *Organization)GetOrganization(id string) error {
	result := DB.Preload("Projects", func(db *gorm.DB) *gorm.DB {
		return DB.Preload(clause.Associations)
	  }).Preload(clause.Associations).First(&o, "id = ?", id); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (o *Organization) GetImage() {
	url := "https://loremflickr.com/320/240?random=1"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	randStr, _ := crypto.MakeRandomStr(15)
	extension := ".png"
	filename := randStr + extension
	path := "upload/organizations/"

	file, err := os.Create(path + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, response.Body)
	o.Image = filename
}