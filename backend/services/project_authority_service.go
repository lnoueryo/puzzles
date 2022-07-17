package services

import (
	"backend/models"
	"errors"
	"net/http"
	"strconv"
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
			return nil
		}
		IDs = append(IDs, id)
	}

	pa, err := models.NewProjectAuthority(r);if err != nil {
		return err
	}

	err = pa.DeleteByUserIDs(DB, IDs); if err != nil {
		return err
	}
	return nil
}