package processing

import (
	"archive/zip"
	"backend/models"
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	// "io"
	// "os"
	"strconv"
	"time"
)

type CSVBuffer struct {
	Name string
	Byte *bytes.Buffer
}

func DLActivity(activities []models.Activity) *bytes.Buffer {
	rows := [][]string{{"id", "user_id", "project_id", "content_id", "created_at", "updated_at"}}
	for _, activity := range activities {
		row := []string{}
		ID := strconv.Itoa(activity.ID)
		UserID := strconv.Itoa(activity.UserID)
		ProjectID := strconv.Itoa(activity.ProjectID)
		ContentID := strconv.Itoa(activity.ContentID)
		CreatedAt := activity.CreatedAt.String()
		UpdatedAt := activity.UpdatedAt.String()
		row = append(row, ID, UserID, ProjectID, ContentID, CreatedAt, UpdatedAt)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	return b
}

func DLActivityContent(activity_contents []models.ActivityContent) *bytes.Buffer {
	rows := [][]string{{"id", "content"}}
	for _, activity_content := range activity_contents {
		row := []string{}
		ID := strconv.Itoa(activity_content.ID)
		Content := activity_content.Content
		row = append(row, ID, Content)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	return b
}

func DLAuthority(authorities []models.Authority) *bytes.Buffer {
	rows := [][]string{{"id", "name"}}
	for _, authority := range authorities {
		row := []string{}
		ID := strconv.Itoa(authority.ID)
		Name := authority.Name
		row = append(row, ID, Name)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	return b
}

func DLComment(comments []models.Comment) *bytes.Buffer {
	rows := [][]string{{"id", "content", "task_id", "user_id", "parent_id", "created_at", "updated_at"}}
	for _, comment := range comments {
		row := []string{}
		ID := strconv.Itoa(comment.ID)
		Content := comment.Content
		TaskID := strconv.Itoa(comment.TaskID)
		UserID := strconv.Itoa(comment.UserID)
		ParentID := strconv.Itoa(*comment.ParentID)
		CreatedAt := comment.CreatedAt.String()
		UpdatedAt := comment.UpdatedAt.String()
		row = append(row, ID, Content, TaskID, UserID, ParentID, CreatedAt, UpdatedAt)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	return b
}

func DLField(fields []models.Field) *bytes.Buffer {
	rows := [][]string{{"id", "project_id", "name", "created_at", "updated_at"}}
	for _, field := range fields {
		row := []string{}
		ID := strconv.Itoa(field.ID)
		ProjectID := strconv.Itoa(field.ProjectID)
		Name := field.Name
		CreatedAt := field.CreatedAt.String()
		UpdatedAt := field.UpdatedAt.String()
		row = append(row, ID, ProjectID, Name, CreatedAt, UpdatedAt)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	return b
}

func DLMilestone(milestones []models.Milestone) *bytes.Buffer {
	rows := [][]string{{"id", "name", "project_id", "created_at", "updated_at"}}
	for _, milestone := range milestones {
		row := []string{}
		ID := strconv.Itoa(milestone.ID)
		Name := milestone.Name
		ProjectID := strconv.Itoa(milestone.ProjectID)
		CreatedAt := milestone.CreatedAt.String()
		UpdatedAt := milestone.UpdatedAt.String()
		row = append(row, ID, Name, ProjectID, CreatedAt, UpdatedAt)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	return b
}

func DLOrganizationAuthority(organizationAuthorities []models.OrganizationAuthority) *bytes.Buffer {
	rows := [][]string{{"id", "user_id", "organization_id", "authority_id", "active", "created_at", "updated_at"}}
	for _, organizationAuthority := range organizationAuthorities {
		row := []string{}
		ID := strconv.Itoa(organizationAuthority.ID)
		UserID := strconv.Itoa(organizationAuthority.UserID)
		OrganizationID := organizationAuthority.OrganizationID
		AuthorityID := strconv.Itoa(organizationAuthority.AuthorityID)
		Active := strconv.FormatBool(organizationAuthority.Active)
		CreatedAt := organizationAuthority.CreatedAt.String()
		UpdatedAt := organizationAuthority.UpdatedAt.String()
		row = append(row, ID, UserID, OrganizationID, AuthorityID, Active, CreatedAt, UpdatedAt)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	return b
}

func DLOrganization(organizations []models.Organization) *bytes.Buffer {
	rows := [][]string{{"id", "name", "address", "number", "founded", "image", "description", "plan", "credit_card", "expiry", "created_at", "updated_at"}}
	for _, organization := range organizations {
		row := []string{}
		ID := organization.ID
		Name := organization.Name
		Address := organization.Address
		Number := organization.Number
		Founded := organization.Founded
		Image := organization.Image
		Description := organization.Description
		Plan := organization.Plan
		CreditCard := organization.CreditCard
		Expiry := organization.Expiry
		CreatedAt := organization.CreatedAt.String()
		UpdatedAt := organization.UpdatedAt.String()
		row = append(row, ID, Name, Address, Number, Founded, Image, Description, Plan, CreditCard, Expiry, CreatedAt, UpdatedAt)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	return b
}

func DLPriority(priorities []models.Priority) *bytes.Buffer {
	rows := [][]string{{"id", "name"}}
	for _, priority := range priorities {
		row := []string{}
		ID := strconv.Itoa(priority.ID)
		Name := priority.Name
		row = append(row, ID, Name)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	return b
}

func DLProjectAuthority(project_authorities []models.ProjectAuthority) *bytes.Buffer {
	rows := [][]string{{"id", "project_id", "user_id", "authority_id", "active", "created_at", "updated_at"}}
	for _, project_authority := range project_authorities {
		row := []string{}
		ID := strconv.Itoa(project_authority.ID)
		ProjectID := strconv.Itoa(project_authority.ProjectID)
		UserID := strconv.Itoa(project_authority.UserID)
		AuthorityID := strconv.Itoa(project_authority.AuthorityID)
		Active := strconv.FormatBool(project_authority.Active)
		CreatedAt := project_authority.CreatedAt.String()
		UpdatedAt := project_authority.UpdatedAt.String()
		row = append(row, ID, ProjectID, UserID, AuthorityID, Active, CreatedAt, UpdatedAt)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	return b
}

func DLProject(projects []models.Project) *bytes.Buffer {
	rows := [][]string{{"id", "organization_id", "name", "description", "image", "created_at", "updated_at"}}
	for _, project := range projects {
		row := []string{}
		ID := strconv.Itoa(project.ID)
		OrganizationID := project.OrganizationID
		Name := project.Name
		Description := project.Description
		Image := project.Image
		CreatedAt := project.CreatedAt.String()
		UpdatedAt := project.UpdatedAt.String()
		row = append(row, ID, OrganizationID, Name, Description, Image, CreatedAt, UpdatedAt)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	return b
}

func DLStatus(statuses []models.Status) *bytes.Buffer {
	rows := [][]string{{"id", "name"}}
	for _, status := range statuses {
		row := []string{}
		ID := strconv.Itoa(status.ID)
		Name := status.Name
		row = append(row, ID, Name)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	return b
}

func DLTask(tasks []models.Task) *bytes.Buffer {
	rows := [][]string{{"id", "assignee_id", "assigner_id", "status_id", "field_id", "milestone_id", "priority_id", "type_id", "project_id", "parent_id", "key", "title", "detail", "estimated_time", "actual_time", "start_time", "deadline", "created_at", "updated_at"}}
	for _, task := range tasks {
		row := []string{}
		ID := strconv.Itoa(task.ID)
		AssigneeID := strconv.Itoa(task.AssigneeID)
		AssignerID := strconv.Itoa(task.AssignerID)
		StatusID := strconv.Itoa(task.StatusID)
		FieldID := strconv.Itoa(*task.FieldID)
		MilestoneID := strconv.Itoa(*task.MilestoneID)
		PriorityID := strconv.Itoa(task.PriorityID)
		TypeID := strconv.Itoa(task.TypeID)
		ProjectID := strconv.Itoa(task.ProjectID)
		ParentID := strconv.Itoa(task.ParentID)
		Key := task.Key
		Title := task.Title
		Detail := task.Detail
		EstimatedTime := fmt.Sprintf("%f", task.EstimatedTime)
		ActualTime := fmt.Sprintf("%f", task.ActualTime)
		StartTime := task.StartTime.String()
		Deadline := task.Deadline.String()
		CreatedAt := task.CreatedAt.String()
		UpdatedAt := task.UpdatedAt.String()
		row = append(row, ID, AssigneeID, AssignerID, StatusID, FieldID, MilestoneID, PriorityID, TypeID, ProjectID, ParentID, Key, Title, Detail, EstimatedTime, ActualTime, StartTime, Deadline, CreatedAt, UpdatedAt)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	return b
}

func DLType(types []models.Type) *bytes.Buffer {
	rows := [][]string{{"id", "name"}}
	for _, typeTable := range types {
		row := []string{}
		ID := strconv.Itoa(typeTable.ID)
		Name := typeTable.Name
		row = append(row, ID, Name)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	return b
}

func DLUser(users []User) *bytes.Buffer {
	rows := [][]string{{"id", "name", "age", "sex", "email", "address", "password", "image", "description", "created_at", "updated_at"}}
	for _, user := range users {
		row := []string{}
		ID := strconv.Itoa(user.ID)
		Name := user.Name
		Age := strconv.Itoa(user.Age)
		Sex := user.Sex
		Email := user.Email
		Address := user.Address
		Password := user.Password
		Image := user.Image
		Description := user.Description
		CreatedAt := user.CreatedAt.String()
		UpdatedAt := user.UpdatedAt.String()
		row = append(row, ID, Name, Age, Sex, Email, Address, Password, Image, Description, CreatedAt, UpdatedAt)
		rows = append(rows, row)
	}
    b := new(bytes.Buffer)
    writer := csv.NewWriter(b)
	writer.WriteAll(rows)
	return b
}

func ZipCSVByteArray(bufArray *[]CSVBuffer) *bytes.Buffer {

    buf := new(bytes.Buffer)
    w := zip.NewWriter(buf)
	for _, buf := range *bufArray {
		body, _ := ioutil.ReadAll(buf.Byte)
		fh := &zip.FileHeader{
			Name:     buf.Name + ".csv",
			Modified: time.Now(),
			Method:   8,
		}
		f, err := w.CreateHeader(fh)
		if err != nil {
			fmt.Print(err)
		}
		if _, err := f.Write(body); err != nil {
			fmt.Print(err)
		}
	}
	w.Close()
	return buf
}

type User struct {
	ID						int						`gorm:"AUTO_INCREMENT"json:"id"`
	Name					string					`json:"name"`
	Age						int						`json:"age"`
	Sex						string					`json:"sex"`
	Email					string					`json:"email"`
	Address					string					`json:"address"`
	Password				string					`json:"password"`
	Image					string					`json:"image"`
	Description				string					`json:"description"`
	CreatedAt				time.Time				`gorm:"->:false;<-:create;autoCreateTime;"json:"-"`
	UpdatedAt				time.Time				`gorm:"autoUpdateTime;"json:"updated_at"`
}