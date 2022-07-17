package services

import (
	"backend/models"
	"backend/modules/crypto"
	"backend/modules/image"
	"backend/modules/storage"
	"net/http"
)

type MainUser struct {
	User                  models.User                  `json:"user"`
	OrganizationAuthority models.OrganizationAuthority `json:"organizationAuthority"`
	Projects              []models.Project             `json:"projects"`
}

// 新たにセッション作成
func CreateMainUser(r *http.Request) (MainUser, error) {
	// BUG(inoueryo) 何かしらデータ更新を行う際に新しくセッションを作成し、古いものは削除されていない可能性あり
	var mainUser MainUser
	var u models.User
	s, err := GetSession(r)
	err = u.GetMainUser(s.UserID, s.Organization)
	if err != nil {
		return mainUser, err
	}

	mainUser.Projects = projectFilter(u.Organizations[0].Organization.Projects, func(userID int) bool {
		return userID == u.ID
	})
	u.Organizations[0].Organization.Projects = nil
	mainUser.OrganizationAuthority = u.Organizations[0]
	u.Organizations = nil
	mainUser.User = u
	return mainUser, nil
}

func UpdateMainUser(r *http.Request) error {
	// リクエストで送られてきたユーザー情報のjsonを変換
	u, err := models.GetUserJson(r); if err != nil {
		return err
	}

	// イメージが変更されている場合
	if u.ImageData != "" {
		// 既存のイメージの名前を格納
		deleteImageName := u.Image

		fileName, err := storage.UploadToGCS("users", u.ImageData); if err != nil {
			return err
		}
		u.Image = fileName

		// 過去のイメージを削除
		if deleteImageName != "" {
			storage.DeleteImage(deleteImageName, "users")
		}

	} else {

		// ユーザー登録の場合
		if u.Image == "" {

			var nameLength uint32 = 20
			u.Image, _ = crypto.MakeRandomStr(nameLength)

			// 名前のイメージ作成
			buf, err := image.CreateImage(u.Name, u.Image); if err != nil {
				return err
			}

			// 環境による保存場所の変更
			path := "users/" + u.Image
			err = storage.StoreImageToGCS(buf.Bytes(), path)
			if err != nil {
				return err
			}
		}
	}

	err = u.Update(); if err != nil {
		return err
	}

	return nil
}

// ユーザーが参加しているプロジェクトをフィルター
func projectFilter(projects []models.Project, condition func(int) bool) []models.Project {

	var mainUserProjects []models.Project
	for _, project := range projects {
		isUser := false
		for _, user := range project.AuthorityUsers {
			if condition(user.UserID) {
				isUser = true
				break
			}
		}

		if isUser {
			mainUserProjects = append(mainUserProjects, project)
		}
	}

	return mainUserProjects
}