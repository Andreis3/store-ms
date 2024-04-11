package group_dto

import (
	entity_group "github.com/andreis3/stores-ms/internal/domain/entity/group"
	"github.com/andreis3/stores-ms/internal/domain/valueobject"
)

type GroupInputDTO struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Status string `json:"status"`
}
type GroupOutputDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (g *GroupInputDTO) MapperInputDtoToEntity() *entity_group.Group {
	status := valueobject.Status{
		Status: g.Status,
	}
	group := entity_group.NewGroup(g.Name, g.Code, &status)
	return group
}
