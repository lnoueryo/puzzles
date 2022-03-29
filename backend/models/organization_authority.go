package models

import "time"



type OrganizationAuthority struct {
	ID				int			`gorm:"AUTO_INCREMENT"json:"id"`
    OrganizationID	string		`gorm:"not null"json:"-"`
    UserID			int			`gorm:"not null"json:"-"`
    AuthID			int			`json:"auth_id"`
	Active			bool		`json:"active"`
	Type			Authority	`gorm:"foreignkey:AuthID;migrate"json:"type"`
	CreatedAt		time.Time	`gorm:"<-:false;autoCreateTime;"json:"-"`
	UpdatedAt		time.Time	`gorm:"<-;autoUpdateTime;"json:"-"`
}

