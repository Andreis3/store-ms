package group_controller_test

import (
	"context"
	"net/http"

	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/interface/http/helpers"
	"github.com/andreis3/stores-ms/internal/util"
	"github.com/andreis3/stores-ms/tests/mock/app/command/group"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/logger"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/metric/prometheus"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/uuid_mock"
)

// TODO: Create tests for the GetGroup method
func ReturnErroWhenIDInRouterPathIsNotFormatUUID(
	groupCommandMock *group_command_mock.SelectGroupCommandMock,
	prometheusMock *metric_prometheus_mock.PrometheusAdapterMock,
	loggerMock *logger_mock.LoggerMock,
	uuidMock *uuid_mock.UUIDMock) {
	groupCommandMock.On(group_command_mock.Execute, "123").Return(group_dto.GroupOutputDTO{}, &util.ValidationError{})
	uuidMock.On(uuid_mock.Generate).Return("123")
	loggerMock.On(logger_mock.Error, "Get Group Error", ([]any{"REQUEST_ID", "123", "CODE_ERROR", "PR-001", "ORIGIN", "PathRouterValidate", "ERROR_MESSAGE", "invalid path parameter id"}))
	prometheusMock.On(metric_prometheus_mock.CounterRequestHttpStatusCode, context.Background(), helpers.GET_GROUP_V1, http.StatusBadRequest)
	prometheusMock.On(metric_prometheus_mock.HistogramRequestDuration, context.Background(), helpers.GET_GROUP_V1, http.StatusBadRequest, float64(0))
}
