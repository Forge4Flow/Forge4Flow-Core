package authn

import (
	"context"
	"database/sql"
	"time"

	"github.com/auth4flow/auth4flow-core/pkg/database"
	"github.com/auth4flow/auth4flow-core/pkg/service"
	"github.com/pkg/errors"
)

type MySQLRepository struct {
	database.SQLRepository
}

func NewMySQLRepository(db *database.MySQL) MySQLRepository {
	return MySQLRepository{
		database.NewSQLRepository(db),
	}
}

func (repo MySQLRepository) Create(ctx context.Context, model Model) (int64, error) {
	result, err := repo.DB(ctx).ExecContext(
		ctx,
		`
			INSERT INTO nonce (
				nonce,
				expDate,
			) VALUES (?, ?)
		`,
		model.GetNonce(),
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
	var nonce Nonce
	err := repo.DB(ctx).GetContext(
		ctx,
		&nonce,
		`
			SELECT id, nonce, expDate, createdAt, updatedAt, deletedAt
			FROM nonce
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

	return &nonce, nil
}

func (repo MySQLRepository) GetByNonce(ctx context.Context, nonce string) (Model, error) {
	var nonceObject Nonce
	err := repo.DB(ctx).GetContext(
		ctx,
		&nonceObject,
		`
			SELECT id, nonce, expDate, createdAt, updatedAt, deletedAt
			FROM nonce
			WHERE
				nonce = ? AND
				deletedAt IS NULL
		`,
		nonce,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, service.NewRecordNotFoundError("User", nonce)
		default:
			return nil, errors.Wrapf(err, "error getting user %s", nonce)
		}
	}

	return &nonceObject, nil
}

func (repo MySQLRepository) DeleteById(ctx context.Context, id int64) error {
	_, err := repo.DB(ctx).ExecContext(
		ctx,
		`
			UPDATE nonce
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

func (repo MySQLRepository) DeleteByNonce(ctx context.Context, nonce string) error {
	_, err := repo.DB(ctx).ExecContext(
		ctx,
		`
			UPDATE nonce
			SET deletedAt = ?
			WHERE
				nonce = ? AND
				deletedAt IS NULL
		`,
		time.Now().UTC(),
		nonce,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return service.NewRecordNotFoundError("Nonce", nonce)
		default:
			return errors.Wrapf(err, "error deleting nonce %s", nonce)
		}
	}

	return nil
}

//TODO: Create DeleAllExpired function to delete all expired Nonce Objects
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
