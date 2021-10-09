package attachment

import "time"

type AttachmentResponse struct {
	Id        int       `json:"id"`
	TodoId    int       `json:"todo_id"`
	File      string    `json:"file"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
