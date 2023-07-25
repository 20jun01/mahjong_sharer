package database_query

import (
	"context"
)

type UserQueryInterface interface {
	GetUser(ctx context.Context, id int32) (string, error)
}
