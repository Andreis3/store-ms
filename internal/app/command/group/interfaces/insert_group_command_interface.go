package igroup_command

import (
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type IInsertGroupCommand interface {
	Execute(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError)
}
