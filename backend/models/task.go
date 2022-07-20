package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


type Task struct {
	ID				int			`gorm:"AUTO_INCREMENT"json:"id"`
	AssigneeID		int			`gorm:"<-;not null"json:"assignee_id"`
	AssignerID		int			`gorm:"<-;not null"json:"assigner_id"`
	StatusID		int			`gorm:"<-;not null"json:"status_id"`
	FieldID			*int		`gorm:"<-;"json:"field_id"`
	MilestoneID		*int		`gorm:"<-;"json:"milestone_id"`
	VersionID		*int		`gorm:"<-;"json:"version_id"`
	PriorityID		int			`gorm:"<-;not null"json:"priority_id"`
	TypeID			int			`gorm:"<-;not null"json:"type_id"`
	ProjectID		int			`gorm:"<-;not null"json:"project_id"`
	ParentID		int			`gorm:"<-;not null"json:"parent_id"`
	Key				string		`json:"key"`
	Title			string		`gorm:"<-;not null"json:"title" sql:"CHARACTER SET utf8 COLLATE utf8_unicode_ci"`
	Detail			string		`gorm:"<-;"json:"detail"`
	EstimatedTime	float32		`json:"estimated_time"`
	ActualTime		float32		`json:"actual_time"`
	StartTime		time.Time	`gorm:"default:null;"json:"start_time"`
	Deadline		time.Time	`gorm:"default:null;"json:"deadline"`
	// Status        Status    `gorm:"embedded;embeddedPrefix:status_"`
	Status			Status		`gorm:"references:ID;"json:"status"`
	Field			Field    	`gorm:"references:ID"json:"field"`
	Milestone		Milestone	`gorm:"references:ID"json:"milestone"`
	Version			Version		`gorm:"references:ID"json:"version"`
	Type			Type    	`gorm:"references:ID"json:"type"`
	Priority		Priority	`gorm:"references:ID"json:"priority"`
	Assignee		User		`gorm:"references:ID;foreignKey:AssigneeID"json:"assignee"`
	Assigner		User		`gorm:"references:ID;foreignKey:AssignerID"json:"assigner"`
	Comments		[]Comment	`json:"comments"`
	CreatedAt		time.Time	`gorm:"<-:create;autoCreateTime;"json:"created_at"`
	UpdatedAt		time.Time	`gorm:"autoUpdateTime;"json:"updated_at"`
}


func NewTask(r *http.Request) (Task, error) {
	task, _ := GetTaskJson(r)
	return task, nil
}

func GetTasksByProjectID(DB *gorm.DB, projectID int) ([]Task, error) {
	var t []Task
	result := DB.Preload("Comments", "id = ?", 0).Preload(clause.Associations).Find(&t, "project_id = ?", projectID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return t, result.Error
	}
	return t, nil
}

func (t *Task)GetTask(DB *gorm.DB) error {
	tx := DB.Preload("Comments", "parent_id = ?", 0).Preload(clause.Associations)
	tx = RecursivePreload(tx)
	result := tx.Find(&t, t.ID)
	fmt.Println(t)
	// result := DB.Preload(RecursivePreload()).Preload("Comments", "parent_id = ?", 1).Preload(clause.Associations).First(&t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *Task)Create(DB *gorm.DB) error {
	result := DB.Debug().Omit("Assignee", "Assigner", "Field", "Status", "Milestone", "Version", "Type", "Priority").Create(&t); if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *Task)Update(DB *gorm.DB) error {
	fmt.Println(t.EstimatedTime)
	result := DB.Debug().Omit("Assignee", "Assigner", "Field", "Status", "Milestone", "Version", "Type", "Priority").Save(&t).Clauses(clause.Returning{}); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (t *Task)CountProjectTask(DB *gorm.DB) (Project, error) {
	var project Project
	result := DB.Preload("Tasks").Find(&project, t.ProjectID); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return project, result.Error
	}
	return project, nil
}

func GetTaskJson(r *http.Request) (Task, error) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		message := "couldn't decode json"
		err := errors.New(message)
		return task, err
	}
	return task, nil
}