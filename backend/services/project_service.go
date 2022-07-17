package services

import (
	"backend/models"
	"backend/modules/storage"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type ProjectUpdateRequest struct {
	Project			models.Project	`json:"project"`
	FieldDelete		bool			`json:"field_delete"`
	MilestoneDelete	bool			`json:"milestone_delete"`
	VersionDelete	bool			`json:"versione_delete"`
}


func CreateProject(r *http.Request) error {

    p, err := models.NewProject(r);if err != nil {
        return err
    }

	// イメージが更新された場合
	if p.ImageData != "" {
		fileName, err := storage.UploadToGCS("projects", p.ImageData); if err != nil {
			return err
		}
		p.Image = fileName
	}

	err = p.Create(DB); if err != nil {
		return err
	}

	return nil
}

func UpdateProject(r *http.Request) error {

    pur, err := getProjectUpdateRequestJson(r);if err != nil {
        return err
    }

	err = pur.BulkUpdateProject(); if err != nil {
		return err
	}

	return nil
}


func getProjectUpdateRequestJson(r *http.Request) (ProjectUpdateRequest, error) {
	var projectUpdateRequest ProjectUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&projectUpdateRequest);if err != nil {
		return projectUpdateRequest, err
	}
	return projectUpdateRequest, nil
}

// プロジェクトの内容を変更する関数
// プロジェクトのみにかかわらずフィールド、ヴァージョン、マイルストーン、も変更する
func (pur *ProjectUpdateRequest)BulkUpdateProject() error {

	err := DB.Transaction(func(tx *gorm.DB) error {

		if pur.FieldDelete {
			var f models.Field
			if err := f.DeleteByProjectID(tx, pur.Project.ID); err != nil {
				return err
			}
		}

		if pur.MilestoneDelete {

			var m models.Milestone
			if err := m.DeleteByProjectID(tx, pur.Project.ID); err != nil {
				return err
			}
		}

		if pur.VersionDelete {

			var v models.Version
			if err := v.DeleteByProjectID(tx, pur.Project.ID); err != nil {
				return err
			}
		}

		// イメージを変更した場合
		if pur.Project.ImageData != "" {

			deleteImageName := pur.Project.Image
			fileName, err := storage.UploadToGCS("projects", pur.Project.ImageData); if err != nil {
				return err
			}

			pur.Project.Image = fileName
			if err = pur.Project.Update(tx);err != nil {
				return err
			}

			storage.DeleteImage(deleteImageName, "projects")
			return nil
		}
		// 変更しないカラム
		if err := pur.Project.Update(tx);err != nil {
			return err
		}
		return nil
	})
	return err
}
