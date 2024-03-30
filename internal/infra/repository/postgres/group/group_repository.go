package repo_group

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"

	"github.com/andreis3/stores-ms/internal/util"
)

type GroupRepository struct {
	db *pgxpool.Pool
	*pgconn.PgError
}

func NewGroupRepository(db *pgxpool.Pool) *GroupRepository {
	return &GroupRepository{
		db: db,
	}
}

func (r *GroupRepository) InsertGroup(data GroupModel) (string, *util.ValidationError) {
	query := `INSERT INTO groups (id, group_name, code, status, created_at, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`
	rows, _ := r.db.Query(context.Background(), query,
		data.ID,
		data.GroupName,
		data.Code,
		data.Status,
		data.CreatedAt,
		data.UpdatedAt)
	defer rows.Close()
	group, err := pgx.CollectOneRow[GroupModel](rows, pgx.RowToStructByPos[GroupModel])
	if errors.As(err, &r.PgError) {
		return "", &util.ValidationError{
			Code:        fmt.Sprintf("PIDB-%s", r.Code),
			Status:      http.StatusInternalServerError,
			LogError:    []string{fmt.Sprintf("%s, %s", r.Message, r.Detail)},
			ClientError: []string{"Internal Server Error"},
		}
	}
	return group.ID, nil
}
