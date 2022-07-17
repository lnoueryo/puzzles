package services

import (
	"backend/models"
	"errors"
	"net/http"
	"strconv"
)

func UpdateOrganizationAuthority(r *http.Request) error {
	oa, err := models.GetOrganizationAuthorityJson(r)
	if err != nil {
		return err
	}

	// 権限の変更
	err = oa.Update(DB)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOrganizationAuthority(r *http.Request) error {
	query := r.URL.Query()
    idSlice, ok := query["id"]; if !ok {
		message := "couldn't get id"
		err := errors.New(message)
		return err
    }
	var IDs []int
	for _, ID := range idSlice {
		id, err := strconv.Atoi(ID)
		if err != nil {
			return err
		}
		IDs = append(IDs, id)
	}

    oa, err := models.NewOrganizationAuthority(r)
    if err != nil {
        return err
    }

	err = oa.DeleteByUserIDs(DB, IDs); if err != nil {
		return err
	}

	return nil
}