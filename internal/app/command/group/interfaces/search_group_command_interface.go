package igroup_command

import (
	iuow "github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type ISearchGroupCommand interface {
	Execute(id string, uow iuow.IUnitOfWork) (group_dto.GroupOutputDTO, *util.ValidationError)
}
