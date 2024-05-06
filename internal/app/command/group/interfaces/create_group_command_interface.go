package igroup_command

import (
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type ICreateGroupCommand interface {
	Execute(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError)
}
