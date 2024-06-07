package repo_group

import (
	"context"
	"fmt"
	entity_group "github.com/andreis3/stores-ms/internal/domain/entity"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"

	ipostgres "github.com/andreis3/stores-ms/internal/infra/adapters/database/postgres/interfaces"
	imetric "github.com/andreis3/stores-ms/internal/infra/common/metrics/interface"
	"github.com/andreis3/stores-ms/internal/util"
)

type GroupRepository struct {
	DB ipostgres.IInstructionDB
	*pgconn.PgError
	metrics imetric.IMetricAdapter
}

func NewGroupRepository(metrics imetric.IMetricAdapter) *GroupRepository {
	return &GroupRepository{
		metrics: metrics,
	}
}
func (r *GroupRepository) InsertGroup(data entity_group.Group) (*GroupModel, *util.ValidationError) {
	start := time.Now()
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
	//ERROR: duplicate key value violates unique constraint "groups_name_code_key" (SQLSTATE 23505)
	if errors.As(err, &r.PgError) {
		return &GroupModel{}, &util.ValidationError{
			Code:        fmt.Sprintf("PIDB-%s", r.Code),
			Origin:      "GroupRepository.CreateGroup",
			Status:      http.StatusInternalServerError,
			LogError:    []string{fmt.Sprintf("%s, %s", r.Message, r.Detail)},
			ClientError: []string{"Internal Server Error"},
		}
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	r.metrics.HistogramInstructionTableDuration(context.Background(), "postgres", "groups", "insert", duration)
	return &group, nil
}
func (r *GroupRepository) SelectOneGroupByNameAndCode(groupName, code string) (*GroupModel, *util.ValidationError) {
	start := time.Now()
	query := `SELECT * FROM groups WHERE name = $1 AND code = $2`
	rows, _ := r.DB.Query(context.Background(), query, groupName, code)
	defer rows.Close()
	group, err := pgx.CollectOneRow[GroupModel](rows, pgx.RowToStructByName[GroupModel])
	if errors.As(err, &r.PgError) {
		return nil, &util.ValidationError{
			Code:        fmt.Sprintf("PIDB-%s", r.Code),
			Origin:      "GroupRepository.SelectOneGroupByNameAndCode",
			Status:      http.StatusInternalServerError,
			LogError:    []string{fmt.Sprintf("%s, %s", r.Message, r.Detail)},
			ClientError: []string{"Internal Server Error"},
		}
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	r.metrics.HistogramInstructionTableDuration(context.Background(), "postgres", "groups", "select", duration)
	return &group, nil
}
func (r *GroupRepository) SelectOneGroupByID(id string) (*GroupModel, *util.ValidationError) {
	start := time.Now()
	query := `SELECT * FROM groups WHERE id = $1`
	rows, _ := r.DB.Query(context.Background(), query, id)
	defer rows.Close()
	group, err := pgx.CollectOneRow[GroupModel](rows, pgx.RowToStructByName[GroupModel])
	if errors.As(err, &r.PgError) {
		return nil, &util.ValidationError{
			Code:        fmt.Sprintf("PIDB-%s", r.Code),
			Origin:      "GroupRepository.SelectOneGroupByID",
			Status:      http.StatusInternalServerError,
			LogError:    []string{fmt.Sprintf("%s, %s", r.Message, r.Detail)},
			ClientError: []string{"Internal Server Error"},
		}
	}
	end := time.Now()
	duration := float64(end.Sub(start).Milliseconds())
	r.metrics.HistogramInstructionTableDuration(context.Background(), "postgres", "groups", "select", duration)
	return &group, nil
}
