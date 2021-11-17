package models

import "time"

type Category struct {
	Id          int    `json:"id,string"`
	Name        string `json:"Name" gorm:"unique"`
	Description string `json:"description"`
	CreatedBy   int
	User        User `gorm:"foreignKey:CreatedBy" json:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
