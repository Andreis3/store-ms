package igroup_service

import (
	entity_group "github.com/andreis3/stores-ms/internal/domain/entity/group"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type IInsertGroupService interface {
	InsertGroup(data entity_group.Group) (group_dto.GroupOutputDTO, *util.ValidationError)
}
