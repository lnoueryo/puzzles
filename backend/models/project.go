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



type Project struct {
	ID				int				`gorm:"AUTO_INCREMENT"json:"id"`
    OrganizationID	string			`gorm:"<-:create;"json:"organization_id"`
	Name			string			`json:"name"`
	Description		string			`json:"description"`
	Image			string			`json:"image"`
	ImageData		string			`gorm:"migration;"json:"image_data"`
	Authority		string			`gorm:"migration"json:"authority"`
	Organization	Organization	`gorm:"->;references:ID;"json:"organization"`
	Tasks			[]Task			`json:"tasks"`
	Milestones		[]Milestone		`json:"milestones"`
	Fields			[]Field			`json:"fields"`
	AuthorityUsers  []ProjectAuthority	`json:"authority_users"`
	Users			[]User			`gorm:"many2many:project_authorities;"json:"users"`
	CreatedAt		time.Time		`gorm:"autoCreateTime;"json:"created_at"`
	UpdatedAt		time.Time		`gorm:"autoUpdateTime;"json:"updated_at"`
}

type ProjectRequest struct {
	Project
	ImageData		string			`gorm:"<-:false;migration;"json:"image_data"`
	Authority		string			`gorm:"migration"json:"authority"`
}

func NewProject(r *http.Request) (Project, error) {
	project, _ := GetProjectJson(r)
	return project, nil
}

func (p *Project)GetProjectAuthority(s Session) (ProjectAuthority, error) {
	fmt.Println(p)
	var pa ProjectAuthority
	result := DB.Preload("Project.Milestones").Preload("Project.Fields").Preload("ProjectUsers.Type").Preload("ProjectUsers.User").Preload("ProjectUsers").Preload(clause.Associations).Find(&pa, "user_id = ? and project_id = ?", s.UserID, p.ID); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return pa, result.Error
	}
	return pa, nil
}

func(p *Project)Create() error {
	DB.SetupJoinTable(&Project{}, "Users", &ProjectAuthority{})
	result := DB.Debug().Omit("Users.*").Create(&p); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (p *Project)Update() error {
	var projectAuthorities []ProjectAuthority
	for _, user := range p.AuthorityUsers {
		projectAuthority := ProjectAuthority{
			ID: user.ID,
			ProjectID: user.ProjectID,
			UserID: user.UserID,
			AuthorityID: user.AuthorityID,
			Active: user.Active,
		}
		projectAuthorities = append(projectAuthorities, projectAuthority)
	}
	result := DB.Debug().Updates(&projectAuthorities); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	result = DB.Debug().Updates(&p); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
	// var projectAuthorities []ProjectAuthority
	// for _, user := range p.AuthorityUsers {
	// 	projectAuthority := ProjectAuthority{
	// 		ID: user.ID,
	// 		ProjectID: user.ProjectID,
	// 		UserID: user.UserID,
	// 		AuthorityID: user.AuthorityID,
	// 		Active: user.Active,
	// 	}
	// 	projectAuthorities = append(projectAuthorities, projectAuthority)
	// 	fmt.Println(user.AuthorityID)
	// 	fmt.Println(user.User.Name)
	// }
	// result := DB.Save(projectAuthorities); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	return result.Error
	// }
	// result = DB.Save(&p); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	return result.Error
	// }
}

func (p *Project)GetProject(s Session, id int) error {
	result := DB.Preload("Tasks", func(DB *gorm.DB) *gorm.DB {
		return DB.Preload(clause.Associations)
	  }).Preload(clause.Associations).First(&p, "id = ?", id); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	var pas []ProjectAuthority
	result = DB.Preload(clause.Associations).Find(&pas, "project_id = ? and active = true", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	for _ ,pa := range pas {
		if s.UserID == pa.User.ID {
			p.Authority = pa.Type.Name
		}
	}
	p.AuthorityUsers = pas
	return nil
}

func(p *Project)GetEditProject(s Session, id int) error {
	result := DB.Preload(clause.Associations).First(&p, "id = ?", id); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	var pas []ProjectAuthority
	result = DB.Preload(clause.Associations).Find(&pas, "project_id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	for _ ,pa := range pas {
		if s.UserID == pa.User.ID {
			p.Authority = pa.Type.Name
		}
	}
	p.AuthorityUsers = pas
	return nil
}

func (pa *ProjectAuthority) BeforeSave(DB *gorm.DB) error {
	if pa.AuthorityID == 0 {
		pa.AuthorityID = 1
	}
	fmt.Println(pa)
    return nil
}

func GetProjectJson(r *http.Request) (Project, error) {
	var project Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		message := "couldn't decode json"
		err := errors.New(message)
		return project, err
	}
	return project, nil
}

func (p *Project) GetImage() {
	url := "https://loremflickr.com/320/240?random=1"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	randStr, _ := crypto.MakeRandomStr(15)
	extension := ".png"
	filename := randStr + extension
	path := "upload/projects/"

	file, err := os.Create(path + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, response.Body)
	p.Image = filename
}