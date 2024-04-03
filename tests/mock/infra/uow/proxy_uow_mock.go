package uow_mock

type RegisterRepository struct {
	Key  string
	Repo any
}

type MapRegisterRepository []RegisterRepository

func NewProxyUnitOfWorkMock(register MapRegisterRepository) *UnitOfWorkMock {
	uowMock := NewUnitOfWorkMock()

	for _, value := range register {
		uowMock.Register(value.Key, func(tx any) any {
			return value.Repo
		})
	}

	return uowMock
}
