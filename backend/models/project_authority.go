package models

import (
	// "backend/modules/image"
	// "os"
	"encoding/json"
	"errors"
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
	CreatedAt		time.Time			`gorm:"<-:create;autoCreateTime;"json:"-"`
	UpdatedAt		time.Time			`json:"-"`
}

func NewProjectAuthority(r *http.Request) (ProjectAuthority, error) {
	projectAuthority, _ := GetProjectAuthorityJson(r)
	return projectAuthority, nil
}

func GetProjectAuthority(pid int, uid int) (ProjectAuthority, error) {
	var pa ProjectAuthority
	result := DB.Preload("Project.AuthorityUsers.Type").Preload("Project.AuthorityUsers.User").Preload("Project.AuthorityUsers", "active = ?", true).Preload("Project." + clause.Associations).Preload(clause.Associations).First(&pa, "project_id = ? and user_id = ?", pid, uid); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return pa, result.Error
	}
	return pa, nil
}

func (pa *ProjectAuthority)Update() error {
	var projectAuthorities []ProjectAuthority
	result := DB.Save(&projectAuthorities); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	result = DB.Save(&pa.Project); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (pur *ProjectUpdateRequest)BulkUpdateProject() error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if pur.FieldDelete {
			var f Field
			if err := tx.Delete(&f, "project_id = ?", pur.Project.ID).Error; err != nil {
				return errors.New("couldn't delete")
			}
		}
		if pur.MilestoneDelete {
			var m Milestone
			if err := tx.Delete(&m, "project_id = ?", pur.Project.ID).Error; err != nil {
				return errors.New("couldn't delete")
			}
		}
		if pur.Project.ImageData != "" {
			fileName, err := StoreImage("projects", pur.Project.ImageData); if err != nil {
				return errors.New("couldn't save the image")
			}
			pur.Project.Image = fileName
			result := tx.Debug().Omit("Organization", "Tasks", "AuthorityUsers.User", "AuthorityUsers.Type", "AuthorityUsers.Project").Session(&gorm.Session{FullSaveAssociations: true}).Save(&pur.Project); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return result.Error
			}
			DeleteImage("9MSwg3CWQlIQYiVNkZcK.jpeg", "projects")
			return nil
		}
		result := tx.Omit("Organization", "Tasks", "AuthorityUsers.User", "AuthorityUsers.Type", "AuthorityUsers.Project").Session(&gorm.Session{FullSaveAssociations: true}).Save(&pur.Project); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		return nil
	})
	return err
}

func GetProjectAuthorityJson(r *http.Request) (ProjectAuthority, error) {
	var projectAuthority ProjectAuthority
	err := json.NewDecoder(r.Body).Decode(&projectAuthority)
	if err != nil {
		message := "couldn't decode json"
		err := errors.New(message)
		return projectAuthority, err
	}
	return projectAuthority, nil
}
