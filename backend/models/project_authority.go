package models

import (
	// "backend/modules/image"
	// "os"
	"encoding/json"
	"errors"
	"fmt"
	// "io/ioutil"
	"net/http"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProjectAuthority struct {
	ID				int					`gorm:"AUTO_INCREMENT"json:"id"`
    ProjectID		int					`json:"project_id"`
    UserID			int					`json:"user_id"`
    AuthorityID		int					`json:"auth_id"`
	Active			bool				`json:"active"`
	Type			Authority			`gorm:"foreignkey:AuthorityID;"json:"type"`
	User			User				`gorm:"foreignkey:UserID;"json:"user"`
	Project			Project				`gorm:"foreignkey:ProjectID;"json:"project"`
	CreatedAt		time.Time			`gorm:"<-:create;autoCreateTime;"json:"created_at"`
	UpdatedAt		time.Time			`json:"-"`
}

func NewProjectAuthority(r *http.Request) (ProjectAuthority, error) {
	projectAuthority, _ := GetProjectAuthorityJson(r)
	return projectAuthority, nil
}

func GetProjectAuthority(DB *gorm.DB, pid int, uid int) (ProjectAuthority, error) {
	var pa ProjectAuthority
	result := DB.Preload("Project.AuthorityUsers.Type").Preload("Project.AuthorityUsers.User").Preload("Project.AuthorityUsers", "active = ?", true).Preload("Project." + clause.Associations).Preload(clause.Associations).First(&pa, "project_id = ? and user_id = ?", pid, uid); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return pa, result.Error
	}
	return pa, nil
}

func (pa *ProjectAuthority)Create(DB *gorm.DB) error {
	result := DB.Omit("User", "Type", "Project").Create(&pa); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (pa *ProjectAuthority)Update(DB *gorm.DB) error {
	result := DB.Omit("User", "Type", "Project").Save(&pa); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (pa *ProjectAuthority)DeleteByUserIDs(DB *gorm.DB, ids []int) error {
	result := DB.Delete(&pa, ids); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func FindProjectAuthority(DB *gorm.DB, pid int, uid int) (ProjectAuthority, error) {
	var pa ProjectAuthority
	result := DB.Preload(clause.Associations).First(&pa, "project_id = ? and user_id = ? and active = true", pid, uid); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return pa, result.Error
	}
	return pa, nil
}

func GetProjectAuthorityJson(r *http.Request) (ProjectAuthority, error) {
	var projectAuthority ProjectAuthority
	err := json.NewDecoder(r.Body).Decode(&projectAuthority)
	if err != nil {
		fmt.Println(err)
	}
	return projectAuthority, nil
}
