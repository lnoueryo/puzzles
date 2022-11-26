package services

import (
	"backend/config"
	"errors"
	"net/http"
	"strconv"
)

var DB = config.DB
var projectID = config.App.ProjectID
var email = config.App.Email
var origin = config.App.Host
var allowOrigin = config.App.AllowOrigin
var infolog = config.App.InfoLog
var cookieKey = config.App.CookieKey
var bucketName = config.BUCKET_NAME

func GetIDs(r *http.Request) ([]int, int, error) {
	var ids []int
	page := -1
	query := r.URL.Query()
    idSlice, ok := query["id"]; if !ok {
		message := "couldn't get id"
		err := errors.New(message)
		return ids, page, err
    }

	for _, ID := range idSlice {
		id, err := strconv.Atoi(ID);if err != nil {
			return ids, page, err
		}
		ids = append(ids, id)
	}

    pageSlice, ok := query["page"]; if ok {
		intPage, err := strconv.Atoi(pageSlice[0]);if err != nil {
			return ids, page, err
		}
		page = intPage
    }

	return ids, page, nil
}