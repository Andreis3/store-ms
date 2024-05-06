package group_service

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type SearchGroupService struct {
}

func NewSearchGroupService() *SearchGroupService {
	return &SearchGroupService{}
}
func (s *SearchGroupService) SearchOneGroup(id string, unitOfWorker iuow.IUnitOfWork) (group_dto.GroupOutputDTO, *util.ValidationError) {
	var groupModel = new(repo_group.GroupModel)
	err := unitOfWorker.Do(func(uow iuow.IUnitOfWork) *util.ValidationError {
		var err *util.ValidationError
		groupRepository := unitOfWorker.GetRepository(util.GROUP_REPOSITORY_KEY).(irepo_group.IGroupRepository)
		groupModel, err = groupRepository.SelectOneGroupByID(id)
		if err != nil {
			return err
		}
		if groupModel.ID == nil {
			return &util.ValidationError{
				Code:        "VB-0003",
				Status:      http.StatusNotFound,
				LogError:    []string{"group not found"},
				ClientError: []string{"group not found"},
				Origin:      "SearchGroupService.SearchOneGroup",
			}
		}
		return nil
	})
	if err != nil {
		return group_dto.GroupOutputDTO{}, err
	}

	return group_dto.GroupOutputDTO{
		ID:        *groupModel.ID,
		Name:      *groupModel.Name,
		Code:      *groupModel.Code,
		Status:    *groupModel.Status,
		CreatedAt: util.FormatDateString(*groupModel.CreatedAt),
		UpdatedAt: util.FormatDateString(*groupModel.UpdatedAt),
	}, nil
}
