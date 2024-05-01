package igroup_command

import (
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type ISearchGroupCommand interface {
	Execute(id string) (group_dto.GroupOutputDTO, *util.ValidationError)
}
