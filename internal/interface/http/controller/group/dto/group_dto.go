package group_dto

import (
	entity_group "github.com/andreis3/stores-ms/internal/domain/entity/group"
	"github.com/andreis3/stores-ms/internal/domain/valueobject"
)

type GroupInputDTO struct {
	GroupName string `json:"group_name"`
	Code      string `json:"code"`
	Status    string `json:"status"`
}
type GroupOutputDTO struct {
	ID        string `json:"id"`
	GroupName string `json:"group_name"`
	Code      string `json:"code"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (g *GroupInputDTO) MapperInputDtoToEntity() *entity_group.Group {
	status := valueobject.Status{
		Status: g.Status,
	}
	group := entity_group.NewGroup(g.GroupName, g.Code, &status)
	return group
}
