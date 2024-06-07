package iservices

import (
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type ISearchGroupService interface {
	SearchOneGroup(id string) (group_dto.GroupOutputDTO, *util.ValidationError)
}
