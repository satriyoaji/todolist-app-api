package domain

import "time"

type Attachment struct {
	Id        int
	TodoId    int
	File      string
	Caption   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
