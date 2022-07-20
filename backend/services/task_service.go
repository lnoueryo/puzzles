package services

import (
	"backend/models"
	"net/http"
	"strconv"
)

func GetTask(id int) ([]models.Task, error) {

	t, err := models.GetTasksByProjectID(DB, id); if err != nil {
		return t, err
	}

	return t, nil
}

func CreateTask(r *http.Request) (models.Task, error) {
    t, err := models.NewTask(r);if err != nil {
        return t, err
    }

	project, err := t.CountProjectTask(DB); if err != nil {
        return t, err
    }

	t.Key = project.Name + "_" + strconv.Itoa(len(project.Tasks) + 1)
	err = t.Create(DB); if err != nil {
		return t, err
	}
	return t, nil
}

func UpdateTask(r *http.Request) (models.Task, error) {
    t, err := models.NewTask(r);if err != nil {
        return t, err
    }

	err = t.Update(DB); if err != nil {
		return t, err
	}
	return t, nil
}
