package services

import (
	"backend/models"
	"net/http"
)

func UpdateOrganization(r *http.Request) error {
	o, err := models.GetOrganizationJson(r)
	if err != nil {
		return err
	}

	err = o.Update()
	if err != nil {
		return err
	}

	return nil
}