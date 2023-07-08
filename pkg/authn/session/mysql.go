package authn

import (
	"context"
	"database/sql"

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
			INSERT INTO session (
				sessionId,
				userId,
				idleTimeout,
				expTime,
				userAgent,
				clientIp
			) VALUES (?, ?, ?, ?, ?, ?)
		`,
		model.GetSessionId(),
		model.GetUserId(),
		int64(model.GetIdleTimeout()),
		model.GetExpTime(),
		model.GetUserAgent(),
		model.GetClientIp(),
	)
	if err != nil {
		return -1, errors.Wrap(err, "error creating session")
	}

	newSessionId, err := result.LastInsertId()
	if err != nil {
		return -1, errors.Wrap(err, "error creating session")
	}

	return newSessionId, nil
}

func (repo MySQLRepository) GetById(ctx context.Context, id int64) (Model, error) {
	var nonce Session
	err := repo.DB(ctx).GetContext(
		ctx,
		&nonce,
		`
			SELECT id, sessionId, userId, lastActivity, idleTimeout, expTime, userAgent, clientIp, createdAt, updatedAt
			FROM session
			WHERE
				id = ? AND
				deletedAt IS NULL
		`,
		id,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, service.NewRecordNotFoundError("Session", id)
		default:
			return nil, errors.Wrapf(err, "error getting session %d", id)
		}
	}

	return &nonce, nil
}
