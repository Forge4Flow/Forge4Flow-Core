package authn

import (
	"context"
	"fmt"

	"github.com/forge4flow/forge4flow-core/pkg/database"
	"github.com/pkg/errors"
)

type ApiKeyRepository interface {
	Create(ctx context.Context, key Model) (int64, error)
	GetById(ctx context.Context, id int64) (Model, error)
	GetByKey(ctx context.Context, key string) (Model, error)
	DeleteById(ctx context.Context, id int64) error
	DeleteByKey(ctx context.Context, key string) error
	// DeleteAllExpired(ctx context.Context, userId string) error
}

func NewRepository(db database.Database) (ApiKeyRepository, error) {
	switch db.Type() {
	case database.TypeMySQL:
		mysql, ok := db.(*database.MySQL)
		if !ok {
			return nil, errors.New(fmt.Sprintf("invalid %s database config", database.TypeMySQL))
		}

		return NewMySQLRepository(mysql), nil
	//TODO: Finish Repositories for PostgresSQL and SQLite
	// case database.TypePostgres:
	// 	postgres, ok := db.(*database.Postgres)
	// 	if !ok {
	// 		return nil, errors.New(fmt.Sprintf("invalid %s database config", database.TypePostgres))
	// 	}

	// 	return NewPostgresRepository(postgres), nil
	// case database.TypeSQLite:
	// 	sqlite, ok := db.(*database.SQLite)
	// 	if !ok {
	// 		return nil, errors.New(fmt.Sprintf("invalid %s database config", database.TypeSQLite))
	// 	}

	// 	return NewSQLiteRepository(sqlite), nil
	default:
		return nil, errors.New(fmt.Sprintf("unsupported database type %s specified", db.Type()))
	}
}
