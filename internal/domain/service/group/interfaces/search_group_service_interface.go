package igroup_service

import (
	"github.com/andreis3/stores-ms/internal/interface/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type ISearchGroupService interface {
	SearchGroup(id string) (group_dto.GroupOutputDTO, *util.ValidationError)
}
