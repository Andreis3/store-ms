package group_controller_test

import (
	"context"
	"net/http"

	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/interfaces/http/helpers"
	"github.com/andreis3/stores-ms/internal/util"
	"github.com/andreis3/stores-ms/tests/mock/app/command/group"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/logger"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/metric/prometheus"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/uuid_mock"
)

// TODO: Create tests for the SearchGroup method
func ReturnErroWhenIDInRouterPathIsNotFormatUUID(
	groupCommandMock *group_command_mock.SearchGroupCommandMock,
	prometheusMock *metric_prometheus_mock.PrometheusAdapterMock,
	loggerMock *logger_mock.LoggerMock,
	uuidMock *uuid_mock.UUIDMock) {
	groupCommandMock.On(group_command_mock.Execute, "123").Return(group_dto.GroupOutputDTO{}, &util.ValidationError{})
	uuidMock.On(uuid_mock.Generate).Return("123")
	loggerMock.On(logger_mock.Error, "Get Group Error", ([]any{"REQUEST_ID", "123", "CODE_ERROR", "PR-001", "ORIGIN", "PathRouterValidate", "ERROR_MESSAGE", "invalid path parameter id"}))
	prometheusMock.On(metric_prometheus_mock.CounterRequestHttpStatusCode, context.Background(), helpers.GET_GROUP_V1, http.StatusBadRequest)
	prometheusMock.On(metric_prometheus_mock.HistogramRequestDuration, context.Background(), helpers.GET_GROUP_V1, http.StatusBadRequest, float64(0))
}

func ReturnErroWhenGroupNotFound(
	groupCommandMock *group_command_mock.SearchGroupCommandMock,
	prometheusMock *metric_prometheus_mock.PrometheusAdapterMock,
	loggerMock *logger_mock.LoggerMock,
	uuidMock *uuid_mock.UUIDMock) {
	groupCommandMock.On(group_command_mock.Execute, "7eef288f-dc7d-43b7-98a3-b5aacc717b8b").Return(group_dto.GroupOutputDTO{}, &util.ValidationError{
		Code:        "PR-002",
		Origin:      "GetGroupService.SelectGroup",
		ClientError: []string{"group not found"},
		LogError:    []string{"group not found"},
		Status:      http.StatusNotFound,
	})
	uuidMock.On(uuid_mock.Generate).Return("123")
	loggerMock.On(logger_mock.Error, "Select One Group Error", ([]any{"REQUEST_ID", "123", "CODE_ERROR", "PR-002", "ORIGIN", "GetGroupService.SelectGroup", "ERROR_MESSAGE", "group not found"}))
	prometheusMock.On(metric_prometheus_mock.CounterRequestHttpStatusCode, context.Background(), helpers.GET_GROUP_V1, http.StatusNotFound)
	prometheusMock.On(metric_prometheus_mock.HistogramRequestDuration, context.Background(), helpers.GET_GROUP_V1, http.StatusNotFound, float64(0))
}

func ReturnErroWhenCallSelectGroup(
	groupCommandMock *group_command_mock.SearchGroupCommandMock,
	prometheusMock *metric_prometheus_mock.PrometheusAdapterMock,
	loggerMock *logger_mock.LoggerMock,
	uuidMock *uuid_mock.UUIDMock) {
	groupCommandMock.On(group_command_mock.Execute, "7eef288f-dc7d-43b7-98a3-b5aacc717b8b").Return(group_dto.GroupOutputDTO{}, &util.ValidationError{
		Code:        "PR-003",
		Origin:      "GetGroupService.SelectGroup",
		ClientError: []string{"internal server error"},
		LogError:    []string{"internal server error"},
		Status:      http.StatusInternalServerError,
	})
	uuidMock.On(uuid_mock.Generate).Return("123")
	loggerMock.On(logger_mock.Error, "Select One Group Error", ([]any{"REQUEST_ID", "123", "CODE_ERROR", "PR-003", "ORIGIN", "GetGroupService.SelectGroup", "ERROR_MESSAGE", "internal server error"}))
	prometheusMock.On(metric_prometheus_mock.CounterRequestHttpStatusCode, context.Background(), helpers.GET_GROUP_V1, http.StatusInternalServerError)
	prometheusMock.On(metric_prometheus_mock.HistogramRequestDuration, context.Background(), helpers.GET_GROUP_V1, http.StatusInternalServerError, float64(0))
}

func ReturnSuccessWhenCallSelectGroup(
	groupCommandMock *group_command_mock.SearchGroupCommandMock,
	prometheusMock *metric_prometheus_mock.PrometheusAdapterMock,
	loggerMock *logger_mock.LoggerMock,
	uuidMock *uuid_mock.UUIDMock) {
	groupCommandMock.On(group_command_mock.Execute, "7eef288f-dc7d-43b7-98a3-b5aacc717b8b").Return(group_dto.GroupOutputDTO{
		ID:        "7eef288f-dc7d-43b7-98a3-b5aacc717b8b",
		Name:      "Group 1",
		Code:      "G1",
		Status:    "active",
		CreatedAt: "2021-09-01T00:00:00Z",
		UpdatedAt: "2021-09-01T00:00:00Z",
	}, (*util.ValidationError)(nil))
	uuidMock.On(uuid_mock.Generate).Return("123")
	prometheusMock.On(metric_prometheus_mock.CounterRequestHttpStatusCode, context.Background(), helpers.GET_GROUP_V1, http.StatusOK)
	prometheusMock.On(metric_prometheus_mock.HistogramRequestDuration, context.Background(), helpers.GET_GROUP_V1, http.StatusOK, float64(0))
}
