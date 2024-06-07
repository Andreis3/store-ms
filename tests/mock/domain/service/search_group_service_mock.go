package service_mock

import (
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
	"github.com/stretchr/testify/mock"
)

const (
	SearchOneGroup = "SearchOneGroup"
)

type SearchGroupServiceMock struct {
	mock.Mock
}

func (s *SearchGroupServiceMock) SearchOneGroup(id string) (group_dto.GroupOutputDTO, *util.ValidationError) {
	args := s.Called(id)
	return args.Get(0).(group_dto.GroupOutputDTO), args.Get(1).(*util.ValidationError)
}
