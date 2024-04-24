package group_controller_test

import (
	"context"
	"net/http"

	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
	"github.com/andreis3/stores-ms/tests/mock/app/command/group"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/logger"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/metric/prometheus"
	"github.com/andreis3/stores-ms/tests/mock/interface/http/helpers"
)

func ReturnErroWhenInsertGroupCommandOfExecuteIsCalled(
	groupCommandMock *group_command_mock.InsertGroupCommandMock,
	prometheusMock *metric_prometheus_mock.PrometheusAdapterMock,
	loggerMock *logger_mock.LoggerMock,
	requestIDMock *helpers_mock.RequestIDMock) {
	data := group_dto.GroupInputDTO{
		Name:   "teste 1",
		Code:   "23",
		Status: "active",
	}
	groupCommandMock.On(group_command_mock.Execute, data).Return(group_dto.GroupOutputDTO{}, &util.ValidationError{
		Code:        "VBR-400",
		Status:      http.StatusBadRequest,
		ClientError: []string{"error test"},
		LogError:    []string{"error test"},
	})
	requestIDMock.On(helpers_mock.Generate).Return("123")
	loggerMock.On(logger_mock.Error, "Create Group Error", ([]any{"REQUEST_ID", "123", "CODE_ERROR", "VBR-400", "ERROR_MESSAGE", "error test"}))
	prometheusMock.On(metric_prometheus_mock.CounterRequestHttpStatusCode, context.Background(), "/groups", http.StatusBadRequest)
	prometheusMock.On(metric_prometheus_mock.HistogramRequestDuration, context.Background(), "/groups", http.StatusBadRequest, float64(0))
}
