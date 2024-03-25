package group_service

import (
	"context"

	entity_group "github.com/andreis3/stores-ms/internal/domain/entity/group"
	"github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"
	iuow "github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
	"github.com/andreis3/stores-ms/internal/util"
)

type InsertGroupService struct {
	uow       iuow.IUnitOfWork
	repoGroup repo_group.GroupRepository
	ctx       context.Context
}

func NewInsertGroupService(uow iuow.IUnitOfWork) *InsertGroupService {
	ctx := context.Background()
	repoGroup := uow.GetRepository(ctx, util.GROUP_REPOSITORY_KEY).(repo_group.GroupRepository)
	return &InsertGroupService{
		uow:       uow,
		repoGroup: repoGroup,
		ctx:       ctx,
	}
}

func (s *InsertGroupService) InsertGroup(data entity_group.Group) (entity_group.Group, error) {
	err := s.uow.Do(s.ctx, func(tx iuow.IUnitOfWork) error {
		groupModel := repo_group.MapperGroupModel(data)
		_, err := s.repoGroup.InsertGroup(*groupModel)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return entity_group.Group{}, err
	}

	return data, nil
}
