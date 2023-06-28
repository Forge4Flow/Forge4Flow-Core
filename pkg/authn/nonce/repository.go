package authn

import (
	"context"
	"fmt"

	"github.com/auth4flow/auth4flow-core/pkg/database"
	"github.com/pkg/errors"
)

type NonceRepository interface {
	Create(ctx context.Context, nonce Model) (int64, error)
	GetById(ctx context.Context, id int64) (Model, error)
	GetByNonce(ctx context.Context, nonce string) (Model, error)
	DeleteById(ctx context.Context, id int64) error
	DeleteByNonce(ctx context.Context, nonce string) error
	// DeleteAllExpired(ctx context.Context, userId string) error
}

func NewRepository(db database.Database) (NonceRepository, error) {
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
