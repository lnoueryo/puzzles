package controller

import (
	"backend/models"
	"backend/modules/crypto"
	"backend/modules/image"
	"encoding/json"
	"net/http"
)

type Message struct {
	Name    string
	Message string
}

type MainUser struct {
	User models.User									`json:"user"`
	OrganizationAuthority models.OrganizationAuthority 	`json:"organizationAuthority"`
	Projects []models.Project							`json:"projects"`
}


type Home struct{}

// セッションが有効であるか確認
func (_ *Home) Show(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	// セッションを確認し、ユーザー情報を作成
	mainUser, err := CreateMainUser(r); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}
	uJson, _ := json.Marshal(mainUser)
	w.WriteHeader(http.StatusOK)
	w.Write(uJson)
}

// ユーザー情報の更新
func (h *Home) Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		errMap := map[string]string{"message": "not found"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(errJson)
		return
	}

	// リクエストで送られてきたユーザー情報のjsonを変換
	u, err := models.GetUserJson(r); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "bad connection"}
		errJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errJson)
		return
	}

	// イメージが変更されている場合
	if u.ImageData != "" {
		// 既存のイメージの名前を格納
		deleteImageName := u.Image

		// イメージを保存
		fileName, err := UploadToGCS("users", u.ImageData); if err != nil {
			errorlog.Print(err)
			errMap := map[string]string{"message": "couldn't save the image"}
			errJson, _ := json.Marshal(errMap)
			w.WriteHeader(http.StatusNotFound)
			w.Write(errJson)
			return
		}
		u.Image = fileName

		// 過去のイメージを削除
		if deleteImageName != "" {
			DeleteImage(deleteImageName, "users")
		}
	// イメージの変更がない場合
	} else {

		// 初めてのユーザー登録の場合
		if u.Image == "" {

			var nameLength uint32 = 20
			u.Image, _ = crypto.MakeRandomStr(nameLength)

			// 名前のイメージ作成
			buf, err := image.CreateImage(u.Name, u.Image); if err != nil {
				errorlog.Print(err)
				errMap := map[string]string{"message": "couldn't save the image"}
				errJson, _ := json.Marshal(errMap)
				w.WriteHeader(http.StatusNotFound)
				w.Write(errJson)
				return
			}

			// 環境による保存場所の変更
			path := "users/" + u.Image
			err = StoreImageToGCS(buf.Bytes(), path)
			if err != nil {
				errorlog.Print(err);
				errMap := map[string]string{"message": "couldn't save the image"}
				errJson, _ := json.Marshal(errMap)
				w.WriteHeader(http.StatusNotFound)
				w.Write(errJson)
				return
			}
		}
	}

	err = u.Update(); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}

	// 新しい変更ないようでセッションを作成
	mainUser, err := CreateMainUser(r); if err != nil {
		errorlog.Print(err)
		errMap := map[string]string{"message": "bad connection"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(sessionJson)
		return
	}
	uJson, _ := json.Marshal(mainUser)
	w.WriteHeader(http.StatusOK)
	w.Write(uJson)
}
