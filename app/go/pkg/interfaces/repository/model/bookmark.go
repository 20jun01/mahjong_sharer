package model

type BookMark struct {
	UserId string `gorm:"type:char(36)"`
	PostId string `gorm:"type:char(36)"`
}
