package model

import "time"

type Link struct {
	ShortCode string
	URL       string

	Clicks int
	CreatedAt time.Time
}