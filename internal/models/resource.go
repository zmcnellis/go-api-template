package models

import (
	"time"

	"github.com/lib/pq"
)

// Resource ...
type Resource struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Link        string         `json:"link"`
	Name        string         `json:"name"`
	Author      string         `json:"author"`
	Description string         `json:"description"`
	Tags        pq.StringArray `gorm:"type:varchar(64)[]" json:"tags"`
}
