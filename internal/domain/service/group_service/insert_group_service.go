package group_service

import iuow "github.com/andreis3/stores-ms/internal/infra/uow/interfaces"

type InsertGroupService struct {
	uow iuow.IUnitOfWork
}
