package domain

import "time"

type Attachment struct {
	Id        int
	TodoId    int
	Location  string
	Caption   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
