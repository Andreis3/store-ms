package repo_group

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"

	entity_group "github.com/andreis3/stores-ms/internal/domain/entity/group"
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

func (r *GroupRepository) InsertGroup(data entity_group.Group) (*GroupModel, *util.ValidationError) {
	model := MapperGroupModel(data)
	query := `INSERT INTO groups (id, name, code, status, created_at, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`
	rows, _ := r.DB.Query(context.Background(), query,
		model.ID,
		model.Name,
		model.Code,
		model.Status,
		model.CreatedAt,
		model.UpdatedAt)
	defer rows.Close()
	group, err := pgx.CollectOneRow[GroupModel](rows, pgx.RowToStructByName[GroupModel])
	if errors.As(err, &r.PgError) {
		return &GroupModel{}, &util.ValidationError{
			Code:        fmt.Sprintf("PIDB-%s", r.Code),
			Status:      http.StatusInternalServerError,
			LogError:    []string{fmt.Sprintf("%s, %s", r.Message, r.Detail)},
			ClientError: []string{"Internal Server Error"},
		}
	}
	return &group, nil
}
func (r *GroupRepository) SelectOneGroupByNameAndCode(groupName, code string) (*GroupModel, *util.ValidationError) {
	query := `SELECT * FROM groups WHERE name = $1 AND code = $2`
	rows, _ := r.DB.Query(context.Background(), query, groupName, code)
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
