package flow

import (
	"context"
	"database/sql"
	"fmt"
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
			INSERT INTO flowEvent (
				type
			) VALUES (?)
		`,
		model.GetType(),
	)
	if err != nil {
		return -1, errors.Wrap(err, "error creating event type")
	}

	newEventId, err := result.LastInsertId()
	if err != nil {
		return -1, errors.Wrap(err, "error creating event type")
	}

	return newEventId, nil
}

func (repo MySQLRepository) GetById(ctx context.Context, id int64) (Model, error) {
	var event Event
	err := repo.DB(ctx).GetContext(
		ctx,
		&event,
		`
			SELECT id, type, lastBlockHeight, createdAt, updatedAt, deletedAt
			FROM flowEvent
			WHERE
				id = ? AND
				deletedAt IS NULL
		`,
		id,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, service.NewRecordNotFoundError("Event", id)
		default:
			return nil, errors.Wrapf(err, "error getting Event %d", id)
		}
	}

	return &event, nil
}

func (repo MySQLRepository) GetByType(ctx context.Context, eventType string) (Model, error) {
	var eventObject Event
	err := repo.DB(ctx).GetContext(
		ctx,
		&eventObject,
		`
			SELECT id, type, lastBlockHeight, createdAt, updatedAt, deletedAt
			FROM flowEvent
			WHERE
				type = ? AND
				deletedAt IS NULL
		`,
		eventType,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, service.NewRecordNotFoundError("Event", eventType)
		default:
			return nil, errors.Wrapf(err, "error getting event %s", eventType)
		}
	}

	return &eventObject, nil
}

func (repo MySQLRepository) GetAllEvents(ctx context.Context) ([]Model, error) {
	var eventObjects []Event
	err := repo.DB(ctx).SelectContext(
		ctx,
		&eventObjects,
		`
			SELECT id, type, lastBlockHeight, createdAt, updatedAt, deletedAt
			FROM flowEvent
			WHERE
				deletedAt IS NULL
		`,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, service.NewRecordNotFoundError("Event", "None")
		default:
			return nil, errors.Wrapf(err, "error getting events")
		}
	}

	fmt.Println(eventObjects)

	var models []Model
	for _, event := range eventObjects {
		eventModel := event
		models = append(models, &eventModel)
	}

	return models, nil
}

func (repo MySQLRepository) DeleteById(ctx context.Context, id int64) error {
	_, err := repo.DB(ctx).ExecContext(
		ctx,
		`
			UPDATE flowEvent
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
			return service.NewRecordNotFoundError("FlowId", id)
		default:
			return errors.Wrapf(err, "error deleting event with ID %v", id)
		}
	}

	return nil
}

func (repo MySQLRepository) DeleteByType(ctx context.Context, eventType string) error {
	_, err := repo.DB(ctx).ExecContext(
		ctx,
		`
			UPDATE flowEvent
			SET deletedAt = ?
			WHERE
				type = ? AND
				deletedAt IS NULL
		`,
		time.Now().UTC(),
		eventType,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return service.NewRecordNotFoundError("Event", eventType)
		default:
			return errors.Wrapf(err, "error deleting event %s", eventType)
		}
	}

	return nil
}

func (repo MySQLRepository) UpdateLastBlockHeightByType(ctx context.Context, eventType string, lastBlockHeight uint64) error {
	_, err := repo.DB(ctx).ExecContext(
		ctx,
		`
			UPDATE flowEvent
			SET lastBlockHeight = ?
			WHERE
				type = ? AND
				deletedAt IS NULL
		`,
		lastBlockHeight,
		eventType,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return service.NewRecordNotFoundError("Event", eventType)
		default:
			return errors.Wrapf(err, "error deleting event %s", eventType)
		}
	}

	return nil
}
