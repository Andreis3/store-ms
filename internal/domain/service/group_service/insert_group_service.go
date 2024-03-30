package group_service

import (
	"context"
	"net/http"

	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"
	irepo_group "github.com/andreis3/stores-ms/internal/infra/repository/postgres/group/interfaces"
	iuow "github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
	group_dto "github.com/andreis3/stores-ms/internal/interface/http/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type InsertGroupService struct {
	uow       iuow.IUnitOfWork
	repoGroup irepo_group.IGroupRepository
	ctx       context.Context
}

func NewInsertGroupService(uow iuow.IUnitOfWork) *InsertGroupService {
	ctx := context.Background()
	repoGroup := uow.GetRepository(util.GROUP_REPOSITORY_KEY).(irepo_group.IGroupRepository)
	return &InsertGroupService{
		uow:       uow,
		ctx:       ctx,
		repoGroup: repoGroup,
	}
}
func (s *InsertGroupService) InsertGroup(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
	var groupModel *repo_group.GroupModel
	groupEntity := data.MapperInputDtoToEntity()
	validate := groupEntity.Validate()
	if len(validate) > 0 {
		return group_dto.GroupOutputDTO{}, &util.ValidationError{
			Code:        "VBR-0001",
			LogError:    validate,
			ClientError: validate,
			Status:      http.StatusBadRequest,
		}
	}
	err := s.uow.Do(s.ctx, func(tx iuow.IUnitOfWork) *util.ValidationError {
		groupModel = repo_group.MapperGroupModel(*groupEntity)
		_, err := s.repoGroup.InsertGroup(*groupModel)
		if err != nil {
			return err
		}
		return nil
	})
	if err.ExistError() {
		return group_dto.GroupOutputDTO{}, err
	}
	return group_dto.GroupOutputDTO{
		ID:        groupModel.ID,
		Status:    groupModel.Status,
		Code:      groupModel.Code,
		GroupName: groupModel.GroupName,
		CreatedAt: util.FormatDateString(groupModel.CreatedAt),
		UpdatedAt: util.FormatDateString(groupModel.UpdatedAt),
	}, nil
}
