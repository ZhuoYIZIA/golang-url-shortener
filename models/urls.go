package models

import "time"

type Url struct {
	Id          string    `json:"id"`
	OriginalUrl string    `json:"original_url"`
	CreateAt    time.Time `json:"create_at"`
}
