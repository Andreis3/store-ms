package uow_mock

import iuow "github.com/andreis3/stores-ms/internal/infra/uow/interfaces"

type RegisterRepository struct {
	Key  string
	Repo any
}

type MapRegisterRepository []RegisterRepository

func NewProxyUnitOfWorkMock(uowMock *UnitOfWorkMock, mapRepo MapRegisterRepository) *UnitOfWorkMock {
	uowMock.RepositoryMocks = make(map[string]iuow.RepositoryFactory)
	for _, repo := range mapRepo {
		uowMock.Register(repo.Key, func(tx any) any {
			return repo.Repo
		})
	}
	return uowMock
}
