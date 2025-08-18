package domain

import "time"

// URL represents a shortened URL entity.
// This is a shared domain model used by multiple services.
type URL struct {
	ShortKey  string    `json:"short_key"`
	LongURL   string    `json:"long_url"`
	CreatedAt time.Time `json:"created_at"`
	Redirects int       `json:"redirects"`
}
