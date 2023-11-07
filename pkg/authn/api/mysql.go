package authn

import (
	"context"
	"database/sql"
	"time"

	"github.com/forge4flow/forge4flow-core/pkg/database"
	"github.com/forge4flow/forge4flow-core/pkg/service"
	"github.com/pkg/errors"
)

type MySQLRepository struct {
	database.SQLRepository
}

func NewMySQLRepository(db *database.MySQL) *MySQLRepository {
	return &MySQLRepository{
		database.NewSQLRepository(db),
	}
}

func (repo MySQLRepository) Create(ctx context.Context, model Model) (int64, error) {
	result, err := repo.DB(ctx).ExecContext(
		ctx,
		`
			INSERT INTO keys (
				displayName,
				key,
				expDate
			) VALUES (?, ?)
		`,
		model.GetName(),
		model.GetKey(),
		model.GetExpDate(),
	)
	if err != nil {
		return -1, errors.Wrap(err, "error creating nonce")
	}

	newNonceId, err := result.LastInsertId()
	if err != nil {
		return -1, errors.Wrap(err, "error creating user")
	}

	return newNonceId, nil
}

func (repo MySQLRepository) GetById(ctx context.Context, id int64) (Model, error) {
	var key ApiKey
	err := repo.DB(ctx).GetContext(
		ctx,
		&key,
		`
			SELECT id, displayName, nonce, expDate, createdAt, updatedAt, deletedAt
			FROM keys
			WHERE
				id = ? AND
				deletedAt IS NULL
		`,
		id,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, service.NewRecordNotFoundError("Nonce", id)
		default:
			return nil, errors.Wrapf(err, "error getting nonce %d", id)
		}
	}

	return &key, nil
}

func (repo MySQLRepository) GetByKey(ctx context.Context, key string) (Model, error) {
	var keyObject ApiKey
	err := repo.DB(ctx).GetContext(
		ctx,
		&keyObject,
		`
			SELECT id, displayName, nonce, expDate, createdAt, updatedAt, deletedAt
			FROM keys
			WHERE
				nonce = ? AND
				deletedAt IS NULL
		`,
		key,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, service.NewRecordNotFoundError("User", key)
		default:
			return nil, errors.Wrapf(err, "error getting user %s", key)
		}
	}

	return &keyObject, nil
}

func (repo MySQLRepository) DeleteById(ctx context.Context, id int64) error {
	_, err := repo.DB(ctx).ExecContext(
		ctx,
		`
			UPDATE keys
			SET deletedAt = ?
			WHERE
				id = ? AND
				deletedAt IS NULL
		`,
		time.Now().UTC(),
		id,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return service.NewRecordNotFoundError("NonceID", id)
		default:
			return errors.Wrapf(err, "error deleting nonce with ID %s", id)
		}
	}

	return nil
}

func (repo MySQLRepository) DeleteByKey(ctx context.Context, key string) error {
	_, err := repo.DB(ctx).ExecContext(
		ctx,
		`
			UPDATE keys
			SET deletedAt = ?
			WHERE
				key = ? AND
				deletedAt IS NULL
		`,
		time.Now().UTC(),
		key,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return service.NewRecordNotFoundError("Nonce", key)
		default:
			return errors.Wrapf(err, "error deleting nonce %s", key)
		}
	}

	return nil
}

//TODO: Create DeleAllExpired function to delete all expired Key Objects
// func (repo MySQLRepository) DeleteAllExpired(ctx context.Context, userId string) error {
// 	_, err := repo.DB(ctx).ExecContext(
// 		ctx,
// 		`
// 			UPDATE user
// 			SET deletedAt = ?
// 			WHERE
// 				userId = ? AND
// 				deletedAt IS NULL
// 		`,
// 		time.Now().UTC(),
// 		userId,
// 	)
// 	if err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			return service.NewRecordNotFoundError("User", userId)
// 		default:
// 			return errors.Wrapf(err, "error deleting user %s", userId)
// 		}
// 	}

// 	return nil
// }
