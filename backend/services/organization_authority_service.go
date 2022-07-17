package services

import (
	"backend/models"
	"net/http"
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