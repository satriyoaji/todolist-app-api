package todo

import "time"

type TodoResponse struct {
	Id   int    `json:"id"`
	UserId int `json:"user_id"`
	Title string `json:"title"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
