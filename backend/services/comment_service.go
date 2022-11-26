package services

import (
	"backend/models"
	"net/http"
)

func GetComment(id int) ([]models.Comment, error) {

	c, err := models.GetCommentsByTaskID(DB, id); if err != nil {
		return c, err
	}

	return c, nil
}

func CreateComment(r *http.Request) (models.Comment, error) {
    c, err := models.NewComment(r);if err != nil {
        return c, err
    }
	err = c.Create(DB); if err != nil {
		return c, err
	}

	return c, nil
}

func UpdateComment(r *http.Request) (models.Comment, error) {

    c, err := models.NewComment(r);if err != nil {
        return c, err
    }

	err = c.Update(DB); if err != nil {
		return c, err
	}

	return c, nil
}

func DeleteComment(r *http.Request) error {

	ids, _, err := GetIDs(r);if err != nil {
		return err
	}

    c, err := models.NewComment(r);if err != nil {
        return err
    }

	err = c.Delete(DB, ids); if err != nil {
		return err
	}

	return nil
}