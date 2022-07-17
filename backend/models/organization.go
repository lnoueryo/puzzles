package models

import (
	"backend/modules/crypto"
	"encoding/json"
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
	ImageData			string					`gorm:"-:all;"json:"image_data"`
	Projects			[]Project				`json:"projects"`
	Users				[]OrganizationAuthority	`gorm:"foreignkey:OrganizationID;"json:"users"`
	CreatedAt			time.Time				`gorm:"autoCreateTime;"json:"created_at"`
	UpdatedAt			time.Time				`gorm:"autoUpdateTime;"json:"updated_at"`
}

func (o *Organization)GetOrganization(DB *gorm.DB, id string) error {
	result := DB.Preload("Projects", func(db *gorm.DB) *gorm.DB {
		return DB.Preload(clause.Associations)
	  }).Preload(clause.Associations).First(&o, "id = ?", id); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func(o *Organization) Update(DB *gorm.DB) error {
	if o.ImageData != "" {
		deleteImageName := o.Image
		fileName, err := UploadToGCS("organizations", o.ImageData); if err != nil {
			return errors.New("couldn't save the image")
		}
		o.Image = fileName
		result := DB.Debug().Omit("Organization", "Projects", "Users").Session(&gorm.Session{FullSaveAssociations: true}).Save(&o); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		DeleteImage(deleteImageName, "organizations")
		return nil
	}
	result := DB.Debug().Omit("Organization", "Projects", "Users").Session(&gorm.Session{FullSaveAssociations: true}).Save(&o); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
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


func GetOrganizationJson(r *http.Request) (Organization, error) {
	var organization Organization
	err := json.NewDecoder(r.Body).Decode(&organization)
	if err != nil {
		message := "couldn't decode json"
		err := errors.New(message)
		return organization, err
	}
	return organization, nil
}