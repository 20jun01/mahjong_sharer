package model

import "time"

type Post struct {
	PostId    string    `gorm:"type:char(36);primaryKey;not null"`
	UserId    string    `gorm:"type:char(36);not null"`
	Title     string    `gorm:"type:varchar(255);not null"`
	Content   string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"precision:6"`
	UpdatedAt time.Time `gorm:"precision:6"`
}
