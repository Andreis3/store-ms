package iservices

import (
	entity_group "github.com/andreis3/stores-ms/internal/domain/entity"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type ICreateGroupService interface {
	CreateGroup(data entity_group.Group) (group_dto.GroupOutputDTO, *util.ValidationError)
}
