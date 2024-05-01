package group_service

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/domain/entity/group"
	"github.com/andreis3/stores-ms/internal/infra/common/uuid"
	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group/interfaces"
	"github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

type InsertGroupService struct {
	uow  iuow.IUnitOfWork
	uuid uuid.IUUID
}

func NewInsertGroupService(uow iuow.IUnitOfWork, uuid uuid.IUUID) *InsertGroupService {
	return &InsertGroupService{
		uow:  uow,
		uuid: uuid,
	}
}
func (igs *InsertGroupService) InsertGroup(data entity_group.Group) (group_dto.GroupOutputDTO, *util.ValidationError) {
	groupModel := new(repo_group.GroupModel)
	data.ID = igs.uuid.Generate()
	validate := data.Validate()
	if validate.HasNotification() {
		return group_dto.GroupOutputDTO{}, &util.ValidationError{
			Code:        "VBR-0001",
			Origin:      "InsertGroupService.InsertGroup",
			LogError:    validate.ReturnNotification(),
			ClientError: validate.ReturnNotification(),
			Status:      http.StatusBadRequest,
		}
	}
	err := igs.uow.Do(func(uow iuow.IUnitOfWork) *util.ValidationError {
		repoGroup := uow.GetRepository(util.GROUP_REPOSITORY_KEY).(irepo_group.IGroupRepository)
		res, err := repoGroup.SelectOneGroupByNameAndCode(data.Name, data.Code)
		if err != nil {
			return err
		}
		if res.ID != nil {
			return &util.ValidationError{
				Code:        "VBR-0002",
				Origin:      "InsertGroupService.InsertGroup",
				LogError:    []string{"Group already exists"},
				ClientError: []string{"Group already exists"},
				Status:      http.StatusBadRequest,
			}
		}
		groupModel, err = repoGroup.InsertGroup(data)
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
		Name:      *groupModel.Name,
		CreatedAt: util.FormatDateString(*groupModel.CreatedAt),
		UpdatedAt: util.FormatDateString(*groupModel.UpdatedAt),
	}, nil
}
