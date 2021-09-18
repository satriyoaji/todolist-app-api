package domain

import "time"

type Todo struct {
	Id   int
	UserId int
	Title string
	CreatedAt time.Time
	UpdatedAt time.Time
}
