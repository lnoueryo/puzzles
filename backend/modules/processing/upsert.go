package processing

import (
	"backend/models"
	"encoding/csv"
	"io"
	"strconv"
	"time"
)

var timestamp = "2006-01-02 15:04:05"

func UpsertActivity(file io.Reader) ([]models.Activity, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var activities []models.Activity
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "Activityテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return activities, errorArray
	}
	for i, row := range content[1:] {
		if len(row) != 6 {
			errMessage := "Activity: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ID, err := strconv.Atoi(row[0]); if err != nil {
			errMessage := "Activity: " + strconv.Itoa(i + 2) + "列目のidが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		UserID, err := strconv.Atoi(row[1]); if err != nil {
			errMessage := "Activity: " + strconv.Itoa(i + 2) + "列目のuser_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ProjectID, _ := strconv.Atoi(row[2]); if err != nil {
			errMessage := "Activity: " + strconv.Itoa(i + 2) + "列目のproject_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ContentID, _ := strconv.Atoi(row[3]); if err != nil {
			errMessage := "Activity: " + strconv.Itoa(i + 2) + "列目のcontent_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		CreatedAt, _ := time.Parse(timestamp, row[4])
		UpdatedAt, _ := time.Parse(timestamp, row[5])
		activity := models.Activity{
			ID: ID,
			UserID: UserID,
			ProjectID: ProjectID,
			ContentID: ContentID,
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		activities = append(activities, activity)
	}
	return activities, errorArray
}

func UpsertActivityContent(file io.Reader) ([]models.ActivityContent, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var activity_contents []models.ActivityContent
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "ActivityContentテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return activity_contents, errorArray
	}
	for i, row := range content[1:] {
		if len(row) != 2 {
			errMessage := "ActivityContent: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ID, err := strconv.Atoi(row[0]); if err != nil {
			errMessage := "ActivityContent: " + strconv.Itoa(i + 2) + "列目のidが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		activity_content := models.ActivityContent{
			ID: ID,
			Content: row[1],
		}
		activity_contents = append(activity_contents, activity_content)
	}
	return activity_contents, errorArray
}

func UpsertAuthority(file io.Reader) ([]models.Authority, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var authorities []models.Authority
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "Authorityテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return authorities, errorArray
	}
	for i, row := range content[1:] {
		if len(row) != 2 {
			errMessage := "Authority: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ID, err := strconv.Atoi(row[0]); if err != nil {
			errMessage := "Authority: " + strconv.Itoa(i + 2) + "列目のidが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		authority := models.Authority{
			ID: ID,
			Name: row[1],
		}
		authorities = append(authorities, authority)
	}
	return authorities, errorArray
}

func UpsertComment(file io.Reader) ([]models.Comment, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var comments []models.Comment
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "Commentテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return comments, errorArray
	}
	for i, row := range content[1:] {
		if len(row) != 7 {
			errMessage := "Comment: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ID, err := strconv.Atoi(row[0]); if err != nil {
			errMessage := "Comment: " + strconv.Itoa(i + 2) + "列目のidが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		TaskID, err := strconv.Atoi(row[2]); if err != nil {
			errMessage := "Comment: " + strconv.Itoa(i + 2) + "列目のtask_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		UserID, err := strconv.Atoi(row[3]); if err != nil {
			errMessage := "Comment: " + strconv.Itoa(i + 2) + "列目のuser_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ParentID, err := strconv.Atoi(row[4]); if err != nil {
			errMessage := "Comment: " + strconv.Itoa(i + 2) + "列目のparent_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		CreatedAt, _ := time.Parse(timestamp, row[5])
		UpdatedAt, _ := time.Parse(timestamp, row[6])
		comment := models.Comment{
			ID: ID,
			Content: row[1],
			TaskID: TaskID,
			UserID: UserID,
			ParentID: &ParentID,
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		comments = append(comments, comment)
	}
	return comments, errorArray
}

func UpsertField(file io.Reader) ([]models.Field, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var fields []models.Field
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "Fieldテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return fields, errorArray
	}
	for i, row := range content[1:] {
		if len(row) != 5 {
			errMessage := "Field: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ID, err := strconv.Atoi(row[0]); if err != nil {
			errMessage := "Field: " + strconv.Itoa(i + 2) + "列目のidが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ProjectID, err := strconv.Atoi(row[1]); if err != nil {
			errMessage := "Field: " + strconv.Itoa(i + 2) + "列目のproject_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		CreatedAt, _ := time.Parse(timestamp, row[3])
		UpdatedAt, _ := time.Parse(timestamp, row[4])
		field := models.Field{
			ID: ID,
			ProjectID: ProjectID,
			Name: row[2],
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		fields = append(fields, field)
	}
	return fields, errorArray
}

// fieldとmilestoneのカラムの順番が違う

func UpsertMilestone(file io.Reader) ([]models.Milestone, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var milestones []models.Milestone
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "Milestoneテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return milestones, errorArray
	}
	for i, row := range content[1:] {
		if len(row) != 5 {
			errMessage := "Milestone: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ID, err := strconv.Atoi(row[0]); if err != nil {
			errMessage := "Milestone: " + strconv.Itoa(i + 2) + "列目のidが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ProjectID, err := strconv.Atoi(row[2]); if err != nil {
			errMessage := "Milestone: " + strconv.Itoa(i + 2) + "列目のproject_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		CreatedAt, _ := time.Parse(timestamp, row[3])
		UpdatedAt, _ := time.Parse(timestamp, row[4])
		milestone := models.Milestone{
			ID: ID,
			Name: row[1],
			ProjectID: ProjectID,
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		milestones = append(milestones, milestone)
	}
	return milestones, errorArray
}

func UpsertVersion(file io.Reader) ([]models.Version, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var versions []models.Version
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "Versionテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return versions, errorArray
	}
	for i, row := range content[1:] {
		if len(row) != 5 {
			errMessage := "Milestone: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ID, err := strconv.Atoi(row[0]); if err != nil {
			errMessage := "Milestone: " + strconv.Itoa(i + 2) + "列目のidが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ProjectID, err := strconv.Atoi(row[2]); if err != nil {
			errMessage := "Milestone: " + strconv.Itoa(i + 2) + "列目のproject_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		CreatedAt, _ := time.Parse(timestamp, row[3])
		UpdatedAt, _ := time.Parse(timestamp, row[4])
		version := models.Version{
			ID: ID,
			Name: row[1],
			ProjectID: ProjectID,
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		versions = append(versions, version)
	}
	return versions, errorArray
}

func UpsertOrganizationAuthority(file io.Reader) ([]models.OrganizationAuthority, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var organizationAuthorities []models.OrganizationAuthority
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "OrganizationAuthorityテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return organizationAuthorities, errorArray
	}
	for i, row := range content[1:] {
		if len(row) != 7 {
			errMessage := "OrganizationAuthority: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ID, err := strconv.Atoi(row[0]); if err != nil {
			errMessage := "OrganizationAuthority: " + strconv.Itoa(i + 2) + "列目のidが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		UserID, err := strconv.Atoi(row[1]); if err != nil {
			errMessage := "OrganizationAuthority: " + strconv.Itoa(i + 2) + "列目のuser_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		AuthorityID, err := strconv.Atoi(row[3]); if err != nil {
			errMessage := "OrganizationAuthority: " + strconv.Itoa(i + 2) + "列目のauthority_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		Active, err := strconv.ParseBool(row[4]); if err != nil {
			errMessage := "OrganizationAuthority: " + strconv.Itoa(i + 2) + "列目のactiveが論理型ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		CreatedAt, _ := time.Parse(timestamp, row[5])
		UpdatedAt, _ := time.Parse(timestamp, row[6])
		organizationAuthority := models.OrganizationAuthority{
			ID: ID,
			UserID: UserID,
			OrganizationID: row[2],
			AuthorityID: AuthorityID,
			Active: Active,
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		organizationAuthorities = append(organizationAuthorities, organizationAuthority)
	}
	return organizationAuthorities, errorArray
}

func UpsertOrganization(file io.Reader) ([]models.Organization, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var organizations []models.Organization
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "Organizationテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return organizations, errorArray
	}
	for i, row := range content[1:] {
		if len(row) != 12 {
			errMessage := "Organization: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		CreatedAt, _ := time.Parse(timestamp, row[10])
		UpdatedAt, _ := time.Parse(timestamp, row[11])
		organization := models.Organization{
			ID: row[0],
			Name: row[1],
			Address: row[2],
			Number: row[3],
			Founded: row[4],
			Image: row[5],
			Description: row[6],
			Plan: row[7],
			CreditCard: row[8],
			Expiry: row[9],
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		organizations = append(organizations, organization)
	}
	return organizations, errorArray
}

func UpsertPriority(file io.Reader) ([]models.Priority, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var priorities []models.Priority
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "Priorityテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return priorities, errorArray
	}
	for i, row := range content[1:] {
		if len(row) != 2 {
			errMessage := "Priority: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ID, err := strconv.Atoi(row[0]); if err != nil {
			errMessage := "Priority: " + strconv.Itoa(i + 2) + "列目のidが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		priority := models.Priority{
			ID: ID,
			Name: row[1],
		}
		priorities = append(priorities, priority)
	}
	return priorities, errorArray
}

func UpsertProjectAuthority(file io.Reader) ([]models.ProjectAuthority, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var projectAuthorities []models.ProjectAuthority
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "ProjectAuthorityテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return projectAuthorities, errorArray
	}
	for i, row := range content[1:] {
		if len(row) != 7 {
			errMessage := "ProjectAuthority: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ID, err := strconv.Atoi(row[0]); if err != nil {
			errMessage := "ProjectAuthority: " + strconv.Itoa(i + 2) + "列目のidが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ProjectID, err := strconv.Atoi(row[1]); if err != nil {
			errMessage := "ProjectAuthority: " + strconv.Itoa(i + 2) + "列目のproject_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		UserID, err := strconv.Atoi(row[2]); if err != nil {
			errMessage := "ProjectAuthority: " + strconv.Itoa(i + 2) + "列目のuser_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		AuthorityID, err := strconv.Atoi(row[3]); if err != nil {
			errMessage := "ProjectAuthority: " + strconv.Itoa(i + 2) + "列目のauthority_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		Active, err := strconv.ParseBool(row[4]); if err != nil {
			errMessage := "ProjectAuthority: " + strconv.Itoa(i + 2) + "列目のactiveが論理型ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		CreatedAt, _ := time.Parse(timestamp, row[5])
		UpdatedAt, _ := time.Parse(timestamp, row[6])
		priority := models.ProjectAuthority{
			ID: ID,
			ProjectID: ProjectID,
			UserID: UserID,
			AuthorityID: AuthorityID,
			Active: Active,
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		projectAuthorities = append(projectAuthorities, priority)
	}
	return projectAuthorities, errorArray
}

func UpsertProject(file io.Reader) ([]Project, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var projects []Project
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "Projectテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return projects, errorArray
	}
	for i, row := range content[1:] {
		if len(row) != 7 {
			errMessage := "Project: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ID, err := strconv.Atoi(row[0]); if err != nil {
			errMessage := "Project: " + strconv.Itoa(i + 2) + "列目のidが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		CreatedAt, _ := time.Parse(timestamp, row[5])
		UpdatedAt, _ := time.Parse(timestamp, row[6])
		project := Project{
			ID: ID,
			OrganizationID: row[1],
			Name: row[2],
			Description: row[3],
			Image: row[4],
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		projects = append(projects, project)
	}
	return projects, errorArray
}

func UpsertStatus(file io.Reader) ([]models.Status, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var statuses []models.Status
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "Statusテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return statuses, errorArray
	}
	for i, row := range content[1:] {
		if len(row) != 2 {
			errMessage := "Status: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ID, err := strconv.Atoi(row[0]); if err != nil {
			errMessage := "Status: " + strconv.Itoa(i + 2) + "列目のidが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		status := models.Status{
			ID: ID,
			Name: row[1],
		}
		statuses = append(statuses, status)
	}
	return statuses, errorArray
}

func UpsertTask(file io.Reader) ([]models.Task, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var tasks []models.Task
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "Taskテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return tasks, errorArray
	}

	for i, row := range content[1:] {
		if len(row) != 20 {
			errMessage := "Task: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ID, err := strconv.Atoi(row[0]); if err != nil {
			errMessage := "Task: " + strconv.Itoa(i + 2) + "列目のidが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		AssigneeID, err := strconv.Atoi(row[1]); if err != nil {
			errMessage := "Task: " + strconv.Itoa(i + 2) + "列目のassignee_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		AssignerID, err := strconv.Atoi(row[2]); if err != nil {
			errMessage := "Task: " + strconv.Itoa(i + 2) + "列目のassigner_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		StatusID, err := strconv.Atoi(row[3]); if err != nil {
			errMessage := "Task: " + strconv.Itoa(i + 2) + "列目のstatus_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		FieldID, err := strconv.Atoi(row[4]); if err != nil {
			if row[4] != "NULL" {
				errMessage := "Task: " + strconv.Itoa(i + 2) + "列目のfield_idが数字ではありません"
				errorArray = append(errorArray, errMessage)
				continue
			}
		}
		MilestoneID, err := strconv.Atoi(row[5]); if err != nil {
			if row[5] != "NULL" {
				errMessage := "Task: " + strconv.Itoa(i + 2) + "列目のmilestone_idが数字ではありません"
				errorArray = append(errorArray, errMessage)
				continue
			}
		}
		VersionID, err := strconv.Atoi(row[6]); if err != nil {
			if row[6] != "NULL" {
				errMessage := "Task: " + strconv.Itoa(i + 2) + "列目のmilestone_idが数字ではありません"
				errorArray = append(errorArray, errMessage)
				continue
			}
		}
		PriorityID, err := strconv.Atoi(row[7]); if err != nil {
			errMessage := "Task: " + strconv.Itoa(i + 2) + "列目のpriority_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		TypeID, err := strconv.Atoi(row[8]); if err != nil {
			errMessage := "Task: " + strconv.Itoa(i + 2) + "列目のtype_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ProjectID, err := strconv.Atoi(row[9]); if err != nil {
			errMessage := "Task: " + strconv.Itoa(i + 2) + "列目のproject_idが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ParentID, err := strconv.Atoi(row[10]); if err != nil {
			if row[10] != "NULL" {
				errMessage := "Task: " + strconv.Itoa(i + 2) + "列目のparent_idが数字ではありません"
				errorArray = append(errorArray, errMessage)
				continue
			}
		}
		EstimatedTime, err := strconv.ParseFloat(row[14], 32); if err != nil {
			if row[14] != "NULL" {
				errMessage := "Task: " + strconv.Itoa(i + 2) + "列目のestimated_timeが数字ではありません"
				errorArray = append(errorArray, errMessage)
				continue
			}
		}
		ActualTime, err := strconv.ParseFloat(row[15], 32); if err != nil {
			if row[15] != "NULL" {
				errMessage := "Task: " + strconv.Itoa(i + 2) + "列目のactual_timeが数字ではありません"
				errorArray = append(errorArray, errMessage)
				continue
			}
		}
		StartTime, _ := time.Parse(timestamp, row[16])
		Deadline, _ := time.Parse(timestamp, row[17])
		CreatedAt, _ := time.Parse(timestamp, row[18])
		UpdatedAt, _ := time.Parse(timestamp, row[19])
		task := models.Task{
			ID: ID,
			AssigneeID: AssigneeID,
			AssignerID: AssignerID,
			StatusID: StatusID,
			FieldID: &FieldID,
			MilestoneID: &MilestoneID,
			PriorityID: PriorityID,
			VersionID: &VersionID,
			TypeID: TypeID,
			ProjectID: ProjectID,
			ParentID: ParentID,
			Key: row[11],
			Title: row[12],
			Detail: row[13],
			EstimatedTime: float32(EstimatedTime),
			ActualTime: float32(ActualTime),
			StartTime: StartTime,
			Deadline: Deadline,
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		tasks = append(tasks, task)
	}
	return tasks, errorArray
}

func UpsertType(file io.Reader) ([]models.Type, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var types []models.Type
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "Typeテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return types, errorArray
	}
	for i, row := range content[1:] {
		if len(row) != 2 {
			errMessage := "Type: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ID, err := strconv.Atoi(row[0]); if err != nil {
			errMessage := "Type: " + strconv.Itoa(i + 2) + "列目のidが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		_type := models.Type{
			ID: ID,
			Name: row[1],
		}
		types = append(types, _type)
	}
	return types, errorArray
}

func UpsertUser(file io.Reader) ([]User, []string) {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var users []User
	var errorArray []string
	if len(content) <= 1 {
		errMessage := "Userテーブルのcsvが空です"
		errorArray = append(errorArray, errMessage)
		return users, errorArray
	}
	for i, row := range content[1:] {
		if len(row) != 11 {
			errMessage := "User: " + strconv.Itoa(i + 2) + "列目のレコードに足りない値があります"
			errorArray = append(errorArray, errMessage)
			continue
		}
		ID, err := strconv.Atoi(row[0]); if err != nil {
			errMessage := "User: " + strconv.Itoa(i + 2) + "列目のidが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		Age, err := strconv.Atoi(row[2]); if err != nil {
			errMessage := "User: " + strconv.Itoa(i + 2) + "列目のageが数字ではありません"
			errorArray = append(errorArray, errMessage)
			continue
		}
		CreatedAt, _ := time.Parse(timestamp, row[9])
		UpdatedAt, _ := time.Parse(timestamp, row[10])
		user := User{
			ID: ID,
			Name: row[1],
			Age: Age,
			Sex: row[3],
			Email: row[4],
			Address: row[5],
			Password: row[6],
			Image: row[7],
			Description: row[8],
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		users = append(users, user)
	}
	return users, errorArray
}

type Project struct {
	ID				int				`gorm:"AUTO_INCREMENT"json:"id"`
    OrganizationID	string			`gorm:"<-:create;"json:"organization_id"`
	Name			string			`json:"name"`
	Description		string			`json:"description"`
	Image			string			`json:"image"`
	CreatedAt		time.Time		`gorm:"autoCreateTime;"json:"created_at"`
	UpdatedAt		time.Time		`gorm:"autoUpdateTime;"json:"updated_at"`
}