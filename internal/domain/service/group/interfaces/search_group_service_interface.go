package igroup_service

import (
	iuow "github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type ISearchGroupService interface {
	SearchOneGroup(id string, unitOfWorker iuow.IUnitOfWork) (group_dto.GroupOutputDTO, *util.ValidationError)
}
