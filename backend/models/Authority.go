package models


type Authority struct {
	ID		int		`json:"id"`
	// 権限の種類
	Name	string	`json:"name"`
}