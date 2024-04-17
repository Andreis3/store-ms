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
	DB ipostgres.IInstructionDB
	*pgconn.PgError
}

func NewGroupRepository() *GroupRepository {
	return &GroupRepository{}
}

func (r *GroupRepository) InsertGroup(data GroupModel) (string, *util.ValidationError) {
	query := `INSERT INTO groups (id, group_name, code, status, created_at, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`
	rows, _ := r.DB.Query(context.Background(), query,
		data.ID,
		data.GroupName,
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
