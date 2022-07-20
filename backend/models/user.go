package models

import (
	"backend/modules/crypto"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID						int						`gorm:"AUTO_INCREMENT"json:"id"`
	Name					string					`json:"name"`
	Age						int						`json:"age"`
	Sex						string					`json:"sex"`
	Email					string					`json:"email"`
	Address					string					`json:"address"`
	Password				string					`json:"-"`
	Image					string					`json:"image"`
	ImageData				string					`gorm:"<-:false;-:migration;"json:"image_data"`
	Description				string					`json:"description"`
	Organization			string					`gorm:"-:all;"json:"organization"`
	ChangePassword			string					`gorm:"-:all;"json:"password"`
	Projects				[]ProjectAuthority		`json:"projects"`
	Organizations			[]OrganizationAuthority	`gorm:"foreignkey:UserID;"json:"organizations"`
	Tasks					[]Task					`gorm:"foreignKey:AssigneeID;references:ID"json:"tasks"`
	CreatedAt				time.Time				`gorm:"->:false;<-:create;autoCreateTime;"json:"-"`
	UpdatedAt				time.Time				`gorm:"autoUpdateTime;"json:"updated_at"`
}


func NewUser(r *http.Request) (User, error) {
	user, _ := GetUserJson(r)
	randStr, _ := crypto.MakeRandomStr(20)
	filename := randStr + ".png"
	password := crypto.Encrypt(user.Password)
	user = User{Name: user.Name, Email: user.Email, Image: filename, Password: password}
	return user, nil
}

func (u *User) Create(DB *gorm.DB) error {
	result := DB.Create(u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (u *User) Update(DB *gorm.DB) error {
	if u.ChangePassword != "" {
		u.Password = crypto.Encrypt(u.ChangePassword)
		result := DB.Omit("Organizations").Save(u)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		return nil
	}
	result := DB.Omit("Organizations", "Password").Save(u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (u *User)GetMainUser(DB *gorm.DB, userID int, orgID string) error {
	result := DB.Preload("Organizations.Organization.Projects.AuthorityUsers." + clause.Associations).
	Preload("Organizations.Organization.Projects.AuthorityUsers", "active = ?", true).
	Preload("Organizations.Organization.Projects.Fields").
	Preload("Organizations.Organization.Projects.Milestones").
	Preload("Organizations.Organization.Projects.Versions").
	Preload("Organizations.Organization.Projects").
	Preload("Organizations.Organization.Users.User").
	Preload("Organizations.Organization.Users.Type").
	Preload("Organizations.Organization.Users").
	Preload("Organizations." + clause.Associations).
	Preload("Organizations", "organization_id = ?", orgID).
	Preload("Organizations").
	First(&u, "id = ?", userID); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	u.Password = ""
	return nil
}

func (u *User) GetImage() {
	url := "https://loremflickr.com/320/240?random=1"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	randStr, _ := crypto.MakeRandomStr(15)
	extension := ".png"
	filename := randStr + extension
	path := "upload/user/"

	file, err := os.Create(path + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, response.Body)
	u.Image = filename
}

func (u *User)FindLoginUser(DB *gorm.DB, email string, password string, organization string) error {
	cryptoPassword := crypto.Encrypt(password)
	result := DB.Preload("Organizations", "organization_id = ?", organization).Preload(clause.Associations).First(&u, "email = ? and password = ?", email, cryptoPassword)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		message := "email or password is wrong"
		err := errors.New(message)
		return err
	}

	return nil
}

func GetUserJson(r *http.Request) (User, error) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}