package controller

import (
	"backend/models"
	dp "backend/modules/processing"
	"encoding/json"
	"errors"
	"io/ioutil"
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

type CSVRequest struct {
	Request []string `json:"request"`
}

func(*Data)Index(w http.ResponseWriter, r *http.Request) {
	
}

func(*Data)Download(w http.ResponseWriter, r *http.Request) {
	byteArray, _ := ioutil.ReadAll(r.Body)
	var csvRequest CSVRequest
	err := json.Unmarshal(byteArray, &csvRequest); if err != nil{
		errorlog.Print(err)
		errMap := map[string]string{"message": err.Error()}
		sessionJson, _ := json.Marshal(errMap)
		w.WriteHeader(http.StatusNotFound)
		w.Write(sessionJson)
		return
	}
	var csvBufferArray []dp.CSVBuffer
	for _, request := range csvRequest.Request {
		if request == "activities" {
			var activities []models.Activity
			err := DB.Find(&activities).Error; if err != nil {
				errorlog.Print("Activityテーブルの読み込みに失敗しました。")
			}
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLActivity(activities),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "activity_contents" {
			var activity_contents []models.ActivityContent
			err := DB.Find(&activity_contents).Error; if err != nil {
				errorlog.Print("ActivityContentテーブルの読み込みに失敗しました。")
			}
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLActivityContent(activity_contents),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "authorities" {
			var authorities []models.Authority
			err := DB.Find(&authorities).Error; if err != nil {
				errorlog.Print("ActivityContentテーブルの読み込みに失敗しました。")
			}
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLAuthority(authorities),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "comments" {
			var comments []models.Comment
			err := DB.Find(&comments).Error; if err != nil {
				errorlog.Print("Commentテーブルの読み込みに失敗しました。")
			}
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLComment(comments),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "fields" {
			var fields []models.Field
			err := DB.Find(&fields).Error; if err != nil {
				errorlog.Print("Fieldテーブルの読み込みに失敗しました。")
			}
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLField(fields),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "milestones" {
			var milestones []models.Milestone
			err := DB.Find(&milestones).Error; if err != nil {
				errorlog.Print("Milestoneテーブルの読み込みに失敗しました。")
			}
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLMilestone(milestones),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "organization_authorities" {
			var organization_authorities []models.OrganizationAuthority
			err := DB.Find(&organization_authorities).Error; if err != nil {
				errorlog.Print("OrganizationAuthorityテーブルの読み込みに失敗しました。")
			}
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLOrganizationAuthority(organization_authorities),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "organizations" {
			var organizations []models.Organization
			err := DB.Find(&organizations).Error; if err != nil {
				errorlog.Print("Organizationテーブルの読み込みに失敗しました。")
			}
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLOrganization(organizations),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "priorities" {
			var priorities []models.Priority
			err := DB.Find(&priorities).Error; if err != nil {
				errorlog.Print("Priorityテーブルの読み込みに失敗しました。")
			}
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLPriority(priorities),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "project_authorities" {
			var project_authorities []models.ProjectAuthority
			err := DB.Find(&project_authorities).Error; if err != nil {
				errorlog.Print("ProjectAuthorityテーブルの読み込みに失敗しました。")
			}
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLProjectAuthority(project_authorities),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "projects" {
			var projects []models.Project
			err := DB.Find(&projects).Error; if err != nil {
				errorlog.Print("Projectテーブルの読み込みに失敗しました。")
			}
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLProject(projects),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "statuses" {
			var statuses []models.Status
			err := DB.Find(&statuses).Error; if err != nil {
				errorlog.Print("Statusテーブルの読み込みに失敗しました。")
			}
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLStatus(statuses),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "tasks" {
			var tasks []models.Task
			err := DB.Find(&tasks).Error; if err != nil {
				errorlog.Print("Taskテーブルの読み込みに失敗しました。")
			}
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLTask(tasks),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "types" {
			var types []models.Type
			err := DB.Find(&types).Error; if err != nil {
				errorlog.Print("Typeテーブルの読み込みに失敗しました。")
			}
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLType(types),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "users" {
			var users []dp.User
			err := DB.Find(&users).Error; if err != nil {
				errorlog.Print("Userテーブルの読み込みに失敗しました。")
			}
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLUser(users),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
	}
	b := dp.ZipCSVByteArray(&csvBufferArray)
	w.WriteHeader(http.StatusOK)
	w.Write(b.Bytes())
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
	var errorArray [][]string
	for _, file := range files {
		f, err := file.Open()
		defer f.Close()
		if err != nil {
			return
		}
		if file.Filename == "activities.csv" {
			activities, errArray := dp.UpsertActivity(f)
			err := DB.Save(&activities).Error; if err != nil {
				message := "Activityテーブルの書き込みに失敗しました。"
				errorlog.Print(message)
				errorlog.Print(err)
				errArray = append(errArray, message)
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		if file.Filename == "activity_contents.csv" {
			activity_contents, errArray := dp.UpsertActivityContent(f)
			err := DB.Save(&activity_contents).Error; if err != nil {
				message := "ActivityContentテーブルの書き込みに失敗しました。"
				errorlog.Print(message)
				errorlog.Print(err)
				errArray = append(errArray, message)
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		if file.Filename == "authorities.csv" {
			authorities, errArray := dp.UpsertAuthority(f);
			err := DB.Save(&authorities).Error; if err != nil {
				message := "Authorityテーブルの書き込みに失敗しました。"
				errorlog.Print(message)
				errorlog.Print(err)
				errArray = append(errArray, message)
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		if file.Filename == "comments.csv" {
			comments, errArray := dp.UpsertComment(f);
			var bulk []models.Comment
			for _, comment := range comments {
				bulk = append(bulk, comment)
				if len(bulk) > 500 {
					err := DB.Save(&bulk).Error; if err != nil {
						message := "Commentテーブルの書き込みに失敗しました。"
						errorlog.Print(message)
						errorlog.Print(err)
						errArray = append(errArray, message)
					}
					if len(errArray) != 0 {
						errorArray = append(errorArray, errArray)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(&bulk).Error; if err != nil {
					message := "Commentテーブルの書き込みに失敗しました。"
					errorlog.Print(message)
					errorlog.Print(err)
					errArray = append(errArray, message)
				}
				if len(errArray) != 0 {
					errorArray = append(errorArray, errArray)
				}
			}
			continue
		}
		if file.Filename == "fields.csv" {
			fields, errArray := dp.UpsertField(f);
			bulk := []models.Field{}
			for _, record := range fields {
				bulk = append(bulk, record)
				if len(bulk) > 500 {
					err := DB.Save(bulk).Error; if err != nil {
						message := "Fieldテーブルの書き込みに失敗しました。"
						errorlog.Print(message)
						errorlog.Print(err)
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "Fieldテーブルの書き込みに失敗しました。"
					errorlog.Print(message)
					errorlog.Print(err)
					errArray = append(errArray, message)
				}
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		if file.Filename == "milestones.csv" {
			milestones, errArray := dp.UpsertMilestone(f);
			bulk := []models.Milestone{}
			for _, record := range milestones {
				bulk = append(bulk, record)
				if len(bulk) > 500 {
					err := DB.Save(bulk).Error; if err != nil {
						message := "Milestoneテーブルの書き込みに失敗しました。"
						errorlog.Print(message)
						errorlog.Print(err)
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "Milestoneテーブルの書き込みに失敗しました。"
					errorlog.Print(message)
					errorlog.Print(err)
					errArray = append(errArray, message)
				}
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		if file.Filename == "organization_authorities.csv" {
			organizationAuthorities, errArray := dp.UpsertOrganizationAuthority(f);
			bulk := []models.OrganizationAuthority{}
			for _, record := range organizationAuthorities {
				bulk = append(bulk, record)
				if len(bulk) > 500 {
					err := DB.Save(bulk).Error; if err != nil {
						message := "OrganizationAuthorityテーブルの書き込みに失敗しました。"
						errorlog.Print(message)
						errorlog.Print(err)
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "OrganizationAuthorityテーブルの書き込みに失敗しました。"
					errorlog.Print(message)
					errorlog.Print(err)
					errArray = append(errArray, message)
				}
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		if file.Filename == "organizations.csv" {
			organizations, errArray := dp.UpsertOrganization(f);
			bulk := []models.Organization{}
			for _, record := range organizations {
				bulk = append(bulk, record)
				if len(bulk) > 500 {
					err := DB.Save(bulk).Error; if err != nil {
						message := "Organizationテーブルの書き込みに失敗しました。"
						errorlog.Print(message)
						errorlog.Print(err)
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "Organizationテーブルの書き込みに失敗しました。"
					errorlog.Print(message)
					errorlog.Print(err)
					errArray = append(errArray, message)
				}
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		if file.Filename == "priorities.csv" {
			priorities, errArray := dp.UpsertPriority(f);
			err := DB.Save(priorities).Error; if err != nil {
				message := "Priorityテーブルの書き込みに失敗しました。"
				errorlog.Print(message)
				errorlog.Print(err)
				errArray = append(errArray, message)
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		if file.Filename == "project_authorities.csv" {
			projectAuthorities, errArray := dp.UpsertProjectAuthority(f);
			bulk := []models.ProjectAuthority{}
			for _, record := range projectAuthorities {
				bulk = append(bulk, record)
				if len(bulk) > 500 {
					err := DB.Save(bulk).Error; if err != nil {
						message := "ProjectAuthorityテーブルの書き込みに失敗しました。"
						errorlog.Print(message)
						errorlog.Print(err)
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "ProjectAuthorityテーブルの書き込みに失敗しました。"
					errorlog.Print(message)
					errorlog.Print(err)
					errArray = append(errArray, message)
				}
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		if file.Filename == "projects.csv" {
			projects, errArray := dp.UpsertProject(f);
			bulk := []dp.Project{}
			for _, record := range projects {
				bulk = append(bulk, record)
				if len(bulk) > 500 {
					err := DB.Save(bulk).Error; if err != nil {
						message := "Projectテーブルの書き込みに失敗しました。"
						errorlog.Print(message)
						errorlog.Print(err)
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "Projectテーブルの書き込みに失敗しました。"
					errorlog.Print(message)
					errorlog.Print(err)
					errArray = append(errArray, message)
				}
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		if file.Filename == "statuses.csv" {
			statuses, errArray := dp.UpsertStatus(f);
			err := DB.Save(statuses).Error; if err != nil {
				message := "Statusテーブルの書き込みに失敗しました。"
				errorlog.Print(message)
				errorlog.Print(err)
				errArray = append(errArray, message)
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		if file.Filename == "tasks.csv" {
			tasks, errArray := dp.UpsertTask(f);
			bulk := []models.Task{}
			for _, record := range tasks {
				bulk = append(bulk, record)
				if len(bulk) > 300 {
					err := DB.Save(bulk).Error; if err != nil {
						message := "Taskテーブルの書き込みに失敗しました。"
						errorlog.Print(message)
						errorlog.Print(err)
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "Taskテーブルの書き込みに失敗しました。"
					errorlog.Print(message)
					errorlog.Print(err)
					errArray = append(errArray, message)
				}
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		if file.Filename == "types.csv" {
			types, errArray := dp.UpsertType(f);
			err := DB.Save(types).Error; if err != nil {
				message := "Typeテーブルの書き込みに失敗しました。"
				errorlog.Print(message)
				errorlog.Print(err)
				errArray = append(errArray, message)
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		if file.Filename == "users.csv" {
			users, errArray := dp.UpsertUser(f);
			bulk := []dp.User{}
			for _, record := range users {
				bulk = append(bulk, record)
				if len(bulk) > 500 {
					err := DB.Save(bulk).Error; if err != nil {
						message := "Userテーブルの書き込みに失敗しました。"
						errorlog.Print(message)
						errorlog.Print(err)
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "Userテーブルの書き込みに失敗しました。"
					errorlog.Print(message)
					errorlog.Print(err)
					errArray = append(errArray, message)
				}
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		errorlog.Print(file.Filename + "は除外されました")
		message := []string{file.Filename + "は除外されました"}
		errorArray = append(errorArray, message)
	}
	errJson, _ := json.Marshal(errorArray)
	w.WriteHeader(http.StatusOK)
	w.Write(errJson)
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

