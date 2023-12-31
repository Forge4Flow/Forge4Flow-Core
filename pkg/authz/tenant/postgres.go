package tenant

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"time"

	"github.com/forge4flow/forge4flow-core/pkg/database"
	"github.com/forge4flow/forge4flow-core/pkg/service"
	"github.com/pkg/errors"
)

type PostgresRepository struct {
	database.SQLRepository
}

func NewPostgresRepository(db *database.Postgres) *PostgresRepository {
	return &PostgresRepository{
		database.NewSQLRepository(db),
	}
}

func (repo PostgresRepository) Create(ctx context.Context, model Model) (int64, error) {
	var newTenantId int64
	err := repo.DB(ctx).GetContext(
		ctx,
		&newTenantId,
		`
			INSERT INTO tenant (
				tenant_id,
				object_id,
				name
			) VALUES (?, ?, ?)
			ON CONFLICT (tenant_id) DO UPDATE SET
				object_id = ?,
				name = ?,
				created_at = CURRENT_TIMESTAMP(6),
				deleted_at = NULL
			RETURNING id
		`,
		model.GetTenantId(),
		model.GetObjectId(),
		model.GetName(),
		model.GetObjectId(),
		model.GetName(),
	)

	if err != nil {
		return -1, errors.Wrap(err, "error creating tenant")
	}

	return newTenantId, nil
}

func (repo PostgresRepository) GetById(ctx context.Context, id int64) (Model, error) {
	var tenant Tenant
	err := repo.DB(ctx).GetContext(
		ctx,
		&tenant,
		`
			SELECT id, object_id, tenant_id, name, created_at, updated_at, deleted_at
			FROM tenant
			WHERE
				id = ? AND
				deleted_at IS NULL
		`,
		id,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, service.NewRecordNotFoundError("Tenant", id)
		default:
			return nil, errors.Wrapf(err, "error getting tenant %d", id)
		}
	}

	return &tenant, nil
}

func (repo PostgresRepository) GetByTenantId(ctx context.Context, tenantId string) (Model, error) {
	var tenant Tenant
	err := repo.DB(ctx).GetContext(
		ctx,
		&tenant,
		`
			SELECT id, object_id, tenant_id, name, created_at, updated_at, deleted_at
			FROM tenant
			WHERE
				tenant_id = ? AND
				deleted_at IS NULL
		`,
		tenantId,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, service.NewRecordNotFoundError("Tenant", tenantId)
		default:
			return nil, errors.Wrapf(err, "error getting tenant %s", tenantId)
		}
	}

	return &tenant, nil
}

func (repo PostgresRepository) List(ctx context.Context, listParams service.ListParams) ([]Model, error) {
	models := make([]Model, 0)
	tenants := make([]Tenant, 0)
	query := `
		SELECT id, object_id, tenant_id, name, created_at, updated_at, deleted_at
		FROM tenant
		WHERE
			deleted_at IS NULL

	`
	replacements := []interface{}{}

	if listParams.Query != "" {
		searchTermReplacement := fmt.Sprintf("%%%s%%", listParams.Query)
		query = fmt.Sprintf("%s AND (tenant_id LIKE ? OR name LIKE ?)", query)
		replacements = append(replacements, searchTermReplacement, searchTermReplacement)
	}

	sortBy := regexp.MustCompile("([A-Z])").ReplaceAllString(listParams.SortBy, `_$1`)
	if listParams.AfterId != "" {
		if listParams.AfterValue != nil {
			if listParams.SortOrder == service.SortOrderAsc {
				query = fmt.Sprintf("%s AND (%s > ? OR (tenant_id > ? AND %s = ?))", query, sortBy, sortBy)
				replacements = append(replacements,
					listParams.AfterValue,
					listParams.AfterId,
					listParams.AfterValue,
				)
			} else {
				query = fmt.Sprintf("%s AND (%s < ? OR (tenant_id < ? AND %s = ?))", query, sortBy, sortBy)
				replacements = append(replacements,
					listParams.AfterValue,
					listParams.AfterId,
					listParams.AfterValue,
				)
			}
		} else {
			if listParams.SortOrder == service.SortOrderAsc {
				query = fmt.Sprintf("%s AND tenant_id > ?", query)
				replacements = append(replacements, listParams.AfterId)
			} else {
				query = fmt.Sprintf("%s AND tenant_id < ?", query)
				replacements = append(replacements, listParams.AfterId)
			}
		}
	}

	if listParams.BeforeId != "" {
		if listParams.BeforeValue != nil {
			if listParams.SortOrder == service.SortOrderAsc {
				query = fmt.Sprintf("%s AND (%s < ? OR (tenant_id < ? AND %s = ?))", query, sortBy, sortBy)
				replacements = append(replacements,
					listParams.BeforeValue,
					listParams.BeforeId,
					listParams.BeforeValue,
				)
			} else {
				query = fmt.Sprintf("%s AND (%s > ? OR (tenant_id > ? AND %s = ?))", query, sortBy, sortBy)
				replacements = append(replacements,
					listParams.BeforeValue,
					listParams.BeforeId,
					listParams.BeforeValue,
				)
			}
		} else {
			if listParams.SortOrder == service.SortOrderAsc {
				query = fmt.Sprintf("%s AND tenant_id < ?", query)
				replacements = append(replacements, listParams.AfterId)
			} else {
				query = fmt.Sprintf("%s AND tenant_id > ?", query)
				replacements = append(replacements, listParams.AfterId)
			}
		}
	}

	nullSortClause := "NULLS LAST"
	if listParams.SortOrder == service.SortOrderAsc {
		nullSortClause = "NULLS FIRST"
	}

	if listParams.SortBy != "tenantId" {
		query = fmt.Sprintf("%s ORDER BY %s %s %s, tenant_id %s LIMIT ?", query, sortBy, listParams.SortOrder, nullSortClause, listParams.SortOrder)
		replacements = append(replacements, listParams.Limit)
	} else {
		query = fmt.Sprintf("%s ORDER BY tenant_id %s %s LIMIT ?", query, listParams.SortOrder, nullSortClause)
		replacements = append(replacements, listParams.Limit)
	}

	err := repo.DB(ctx).SelectContext(
		ctx,
		&tenants,
		query,
		replacements...,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return models, nil
		default:
			return models, errors.Wrap(err, "error listing tenants")
		}
	}

	for i := range tenants {
		models = append(models, &tenants[i])
	}

	return models, nil
}

func (repo PostgresRepository) UpdateByTenantId(ctx context.Context, tenantId string, model Model) error {
	_, err := repo.DB(ctx).ExecContext(
		ctx,
		`
			UPDATE tenant
			SET
				name = ?
			WHERE
				tenant_id = ? AND
				deleted_at IS NULL
		`,
		model.GetName(),
		tenantId,
	)
	if err != nil {
		return errors.Wrapf(err, "error updating tenant %s", tenantId)
	}

	return nil
}

func (repo PostgresRepository) DeleteByTenantId(ctx context.Context, tenantId string) error {
	_, err := repo.DB(ctx).ExecContext(
		ctx,
		`
			UPDATE tenant
			SET
				deleted_at = ?
			WHERE
				tenant_id = ? AND
				deleted_at IS NULL
		`,
		time.Now().UTC(),
		tenantId,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return service.NewRecordNotFoundError("Tenant", tenantId)
		default:
			return errors.Wrapf(err, "error deleting tenant %s", tenantId)
		}
	}

	return nil
}
