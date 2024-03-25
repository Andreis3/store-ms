package repo_group

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupRepository struct {
	db *pgxpool.Pool
}

func NewGroupRepository(db *pgxpool.Pool) *GroupRepository {
	return &GroupRepository{
		db: db,
	}
}

func (r *GroupRepository) InsertGroup(group GroupModel) (string, error) {
	query := `INSERT INTO groups (id, group_name, code, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var id string

	err := r.db.QueryRow(context.Background(), query, group.ID, group.GroupName, group.Code, group.Status).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}
