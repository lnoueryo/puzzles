package services

import (
	"backend/models"
	"net/http"
)

func CreateProjectAuthority(r *http.Request) error {
    pa, err := models.GetProjectAuthorityJson(r);if err != nil {
        return err
    }

	err = pa.Create(DB); if err != nil {
		return err
	}
	return nil
}

func UpdateProjectAuthority(r *http.Request) error {
    pa, err := models.GetProjectAuthorityJson(r);if err != nil {
        return err
    }

	err = pa.Update(DB); if err != nil {
		return err
	}
	return nil
}

func DeleteProjectAuthority(r *http.Request) error {

	ids, err := GetIDs(r);if err != nil {
		return err
	}

	pa, err := models.NewProjectAuthority(r);if err != nil {
		return err
	}

	err = pa.DeleteByUserIDs(DB, ids); if err != nil {
		return err
	}
	return nil
}

func FindProjectAuthority(pid int, uid int) bool {
	_, err := models.FindProjectAuthority(DB, pid, uid);if err != nil {
		return false
	}
	return true
}