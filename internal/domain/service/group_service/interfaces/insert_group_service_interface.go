package igroup_service

import (
	"github.com/andreis3/stores-ms/internal/interface/http/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type IInsertGroupService interface {
	InsertGroup(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError)
}
