package domain

import "time"

type Search struct {
	ID			string		`json:"id"`
	Avatar 		string		`json:"avatar"`
	Name		string		`json:"name"`
	Categories	[]string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
