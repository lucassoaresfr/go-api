package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Salary   int64
	Link     string
}

type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Role      string    `json:"deleted_at,omitempty"`
	Company   string    `json:"role"`
	Location  string    `json:"company"`
	Remote    bool      `json:"remote"`
	Salary    int64     `json:"salary"`
}
