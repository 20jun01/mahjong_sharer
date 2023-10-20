package model

import "time"

type User struct {
	UserId    string    `gorm:"type:char(36);primaryKey;not null"`
	Name      string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"precision:6"`
	UpdatedAt time.Time `gorm:"precision:6"`
}
