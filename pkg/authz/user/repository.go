package authz

import (
	"context"
	"fmt"

	"github.com/forge4flow/forge4flow-core/pkg/database"
	"github.com/forge4flow/forge4flow-core/pkg/service"
	"github.com/pkg/errors"
)

type UserRepository interface {
	Create(ctx context.Context, user Model) (int64, error)
	GetById(ctx context.Context, id int64) (Model, error)
	GetByUserId(ctx context.Context, userId string) (Model, error)
	List(ctx context.Context, listParams service.ListParams) ([]Model, error)
	UpdateByUserId(ctx context.Context, userId string, user Model) error
	DeleteByUserId(ctx context.Context, userId string) error
}

func NewRepository(db database.Database) (UserRepository, error) {
	switch db.Type() {
	case database.TypeMySQL:
		mysql, ok := db.(*database.MySQL)
		if !ok {
			return nil, errors.New(fmt.Sprintf("invalid %s database config", database.TypeMySQL))
		}

		return NewMySQLRepository(mysql), nil
	case database.TypePostgres:
		postgres, ok := db.(*database.Postgres)
		if !ok {
			return nil, errors.New(fmt.Sprintf("invalid %s database config", database.TypePostgres))
		}

		return NewPostgresRepository(postgres), nil
	case database.TypeSQLite:
		sqlite, ok := db.(*database.SQLite)
		if !ok {
			return nil, errors.New(fmt.Sprintf("invalid %s database config", database.TypeSQLite))
		}

		return NewSQLiteRepository(sqlite), nil
	default:
		return nil, errors.New(fmt.Sprintf("unsupported database type %s specified", db.Type()))
	}
}
