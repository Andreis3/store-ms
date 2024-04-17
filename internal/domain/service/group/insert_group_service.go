package group_service

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type InsertGroupService struct {
	uow iuow.IUnitOfWork
}

func NewInsertGroupService(uow iuow.IUnitOfWork) *InsertGroupService {
	return &InsertGroupService{
		uow: uow,
	}
}
func (igs *InsertGroupService) InsertGroup(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
	var groupModel *repo_group.GroupModel
	groupEntity := data.MapperInputDtoToEntity()
	validate := groupEntity.Validate()
	if validate.HasNotification() {
		return group_dto.GroupOutputDTO{}, &util.ValidationError{
			Code:        "VBR-0001",
			LogError:    validate.ReturnNotification(),
			ClientError: validate.ReturnNotification(),
			Status:      http.StatusBadRequest,
		}
	}
	err := igs.uow.Do(func(tx iuow.IUnitOfWork) *util.ValidationError {
		repoGroup := igs.uow.GetRepository(util.GROUP_REPOSITORY_KEY).(irepo_group.IGroupRepository)
		groupModel = repo_group.MapperGroupModel(*groupEntity)
		_, err := repoGroup.InsertGroup(*groupModel)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return group_dto.GroupOutputDTO{}, err
	}
	return group_dto.GroupOutputDTO{
		ID:        *groupModel.ID,
		Status:    *groupModel.Status,
		Code:      *groupModel.Code,
		GroupName: *groupModel.GroupName,
		CreatedAt: util.FormatDateString(*groupModel.CreatedAt),
		UpdatedAt: util.FormatDateString(*groupModel.UpdatedAt),
	}, nil
}
