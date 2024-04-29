//go:build unit
// +build unit

package group_controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/andreis3/stores-ms/internal/interface/http/controllers/group/controller"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/metric/prometheus"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/uuid_mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/interface/http/helpers"
	"github.com/andreis3/stores-ms/tests/mock/app/command/group"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/logger"
)

var _ = Describe("INTERFACE :: HTTP :: CONTROLLERS :: GROUP :: GET_GROUP_CONTROLLER", func() {
	Describe("#SelectGroup", func() {
		selectGroupCommandMock := new(group_command_mock.SelectGroupCommandMock)
		prometheusMock := new(metric_prometheus_mock.PrometheusAdapterMock)
		loggerMock := new(logger_mock.LoggerMock)
		requestIDMock := new(uuid_mock.UUIDMock)
		groupController := new(group_controller.GetGroupController)
		Context("When I call the method SelectGroup", func() {
			BeforeEach(func() {
				selectGroupCommandMock = new(group_command_mock.SelectGroupCommandMock)
				prometheusMock = new(metric_prometheus_mock.PrometheusAdapterMock)
				loggerMock = new(logger_mock.LoggerMock)
				requestIDMock = new(uuid_mock.UUIDMock)
			})
			It("Should return a error when of path paramnters is not format uuid", func() {
				ReturnErroWhenIDInRouterPathIsNotFormatUUID(selectGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				groupController = group_controller.NewGetGroupController(selectGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				request, err := http.NewRequest("GET", fmt.Sprint(helpers.GET_GROUP_V1, "/123"), nil)
				writer := httptest.NewRecorder()

				expected := helpers.TypeResponseError{
					RequestID:    "123",
					CodeError:    "PR-001",
					StatusCode:   http.StatusBadRequest,
					ErrorMessage: []interface{}{"invalid parameter id"},
				}
				writer.Code = http.StatusBadRequest

				result := helpers.TypeResponseError{}

				groupController.GetGroup(writer, request)

				json.Unmarshal(writer.Body.Bytes(), &result)

				Expect(err).To(BeNil())
				Expect(request).NotTo(BeNil())
				Expect(writer.Code).To(Equal(http.StatusBadRequest))
				Expect(result).To(Equal(expected))
				Expect(result.RequestID).To(BeAssignableToTypeOf(expected.RequestID))
				Expect(result.StatusCode).To(Equal(expected.StatusCode))
				Expect(loggerMock.AssertCalled(GinkgoT(), logger_mock.Error, "Get Group Error", ([]any{"REQUEST_ID", "123", "CODE_ERROR", "PR-001", "ORIGIN", "PathRouterValidate", "ERROR_MESSAGE", "invalid path parameter id"}))).To(BeTrue())
				Expect(loggerMock.ExpectedCalls).To(HaveLen(1))
			})
		})
	})
})
