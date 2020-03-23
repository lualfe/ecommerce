package models

import "time"

// User model
type User struct {
	ID        int        `json:"id"`
	Email     string     `gorm:"type:varchar(50)" json:"email" form:"email"`
	Password  string     `gorm:"type:varchar(120)" json:"-" form:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at,omitempty"`
}
