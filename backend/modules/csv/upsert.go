package csv

import (
	"backend/models"
	"encoding/csv"
	"io"
	"strconv"
	"time"
)

var timestamp = "2006-01-02 15:04:05"

func UpsertActivity(file io.Reader) []models.Activity {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var activities []models.Activity
	for _, row := range content[1:] {
		ID, _ := strconv.Atoi(row[0])
		UserID, _ := strconv.Atoi(row[1])
		ProjectID, _ := strconv.Atoi(row[2])
		ContentID, _ := strconv.Atoi(row[3])
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
	return activities
}

func UpsertActivityContent(file io.Reader) []models.ActivityContent {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var activity_contents []models.ActivityContent
	for _, row := range content[1:] {
		ID, _ := strconv.Atoi(row[0])
		activity_content := models.ActivityContent{
			ID: ID,
			Content: row[1],
		}
		activity_contents = append(activity_contents, activity_content)
	}
	return activity_contents
}

func UpsertAuthority(file io.Reader) []models.Authority {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var authorities []models.Authority
	for _, row := range content[1:] {
		ID, _ := strconv.Atoi(row[0])
		authority := models.Authority{
			ID: ID,
			Name: row[1],
		}
		authorities = append(authorities, authority)
	}
	return authorities
}

func UpsertComment(file io.Reader) []models.Comment {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var comments []models.Comment
	for _, row := range content[1:] {
		ID, _ := strconv.Atoi(row[0])
		TaskID, _ := strconv.Atoi(row[2])
		UserID, _ := strconv.Atoi(row[3])
		ParentID, _ := strconv.Atoi(row[4])
		CreatedAt, _ := time.Parse(timestamp, row[4])
		UpdatedAt, _ := time.Parse(timestamp, row[5])
		comment := models.Comment{
			ID: ID,
			Content: row[1],
			TaskID: TaskID,
			UserID: UserID,
			ParentID: ParentID,
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		comments = append(comments, comment)
	}
	return comments
}

func UpsertField(file io.Reader) []models.Field {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var fields []models.Field
	for _, row := range content[1:] {
		ID, _ := strconv.Atoi(row[0])
		ProjectID, _ := strconv.Atoi(row[1])
		CreatedAt, _ := time.Parse(timestamp, row[4])
		UpdatedAt, _ := time.Parse(timestamp, row[5])
		field := models.Field{
			ID: ID,
			ProjectID: ProjectID,
			Name: row[2],
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		fields = append(fields, field)
	}
	return fields
}

// fieldとmilestoneのカラムの順番が違う

func UpsertMilestone(file io.Reader) []models.Milestone {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var milestones []models.Milestone
	for _, row := range content[1:] {
		ID, _ := strconv.Atoi(row[0])
		ProjectID, _ := strconv.Atoi(row[2])
		CreatedAt, _ := time.Parse(timestamp, row[4])
		UpdatedAt, _ := time.Parse(timestamp, row[5])
		milestone := models.Milestone{
			ID: ID,
			Name: row[1],
			ProjectID: ProjectID,
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		milestones = append(milestones, milestone)
	}
	return milestones
}

func UpsertOrganizationAuthority(file io.Reader) []models.OrganizationAuthority {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var organizationAuthorities []models.OrganizationAuthority
	for _, row := range content[1:] {
		ID, _ := strconv.Atoi(row[0])
		UserID, _ := strconv.Atoi(row[1])
		AuthorityID, _ := strconv.Atoi(row[3])
		Active, _ := strconv.ParseBool(row[4])
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
	return organizationAuthorities
}

func UpsertOrganization(file io.Reader) []models.Organization {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var organizations []models.Organization
	for _, row := range content[1:] {
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
	return organizations
}

func UpsertPriority(file io.Reader) []models.Priority {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var priorities []models.Priority
	for _, row := range content[1:] {
		ID, _ := strconv.Atoi(row[0])
		priority := models.Priority{
			ID: ID,
			Name: row[1],
		}
		priorities = append(priorities, priority)
	}
	return priorities
}

func UpsertProjectAuthority(file io.Reader) []models.ProjectAuthority {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var projectAuthorities []models.ProjectAuthority
	for _, row := range content[1:] {
		ID, _ := strconv.Atoi(row[0])
		ProjectID, _ := strconv.Atoi(row[1])
		UserID, _ := strconv.Atoi(row[2])
		AuthorityID, _ := strconv.Atoi(row[3])
		Active, _ := strconv.ParseBool(row[4])
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
	return projectAuthorities
}

func UpsertProject(file io.Reader) []models.Project {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var projects []models.Project
	for _, row := range content[1:] {
		ID, _ := strconv.Atoi(row[0])
		CreatedAt, _ := time.Parse(timestamp, row[6])
		UpdatedAt, _ := time.Parse(timestamp, row[7])
		project := models.Project{
			ID: ID,
			OrganizationID: row[1],
			Name: row[2],
			Description: row[3],
			Image: row[4],
			ImageData: row[5],
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		projects = append(projects, project)
	}
	return projects
}

func UpsertStatus(file io.Reader) []models.Status {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var statuses []models.Status
	for _, row := range content[1:] {
		ID, _ := strconv.Atoi(row[0])
		status := models.Status{
			ID: ID,
			Name: row[1],
		}
		statuses = append(statuses, status)
	}
	return statuses
}

func UpsertTask(file io.Reader) []models.Task {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var tasks []models.Task
	for _, row := range content[1:] {
		ID, _ := strconv.Atoi(row[0])
		AssigneeID, _ := strconv.Atoi(row[1])
		AssignerID, _ := strconv.Atoi(row[2])
		StatusID, _ := strconv.Atoi(row[3])
		FieldID, _ := strconv.Atoi(row[4])
		MilestoneID, _ := strconv.Atoi(row[5])
		PriorityID, _ := strconv.Atoi(row[6])
		TypeID, _ := strconv.Atoi(row[7])
		ProjectID, _ := strconv.Atoi(row[8])
		ParentID, _ := strconv.Atoi(row[9])
		EstimatedTime, _ := strconv.ParseFloat(row[13], 32)
		ActualTime, _ := strconv.ParseFloat(row[14], 32)
		StartTime, _ := time.Parse(timestamp, row[15])
		Deadline, _ := time.Parse(timestamp, row[16])
		CreatedAt, _ := time.Parse(timestamp, row[17])
		UpdatedAt, _ := time.Parse(timestamp, row[18])
		task := models.Task{
			ID: ID,
			AssigneeID: AssigneeID,
			AssignerID: AssignerID,
			StatusID: StatusID,
			FieldID: &FieldID,
			MilestoneID: &MilestoneID,
			PriorityID: PriorityID,
			TypeID: TypeID,
			ProjectID: ProjectID,
			ParentID: ParentID,
			EstimatedTime: float32(EstimatedTime),
			ActualTime: float32(ActualTime),
			StartTime: StartTime,
			Deadline: Deadline,
			CreatedAt: CreatedAt,
			UpdatedAt: UpdatedAt,
		}
		tasks = append(tasks, task)
	}
	return tasks
}

func UpsertType(file io.Reader) []models.Type {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var types []models.Type
	for _, row := range content[1:] {
		ID, _ := strconv.Atoi(row[0])
		_type := models.Type{
			ID: ID,
			Name: row[1],
		}
		types = append(types, _type)
	}
	return types
}

func UpsertUser(file io.Reader) []models.User {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()
	var users []models.User
	for _, row := range content[1:] {
		ID, _ := strconv.Atoi(row[0])
		Age, _ := strconv.Atoi(row[2])
		CreatedAt, _ := time.Parse(timestamp, row[10])
		UpdatedAt, _ := time.Parse(timestamp, row[11])
		user := models.User{
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
	return users
}