package domain

import "time"

type User struct {
	Id             int
	Fullname       string
	Email          string
	Password       string
	ForgotPassword string
	RoleId         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
