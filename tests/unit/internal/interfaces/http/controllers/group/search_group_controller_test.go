//go:build unit
// +build unit

package group_controller_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/metric/prometheus"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/uuid_mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/interfaces/http/helpers"
	"github.com/andreis3/stores-ms/tests/mock/app/command/group"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/logger"
)

var _ = Describe("INTERFACES :: HTTP :: CONTROLLERS :: GROUP :: SEARCH_GROUP_CONTROLLER", func() {
	Describe("#SearchOneGroup", func() {
		selectGroupCommandMock := new(group_command_mock.SearchGroupCommandMock)
		prometheusMock := new(metric_prometheus_mock.PrometheusAdapterMock)
		loggerMock := new(logger_mock.LoggerMock)
		requestIDMock := new(uuid_mock.UUIDMock)
		groupController := new(group_controller.SearchGroupController)
		Context("When I call the method SelectGroup", func() {
			BeforeEach(func() {
				selectGroupCommandMock = new(group_command_mock.SearchGroupCommandMock)
				prometheusMock = new(metric_prometheus_mock.PrometheusAdapterMock)
				loggerMock = new(logger_mock.LoggerMock)
				requestIDMock = new(uuid_mock.UUIDMock)
			})
			It("Should return a error when of path paramnters is not format uuid", func() {
				ReturnErrorWhenIDInRouterPathIsNotFormatUUID(selectGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				groupController = group_controller.NewSearchGroupController(selectGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				request, err := http.NewRequest("GET", helpers.SEARCH_GROUP_V1, nil)
				request.SetPathValue("id", "123")
				writer := httptest.NewRecorder()

				expected := helpers.TypeResponseError{
					RequestID:    "123",
					CodeError:    "PR-001",
					StatusCode:   http.StatusBadRequest,
					ErrorMessage: []interface{}{"invalid parameter id"},
				}
				writer.Code = http.StatusBadRequest

				result := helpers.TypeResponseError{}

				groupController.SearchOneGroup(writer, request)

				json.Unmarshal(writer.Body.Bytes(), &result)

				Expect(err).To(BeNil())
				Expect(request).NotTo(BeNil())
				Expect(writer.Code).To(Equal(http.StatusBadRequest))
				Expect(result).To(Equal(expected))
				Expect(result.RequestID).To(BeAssignableToTypeOf(expected.RequestID))
				Expect(result.StatusCode).To(Equal(expected.StatusCode))
				Expect(loggerMock.AssertCalled(GinkgoT(), logger_mock.Error, "Get Group Error", []any{"REQUEST_ID", "123", "CODE_ERROR", "PR-001", "ORIGIN", "PathRouterValidate", "ERROR_MESSAGE", "invalid path parameter id"})).To(BeTrue())
				Expect(loggerMock.ExpectedCalls).To(HaveLen(1))
			})
			It("Should return a error when group not found", func() {
				ReturnErrorWhenGroupNotFound(selectGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				groupController = group_controller.NewSearchGroupController(selectGroupCommandMock, prometheusMock, loggerMock, requestIDMock)

				request, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/groups/%s", "7eef288f-dc7d-43b7-98a3-b5aacc717b8b"), nil)
				request.SetPathValue("id", "7eef288f-dc7d-43b7-98a3-b5aacc717b8b")
				writer := httptest.NewRecorder()

				expected := helpers.TypeResponseError{
					RequestID:    "123",
					CodeError:    "PR-002",
					StatusCode:   http.StatusNotFound,
					ErrorMessage: []interface{}{"group not found"},
				}
				writer.Code = http.StatusNotFound

				result := helpers.TypeResponseError{}

				groupController.SearchOneGroup(writer, request)

				json.Unmarshal(writer.Body.Bytes(), &result)

				Expect(err).To(BeNil())
				Expect(request).NotTo(BeNil())
				Expect(writer.Code).To(Equal(http.StatusNotFound))
				Expect(result).To(Equal(expected))
				Expect(result.RequestID).To(BeAssignableToTypeOf(expected.RequestID))
				Expect(result.StatusCode).To(Equal(expected.StatusCode))
				Expect(loggerMock.AssertCalled(GinkgoT(), logger_mock.Error, "Select One Group Error", []any{"REQUEST_ID", "123", "CODE_ERROR", "PR-002", "ORIGIN", "GetGroupService.SelectGroup", "ERROR_MESSAGE", "group not found"})).To(BeTrue())
				Expect(loggerMock.ExpectedCalls).To(HaveLen(1))
			})
			It("Should return any error when call the method SelectGroup", func() {
				ReturnErrorWhenCallSelectGroup(selectGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				groupController = group_controller.NewSearchGroupController(selectGroupCommandMock, prometheusMock, loggerMock, requestIDMock)

				request, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/groups/%s", "7eef288f-dc7d-43b7-98a3-b5aacc717b8b"), nil)
				request.SetPathValue("id", "7eef288f-dc7d-43b7-98a3-b5aacc717b8b")
				writer := httptest.NewRecorder()

				expected := helpers.TypeResponseError{
					RequestID:    "123",
					CodeError:    "PR-003",
					StatusCode:   http.StatusInternalServerError,
					ErrorMessage: []interface{}{"internal server error"},
				}
				writer.Code = http.StatusInternalServerError

				result := helpers.TypeResponseError{}

				groupController.SearchOneGroup(writer, request)

				json.Unmarshal(writer.Body.Bytes(), &result)

				Expect(err).To(BeNil())
				Expect(request).NotTo(BeNil())
				Expect(writer.Code).To(Equal(http.StatusInternalServerError))
				Expect(result).To(Equal(expected))
				Expect(result.RequestID).To(BeAssignableToTypeOf(expected.RequestID))
				Expect(result.StatusCode).To(Equal(expected.StatusCode))
				Expect(loggerMock.AssertCalled(GinkgoT(), logger_mock.Error, "Select One Group Error", []any{"REQUEST_ID", "123", "CODE_ERROR", "PR-003", "ORIGIN", "GetGroupService.SelectGroup", "ERROR_MESSAGE", "internal server error"})).To(BeTrue())
				Expect(loggerMock.ExpectedCalls).To(HaveLen(1))
			})
			It("Should return success when call the method SelectGroup", func() {
				ReturnSuccessWhenCallSelectGroup(selectGroupCommandMock, prometheusMock, requestIDMock)
				groupController = group_controller.NewSearchGroupController(selectGroupCommandMock, prometheusMock, loggerMock, requestIDMock)

				request, err := http.NewRequest("GET", fmt.Sprintf("/api/v1/groups/%s", "7eef288f-dc7d-43b7-98a3-b5aacc717b8b"), nil)
				request.SetPathValue("id", "7eef288f-dc7d-43b7-98a3-b5aacc717b8b")
				writer := httptest.NewRecorder()

				expected := helpers.TypeResponseSuccess{
					RequestID:  "123",
					StatusCode: http.StatusOK,
					Data: map[string]any{
						"id":         "7eef288f-dc7d-43b7-98a3-b5aacc717b8b",
						"name":       "Group 1",
						"code":       "G1",
						"status":     "active",
						"created_at": "2021-09-01T00:00:00Z",
						"updated_at": "2021-09-01T00:00:00Z",
					},
				}
				writer.Code = http.StatusOK

				result := helpers.TypeResponseSuccess{}

				groupController.SearchOneGroup(writer, request)

				json.Unmarshal(writer.Body.Bytes(), &result)

				Expect(err).To(BeNil())
				Expect(request).NotTo(BeNil())
				Expect(writer.Code).To(Equal(http.StatusOK))
				Expect(result).To(Equal(expected))
				Expect(result.RequestID).To(BeAssignableToTypeOf(expected.RequestID))
				Expect(result.Data).To(BeAssignableToTypeOf(expected.Data))
				//Expect(result.Name).To(Equal(expected.Name))
				//Expect(result.Code).To(Equal(expected.Code))
				Expect(loggerMock.ExpectedCalls).To(HaveLen(0))
			})
		})
	})
})
