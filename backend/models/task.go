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
	ID            int       `gorm:"AUTO_INCREMENT"json:"id"`
	AssigneeID    int       `gorm:"<-;not null"json:"assignee_id"`
	AssignerID    int       `gorm:"<-;not null"json:"assigner_id"`
	StatusID      int       `gorm:"<-;not null"json:"status_id"`
	FieldID       *int      `gorm:"<-;"json:"field_id"`
	MilestoneID   *int      `gorm:"<-;"json:"milestone_id"`
	VersionID     *int      `gorm:"<-;"json:"version_id"`
	PriorityID    int       `gorm:"<-;not null"json:"priority_id"`
	TypeID        int       `gorm:"<-;not null"json:"type_id"`
	ProjectID     int       `gorm:"<-;not null"json:"project_id"`
	ParentID      int       `gorm:"<-;not null"json:"parent_id"`
	Key           string    `json:"key"`
	Title         string    `gorm:"<-;not null"json:"title" sql:"CHARACTER SET utf8 COLLATE utf8_unicode_ci"`
	Detail        string    `gorm:"<-;"json:"detail"`
	EstimatedTime float32   `json:"estimated_time"`
	ActualTime    float32   `json:"actual_time"`
	StartTime     time.Time `gorm:"default:null;"json:"start_time"`
	Deadline      time.Time `gorm:"default:null;"json:"deadline"`
	// Status        Status    `gorm:"embedded;embeddedPrefix:status_"`
	Status    Status    `gorm:"references:ID;"json:"status"`
	Field     Field     `gorm:"references:ID"json:"field"`
	Milestone Milestone `gorm:"references:ID"json:"milestone"`
	Version   Version   `gorm:"references:ID"json:"version"`
	Type      Type      `gorm:"references:ID"json:"type"`
	Priority  Priority  `gorm:"references:ID"json:"priority"`
	Assignee  User      `gorm:"references:ID;foreignKey:AssigneeID"json:"assignee"`
	Assigner  User      `gorm:"references:ID;foreignKey:AssignerID"json:"assigner"`
	Comments  []Comment `json:"comments"`
	CreatedAt time.Time `gorm:"<-:create;autoCreateTime;"json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"json:"updated_at"`
}

func NewTask(r *http.Request) (Task, error) {
	task, _ := GetTaskJson(r)
	return task, nil
}

func GetTasksByProjectID(DB *gorm.DB, projectID int) ([]uint8, error) {
	start := time.Now()
	var JSONs [][]uint8
	DB.Raw(query, projectID).Scan(&JSONs)
	duration := time.Since(start)
	fmt.Println(duration)

	return JSONs[0], nil
}

const query = `
SELECT
	JSON_OBJECT(
		'tasks',
		JSON_ARRAYAGG(
			JSON_OBJECT(
			'id',
			t.id,
			'key',
			t.key,
			'title',
			t.title,
			'detail',
			t.detail,
			'estimated_time',
			t.estimated_time,
			'actual_time',
			t.actual_time,
			'start_time',
			t.start_time,
			'deadline',
			t.deadline,
			'created_at',
			t.created_at,
			'updated_at',
			t.updated_at,
			'assignee',
			t.assignee,
			'assigner',
			t.assigner,
			'status',
			t.status,
			'field',
			t.field,
			'milestone',
			t.milestone,
			'version',
			t.version,
			'priority',
			t.priority,
			'type',
			t.type
			)
		),
        'count', COUNT(t.id)
    ) as task
FROM
	(
		SELECT
			t.id,
			t.key,
			t.title,
			t.detail,
			t.estimated_time,
			t.actual_time,
			t.start_time,
			t.deadline,
			t.created_at,
			t.updated_at,
			JSON_OBJECT(
				'id', assignee.id,
				'name', assignee.name,
				'image', assignee.image
			) AS assignee,
			JSON_OBJECT(
				'id', assigner.id,
				'name', assigner.name,
				'image', assigner.image
			) AS assigner,
			JSON_OBJECT(
				'id', s.id,
				'name', s.name
			) AS status,
			JSON_OBJECT(
				'id', f.id,
				'name', f.name
			) AS field,
			JSON_OBJECT(
				'id', m.id,
				'name', m.name
			) AS milestone,
			JSON_OBJECT(
				'id', v.id,
				'name', v.name
			) AS version,
			JSON_OBJECT(
				'id', p.id,
				'name', p.name
			) AS priority,
			JSON_OBJECT(
				'id', tp.id,
				'name', tp.name
			) AS type
		FROM
			puzzles.tasks t
		LEFT JOIN puzzles.users assignee ON assignee.id = t.assignee_id
		LEFT JOIN puzzles.users assigner ON assigner.id = t.assigner_id
		LEFT JOIN puzzles.statuses s ON s.id = t.status_id
		LEFT JOIN puzzles.fields f ON f.id = t.field_id
		LEFT JOIN puzzles.milestones m ON m.id = t.milestone_id
		LEFT JOIN puzzles.versions v ON v.id = t.version_id
		LEFT JOIN puzzles.priorities p ON p.id = t.priority_id
		LEFT JOIN puzzles.types tp ON tp.id = t.type_id
		WHERE
			t.project_id = ?
		ORDER BY
			t.id
	) t
`

func (t *Task) GetTask(DB *gorm.DB) error {
	tx := DB.Preload("Comments", "parent_id = ?", 0).Preload(clause.Associations)
	tx = RecursivePreload(tx)
	result := tx.Find(&t, t.ID)
	// result := DB.Preload(RecursivePreload()).Preload("Comments", "parent_id = ?", 1).Preload(clause.Associations).First(&t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *Task) Create(DB *gorm.DB) error {
	result := DB.Omit("Assignee", "Assigner", "Field", "Status", "Milestone", "Version", "Type", "Priority").Create(&t)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *Task) Update(DB *gorm.DB) error {
	fmt.Print((t))
	result := DB.Omit("Assignee", "Assigner", "Field", "Status", "Milestone", "Version", "Type", "Priority").Save(&t).Clauses(clause.Returning{})
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func (t *Task) CountProjectTask(DB *gorm.DB) (Project, error) {
	var project Project
	result := DB.Preload("Tasks").Find(&project, t.ProjectID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
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
