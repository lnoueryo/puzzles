package models

import (
	"backend/modules/image"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
	"gorm.io/gorm"
)

type ProjectAuthority struct {
	ID				int					`gorm:"AUTO_INCREMENT"json:"id"`
    ProjectID		int					`json:"project_id"`
    UserID			int					`json:"user_id"`
    AuthorityID		int					`json:"auth_id"`
	Active			bool				`json:"active"`
	Type			Authority			`gorm:"foreignkey:AuthorityID;migrate"json:"type"`
	User			User				`gorm:"foreignkey:UserID;migrate"json:"user"`
	Project			Project				`gorm:"foreignkey:ProjectID;migrate"json:"project"`
	ProjectUsers	[]ProjectAuthority	`gorm:"foreignkey:ProjectID;references:ProjectID;migrate;"json:"project_users"`
	CreatedAt		time.Time			`gorm:"<-:create;autoCreateTime;"json:"-"`
	UpdatedAt		time.Time			`json:"-"`
}

type ProjectAuthorityRequest struct {
	ProjectAuthority	ProjectAuthority `json:"project_authority"`
	FieldDelete			bool			 `json:"field_delete"`
	MilestoneDelete		bool			 `json:"milestone_delete"`
}

func NewProjectAuthority(r *http.Request) (ProjectAuthority, error) {
	projectAuthority, _ := GetProjectAuthorityJson(r)
	return projectAuthority, nil
}

func (pa *ProjectAuthority)Update() error {
	var projectAuthorities []ProjectAuthority
	for _, user := range pa.ProjectUsers {
		projectAuthority := ProjectAuthority{
			ID: user.ID,
			ProjectID: user.ProjectID,
			UserID: user.UserID,
			AuthorityID: user.AuthorityID,
			Active: user.Active,
		}
		projectAuthorities = append(projectAuthorities, projectAuthority)
	}
	result := DB.Save(&projectAuthorities); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	result = DB.Save(&pa.Project); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (par *ProjectAuthorityRequest)BulkUpdateProject() error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if par.FieldDelete {
			var f Field
			if err := tx.Delete(&f, "project_id = ?", par.ProjectAuthority.ProjectID).Error; err != nil {
				return errors.New("couldn't delete")
			}
		}
		if par.MilestoneDelete {
			var m Milestone
			if err := tx.Delete(&m, "project_id = ?", par.ProjectAuthority.ProjectID).Error; err != nil {
				return errors.New("couldn't delete")
			}
		}
		if par.ProjectAuthority.Project.ImageData != "" {
			fileName, err := image.StoreImage("projects", par.ProjectAuthority.Project.ImageData); if err != nil {
				return errors.New("couldn't save the image")
			}
			path := "upload/projects/" + par.ProjectAuthority.Project.Image
			par.ProjectAuthority.Project.Image = fileName
			result := tx.Debug().Omit("Fields", "Milestones", "Users", "AuthorityUsers.User", "AuthorityUsers.Type").Session(&gorm.Session{FullSaveAssociations: true}).Save(&par.ProjectAuthority.Project); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return result.Error
			}
			os.Remove(path)
			return nil
		}
		result := tx.Debug().Omit("Fields", "Milestones", "Users", "AuthorityUsers.User", "AuthorityUsers.Type").Session(&gorm.Session{FullSaveAssociations: true}).Save(&par.ProjectAuthority.Project); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
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

func GetProjectAuthorityRequestJson(r *http.Request) (ProjectAuthorityRequest, error) {
	var projectAuthorityRequest ProjectAuthorityRequest
	err := json.NewDecoder(r.Body).Decode(&projectAuthorityRequest)
	if err != nil {
		message := "couldn't decode json"
		err := errors.New(message)
		return projectAuthorityRequest, err
	}
	return projectAuthorityRequest, nil
}