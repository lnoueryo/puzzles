package services

import (
	"backend/models"
	dp "backend/modules/processing"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CSVRequest struct {
	Request []string `json:"request"`
}

func ImportCSV(r *http.Request) ([][]string, error) {
	var errorArray [][]string
	err := r.ParseMultipartForm(20000000);if err != nil {
		return errorArray, err
	}

	formdata := r.MultipartForm // ok, no problem so far, read the Form data
	//get the *fileheaders
	files := formdata.File["files"] // grab the filenames
	for _, file := range files {
		f, err := file.Open()
		defer f.Close()
		if err != nil {
			return errorArray, err
		}
		if file.Filename == "activities.csv" {
			activities, errArray := dp.UpsertActivity(f)
			err := DB.Save(&activities).Error; if err != nil {
				message := "Activityテーブルの書き込みに失敗しました。"
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
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "Fieldテーブルの書き込みに失敗しました。"
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
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "Milestoneテーブルの書き込みに失敗しました。"
					errArray = append(errArray, message)
				}
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		if file.Filename == "versions.csv" {
			versions, errArray := dp.UpsertVersion(f);
			bulk := []models.Version{}
			for _, record := range versions {
				bulk = append(bulk, record)
				if len(bulk) > 500 {
					err := DB.Save(bulk).Error; if err != nil {
						message := "Versionテーブルの書き込みに失敗しました。"
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "Versionテーブルの書き込みに失敗しました。"
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
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "OrganizationAuthorityテーブルの書き込みに失敗しました。"
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
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "Organizationテーブルの書き込みに失敗しました。"
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
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "ProjectAuthorityテーブルの書き込みに失敗しました。"
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
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "Projectテーブルの書き込みに失敗しました。"
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
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "Taskテーブルの書き込みに失敗しました。"
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
						errArray = append(errArray, message)
					}
					bulk = nil
				}
			}
			if len(bulk) != 0 {
				err := DB.Save(bulk).Error; if err != nil {
					message := "Userテーブルの書き込みに失敗しました。"
					errArray = append(errArray, message)
				}
			}
			if len(errArray) != 0 {
				errorArray = append(errorArray, errArray)
			}
			continue
		}
		message := []string{file.Filename + "は除外されました"}
		errorArray = append(errorArray, message)
	}
	return errorArray, nil
}

func ExportCSV(r *http.Request) (*bytes.Buffer, error) {
	var b *bytes.Buffer
	byteArray, _ := ioutil.ReadAll(r.Body)
	var csvRequest CSVRequest
	err := json.Unmarshal(byteArray, &csvRequest); if err != nil{
		return b, err
	}

	var csvBufferArray []dp.CSVBuffer
	for _, request := range csvRequest.Request {
		if request == "activities" {
			var activities []models.Activity
			DB.Find(&activities)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLActivity(activities),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "activity_contents" {
			var activity_contents []models.ActivityContent
			DB.Find(&activity_contents)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLActivityContent(activity_contents),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "authorities" {
			var authorities []models.Authority
			DB.Find(&authorities)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLAuthority(authorities),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "comments" {
			var comments []models.Comment
			DB.Find(&comments)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLComment(comments),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "fields" {
			var fields []models.Field
			DB.Find(&fields)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLField(fields),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "milestones" {
			var milestones []models.Milestone
			DB.Find(&milestones)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLMilestone(milestones),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "versions" {
			var versions []models.Version
			DB.Find(&versions)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLVersion(versions),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "organization_authorities" {
			var organization_authorities []models.OrganizationAuthority
			DB.Find(&organization_authorities)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLOrganizationAuthority(organization_authorities),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "organizations" {
			var organizations []models.Organization
			DB.Find(&organizations)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLOrganization(organizations),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "priorities" {
			var priorities []models.Priority
			DB.Find(&priorities)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLPriority(priorities),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "project_authorities" {
			var project_authorities []models.ProjectAuthority
			DB.Find(&project_authorities)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLProjectAuthority(project_authorities),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "projects" {
			var projects []models.Project
			DB.Find(&projects)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLProject(projects),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "statuses" {
			var statuses []models.Status
			DB.Find(&statuses)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLStatus(statuses),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "tasks" {
			var tasks []models.Task
			DB.Find(&tasks)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLTask(tasks),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "types" {
			var types []models.Type
			DB.Find(&types)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLType(types),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
		if request == "users" {
			var users []dp.User
			DB.Find(&users)
			csvBuffer := dp.CSVBuffer{
				Name: request,
				Byte: dp.DLUser(users),
			}
			csvBufferArray = append(csvBufferArray, csvBuffer)
		}
	}
	b = dp.ZipCSVByteArray(&csvBufferArray)

	return b, nil
}