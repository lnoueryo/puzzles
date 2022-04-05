package controller

import (
	"backend/models"
	"backend/modules/csv"
	"encoding/json"
	"errors"
	"net/http"
)



type Data struct{}

var (
	csvNameList = []string{
		"activities.csv",
		"activity_contents.csv",
		"authorities.csv",
		"comments.csv",
		"fields.csv",
		"milestones.csv",
		"organization_authorities.csv",
		"organizations.csv",
		"priorities.csv",
		"project_authorities.csv",
		"projects.csv",
		"statuses.csv",
		"tasks.csv",
		"types.csv",
		"users.csv",
	}
)

func(*Data)Download(w http.ResponseWriter, r *http.Request) {
	var activities []models.Activity
	err := DB.Find(&activities).Error; if err != nil {
		errorlog.Print("Activityテーブルの読み込みに失敗しました。")
	}
	b := csv.DLActivity(activities)
	w.WriteHeader(http.StatusOK)
	w.Write(b.Bytes())
	// io.Copy(w, b)
}

func (*Data) Upload(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		errMap := map[string]string{"message": "not found"}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(sessionJson)
		return
	}
	err := r.ParseMultipartForm(20000000) // grab the multipart form
	if err != nil {
		return
	}

	formdata := r.MultipartForm // ok, no problem so far, read the Form data
	//get the *fileheaders
	files := formdata.File["files"] // grab the filenames
	for _, file := range files {
		f, err := file.Open()
		defer f.Close()
		if err != nil {
			return
		}
		if file.Filename == "activities.csv" {
			activities := csv.UpsertActivity(f);
			err := DB.Save(&activities).Error; if err != nil {
				errorlog.Print("Activityテーブルの書き込みに失敗しました。")
			}
			continue
		}
		if file.Filename == "activity_contents.csv" {
			activity_contents := csv.UpsertActivityContent(f)
			err := DB.Save(&activity_contents).Error; if err != nil {
				errorlog.Print("ActivityContentテーブルの書き込みに失敗しました。")
			}
			continue
		}
		if file.Filename == "authorities.csv" {
			authorities := csv.UpsertAuthority(f);
			err := DB.Save(&authorities).Error; if err != nil {
				errorlog.Print("Authorityテーブルの書き込みに失敗しました。")
			}
			continue
		}
		if file.Filename == "comments.csv" {
			comments := csv.UpsertComment(f);
			err := DB.Save(comments).Error; if err != nil {
				errorlog.Print("Authorityテーブルの書き込みに失敗しました。")
			}
			continue
		}
		if file.Filename == "fields.csv" {
			fields := csv.UpsertField(f);
			err := DB.Save(fields).Error; if err != nil {
				errorlog.Print("Fieldテーブルの書き込みに失敗しました。")
			}
			continue
		}
		if file.Filename == "milestones.csv" {
			milestones := csv.UpsertMilestone(f);
			err := DB.Save(&milestones).Error; if err != nil {
				errorlog.Print("Milestoneテーブルの書き込みに失敗しました。")
			}
			continue
		}
		if file.Filename == "organization_authorities.csv" {
			organizationAuthorities := csv.UpsertOrganizationAuthority(f);
			bulk := []models.OrganizationAuthority{}
			for _, record := range organizationAuthorities {
				bulk = append(bulk, record)
				if len(bulk) > 500 {
					err := DB.Save(bulk).Error; if err != nil {
						errorlog.Print("OrganizationAuthorityテーブルの書き込みに失敗しました。")
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					errorlog.Print("OrganizationAuthorityテーブルの書き込みに失敗しました。")
				}
			}
			continue
		}
		if file.Filename == "organizations.csv" {
			organizations := csv.UpsertOrganization(f);
			bulk := []models.Organization{}
			for _, record := range organizations {
				bulk = append(bulk, record)
				if len(bulk) > 500 {
					err := DB.Save(bulk).Error; if err != nil {
						errorlog.Print("OrganizationAuthorityテーブルの書き込みに失敗しました。")
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					errorlog.Print("OrganizationAuthorityテーブルの書き込みに失敗しました。")
				}
			}
			continue
		}
		if file.Filename == "priorities.csv" {
			priorities := csv.UpsertPriority(f);
			err := DB.Save(priorities).Error; if err != nil {
				errorlog.Print("Priorityテーブルの書き込みに失敗しました。")
			}
			continue
		}
		if file.Filename == "project_authorities.csv" {
			projectAuthorities := csv.UpsertProjectAuthority(f);
			bulk := []models.ProjectAuthority{}
			for _, record := range projectAuthorities {
				bulk = append(bulk, record)
				if len(bulk) > 500 {
					err := DB.Save(bulk).Error; if err != nil {
						errorlog.Print("ProjectAuthorityテーブルの書き込みに失敗しました。")
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					errorlog.Print("ProjectAuthorityテーブルの書き込みに失敗しました。")
				}
			}
			continue
		}
		if file.Filename == "projects.csv" {
			projects := csv.UpsertProject(f);
			bulk := []models.Project{}
			for _, record := range projects {
				bulk = append(bulk, record)
				if len(bulk) > 500 {
					err := DB.Save(bulk).Error; if err != nil {
						errorlog.Print("Projectテーブルの書き込みに失敗しました。")
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					errorlog.Print("Projectテーブルの書き込みに失敗しました。")
				}
			}
			continue
		}
		if file.Filename == "statuses.csv" {
			statuses := csv.UpsertStatus(f);
			err := DB.Save(statuses).Error; if err != nil {
				errorlog.Print("Statusテーブルの書き込みに失敗しました。")
			}
			continue
		}
		if file.Filename == "tasks.csv" {
			tasks := csv.UpsertTask(f);
			bulk := []models.Task{}
			for _, record := range tasks {
				bulk = append(bulk, record)
				if len(bulk) > 500 {
					err := DB.Save(bulk).Error; if err != nil {
						errorlog.Print("Taskテーブルの書き込みに失敗しました。")
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					errorlog.Print("Taskテーブルの書き込みに失敗しました。")
				}
			}
			continue
		}
		if file.Filename == "types.csv" {
			types := csv.UpsertType(f);
			err := DB.Save(types).Error; if err != nil {
				errorlog.Print("Typeテーブルの書き込みに失敗しました。")
			}
			continue
		}
		if file.Filename == "users.csv" {
			users := csv.UpsertUser(f);
			bulk := []models.User{}
			for _, record := range users {
				bulk = append(bulk, record)
				if len(bulk) > 500 {
					err := DB.Save(bulk).Error; if err != nil {
						errorlog.Print("Userテーブルの書き込みに失敗しました。")
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					errorlog.Print("Userテーブルの書き込みに失敗しました。")
				}
			}
			continue
		}
		errorlog.Print(file.Filename + "は除外されました")
	}
	org := map[string]string{"message": "fail:"}
	orgJson, _ := json.Marshal(org)
	w.WriteHeader(http.StatusOK)
	w.Write(orgJson)
}


func GetJson(r *http.Request) (Project, error) {
	var project Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		message := "couldn't decode json"
		err := errors.New(message)
		return project, err
	}
	return project, nil
}