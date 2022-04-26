package models

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


type Comment struct {
	ID   	  int		`gorm:"AUTO_INCREMENT"json:"id"`
	Content	  string	`json:"content"`
	TaskID 	  int		`json:"task_id"`
	UserID 	  int		`json:"user_id"`
	ParentID  *int		`json:"parent_id"`
	User 	  User		`gorm:"foreignkey:UserID;"json:"user"`
	Replies	[]Comment 	`gorm:"foreignKey:ParentID"json:"replies"`
	// Replies	[]Comment 	`gorm:"many2many:comment_replies"json:"replies"`
	CreatedAt time.Time `gorm:"autoCreateTime;"json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"json:"updated_at"`
}


func NewComment(r *http.Request) (Comment, error) {
	comment, _ := GetCommentJson(r)
	return comment, nil
}

func GetComments(id int) ([]Comment, error) {
	var comments []Comment
	tx := DB.Preload(clause.Associations)
	tx = RecursivePreload(tx)
	result := tx.Find(&comments, "task_id = ? AND parent_id = ?", id, 0); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return comments, result.Error
	}
	return comments, nil
}

func (c *Comment) Create() error {
	result := DB.Debug().Omit("User").Create(&c); if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *Comment)Update() error {
	result := DB.Debug().Omit("User").Session(&gorm.Session{FullSaveAssociations: true}).Save(&c); if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

func RecursivePreload(tx *gorm.DB) *gorm.DB {
	column := "Replies"
	tx.Preload(column + ".User")
	for i := 0; i < 100; i++{
		column += ".Replies"
		tx.Preload(column)
		tx.Preload(column + ".User")
	}
	return tx
}

func GetCommentJson(r *http.Request) (Comment, error) {
	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		message := "couldn't decode json"
		err := errors.New(message)
		return comment, err
	}
	return comment, nil
}