package repo_group

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"

	ipostgres "github.com/andreis3/stores-ms/internal/infra/adapters/database/postgres/interfaces"
	"github.com/andreis3/stores-ms/internal/util"
)

type GroupRepository struct {
	postgres ipostgres.IPostgres
	*pgconn.PgError
}

func NewGroupRepository(pool ipostgres.IPostgres) *GroupRepository {
	return &GroupRepository{
		postgres: pool,
	}
}

func (r *GroupRepository) InsertGroup(data GroupModel) (string, *util.ValidationError) {
	query := `INSERT INTO groups (id, name, code, status, created_at, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`
	rows, _ := r.postgres.Query(context.Background(), query,
		data.ID,
		data.Name,
		data.Code,
		data.Status,
		data.CreatedAt,
		data.UpdatedAt)
	defer rows.Close()
	group, err := pgx.CollectOneRow[GroupModel](rows, pgx.RowToStructByName[GroupModel])
	if errors.As(err, &r.PgError) {
		return "", &util.ValidationError{
			Code:        fmt.Sprintf("PIDB-%s", r.Code),
			Status:      http.StatusInternalServerError,
			LogError:    []string{fmt.Sprintf("%s, %s", r.Message, r.Detail)},
			ClientError: []string{"Internal Server Error"},
		}
	}
	return *group.ID, nil
}
func (r *GroupRepository) SelectOneGroupByNameAndCode(groupName, code string) (*GroupModel, *util.ValidationError) {
	query := `SELECT * FROM groups WHERE name = $1 AND code = $2`
	rows, _ := r.postgres.Query(context.Background(), query, groupName, code)
	defer rows.Close()
	group, err := pgx.CollectOneRow[GroupModel](rows, pgx.RowToStructByName[GroupModel])
	if errors.As(err, &r.PgError) {
		return nil, &util.ValidationError{
			Code:        fmt.Sprintf("PIDB-%s", r.Code),
			Status:      http.StatusInternalServerError,
			LogError:    []string{fmt.Sprintf("%s, %s", r.Message, r.Detail)},
			ClientError: []string{"Internal Server Error"},
		}
	}
	return &group, nil
}
