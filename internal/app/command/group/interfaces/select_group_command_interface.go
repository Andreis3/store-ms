package igroup_command

import (
	"github.com/andreis3/stores-ms/internal/interface/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type ISelectGroupCommand interface {
	Execute(id string) (group_dto.GroupOutputDTO, *util.ValidationError)
}
