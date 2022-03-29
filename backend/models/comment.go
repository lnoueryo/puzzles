package models

import "time"


type Comment struct {
	ID   	  int		`gorm:"AUTO_INCREMENT"json:"id"`
	Content	  string	`json:"name"`
	TaskID 	  int		`json:"-"`
	UserID 	  int		`json:"-"`
	// ParentID  int		`gorm:"primaryKey;foreignKey:ParentID"`
	User 	  User		`gorm:"references:ID"json:"user"`
	ParentID  bool		`json:"parent_id"`
	Replies []Comment 	`gorm:"many2many:comment_replies"`
	// Replies   []Comment `gorm:"many2many:comment_replies"`
	CreatedAt time.Time `gorm:"autoCreateTime;"json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;"json:"updated_at"`
}