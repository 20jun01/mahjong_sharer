package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/mahjong_sharer/pkg/interfaces/repository/model"
)

func Migrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		v1(),
	}
}

func AllTables() []interface{} {
	return []interface{}{
		&model.User{},
		&model.BookMark{},
		&model.Post{},
	}
}
