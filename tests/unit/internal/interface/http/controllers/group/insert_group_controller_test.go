//go:build unit
// +build unit

package group_controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/controller"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/metric/prometheus"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/uuid_mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/interfaces/http/helpers"
	"github.com/andreis3/stores-ms/internal/util"
	"github.com/andreis3/stores-ms/tests/mock/app/command/group"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/logger"
)

var _ = Describe("INTERFACE :: HTTP :: CONTROLLERS :: GROUP :: INSERT_GROUP_CONTROLLER", func() {
	Describe("#CreateGroup", func() {
		insertGroupCommandMock := new(group_command_mock.InsertGroupCommandMock)
		prometheusMock := new(metric_prometheus_mock.PrometheusAdapterMock)
		loggerMock := new(logger_mock.LoggerMock)
		requestIDMock := new(uuid_mock.UUIDMock)
		groupController := new(group_controller.CreateGroupController)
		Context("When I call the method CreateGroup", func() {
			BeforeEach(func() {
				insertGroupCommandMock = new(group_command_mock.InsertGroupCommandMock)
				prometheusMock = new(metric_prometheus_mock.PrometheusAdapterMock)
				loggerMock = new(logger_mock.LoggerMock)
				requestIDMock = new(uuid_mock.UUIDMock)
			})
			It("Should return a error when of method insertGroupCommand.Execute is call", func() {
				ReturnErroWhenInsertGroupCommandOfExecuteIsCalled(insertGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				groupController = group_controller.NewCreateGroupController(insertGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				body := `{
							"name":"teste 1",
							"code": "23",
							"status":"active"
						  }`
				request, err := http.NewRequest("POST", helpers.CREATE_GROUP_V1, strings.NewReader(body))
				writer := httptest.NewRecorder()
				expected := helpers.TypeResponseError{
					RequestID:    "123",
					CodeError:    "VBR-400",
					StatusCode:   http.StatusBadRequest,
					ErrorMessage: []interface{}{"error test"},
				}
				result := helpers.TypeResponseError{}
				groupController.CreateGroup(writer, request)
				json.Unmarshal(writer.Body.Bytes(), &result)

				Expect(err).To(BeNil())
				Expect(request).NotTo(BeNil())
				Expect(writer.Code).To(Equal(http.StatusBadRequest))
				Expect(result).To(Equal(expected))
				Expect(result.RequestID).To(BeAssignableToTypeOf(expected.RequestID))
				Expect(result.StatusCode).To(Equal(expected.StatusCode))
				Expect(loggerMock.AssertCalled(GinkgoT(), logger_mock.Error, "Create Group Error", ([]any{"REQUEST_ID", "123", "CODE_ERROR", "VBR-400", "ORIGIN", "InsertGroupCommand", "ERROR_MESSAGE", "error test"}))).To(BeTrue())
				Expect(loggerMock.ExpectedCalls).To(HaveLen(1))
			})

			It("Should return a error when of method DecoderBodyRequest is call with poorly formatted payload", func() {
				ReturnErroDecoderBodyRequestPoorlyFormattedPayload(insertGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				groupController = group_controller.NewCreateGroupController(insertGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				body := `{
							"group_name":"teste 1",
							"code": "23",
							"status":"active"
						  `
				request, err := http.NewRequest("POST", helpers.CREATE_GROUP_V1, strings.NewReader(body))
				writer := httptest.NewRecorder()
				expected := helpers.TypeResponseError{
					RequestID:    "123",
					CodeError:    "DJ-402",
					StatusCode:   http.StatusBadRequest,
					ErrorMessage: []interface{}{"invalid json"},
				}
				result := helpers.TypeResponseError{}
				groupController.CreateGroup(writer, request)
				json.Unmarshal(writer.Body.Bytes(), &result)

				Expect(err).To(BeNil())
				Expect(request).NotTo(BeNil())
				Expect(writer.Code).To(Equal(http.StatusBadRequest))
				Expect(result).To(Equal(expected))
				Expect(result.RequestID).To(BeAssignableToTypeOf(expected.RequestID))
				Expect(result.StatusCode).To(Equal(expected.StatusCode))
				Expect(loggerMock.AssertCalled(GinkgoT(), logger_mock.Error, "Create Group Error", ([]any{"REQUEST_ID", "123", "CODE_ERROR", "DJ-402", "ORIGIN", "DecoderBodyRequest", "ERROR_MESSAGE", "unexpected EOF"}))).To(BeTrue())
				Expect(loggerMock.ExpectedCalls).To(HaveLen(1))
			})

			It("Should return a error when of method DecoderBodyRequest is call with invalid json syntax", func() {
				ReturnErroWhenDecoderBodyRequestInvalidJsonSyntax(insertGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				groupController = group_controller.NewCreateGroupController(insertGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				body := `{
							"group_name":"teste 1",
							"code": "23",
							"status":"active",
						  }`
				request, err := http.NewRequest("POST", helpers.CREATE_GROUP_V1, strings.NewReader(body))
				writer := httptest.NewRecorder()

				expected := helpers.TypeResponseError{
					RequestID:    "123",
					CodeError:    "DJ-400",
					StatusCode:   http.StatusBadRequest,
					ErrorMessage: []interface{}{"invalid json syntax"},
				}

				result := helpers.TypeResponseError{}

				groupController.CreateGroup(writer, request)

				json.Unmarshal(writer.Body.Bytes(), &result)

				Expect(err).To(BeNil())
				Expect(request).NotTo(BeNil())
				Expect(writer.Code).To(Equal(http.StatusBadRequest))
				Expect(result).To(Equal(expected))
				Expect(result.RequestID).To(BeAssignableToTypeOf(expected.RequestID))
				Expect(result.StatusCode).To(Equal(expected.StatusCode))
				Expect(loggerMock.AssertCalled(GinkgoT(), logger_mock.Error, "Create Group Error", ([]any{"REQUEST_ID", "123", "CODE_ERROR", "DJ-400", "ORIGIN", "DecoderBodyRequest", "ERROR_MESSAGE", "invalid character '}' looking for beginning of object key string"}))).To(BeTrue())
				Expect(loggerMock.ExpectedCalls).To(HaveLen(1))
			})

			It("Should return a error when of method DecoderBodyRequest is call with invalid json field type", func() {
				ReturnErroWhenDecoderBodyRequestInvalidJsonFieldType(insertGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				groupController = group_controller.NewCreateGroupController(insertGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				body := `{
							"group_name":"teste 1",
							"code": 23,
							"status":"active"
						  }`
				request, err := http.NewRequest("POST", helpers.CREATE_GROUP_V1, strings.NewReader(body))
				writer := httptest.NewRecorder()

				expected := helpers.TypeResponseError{
					RequestID:    "123",
					CodeError:    "DJ-401",
					StatusCode:   http.StatusBadRequest,
					ErrorMessage: []interface{}{"invalid json field type"},
				}

				result := helpers.TypeResponseError{}

				groupController.CreateGroup(writer, request)

				json.Unmarshal(writer.Body.Bytes(), &result)

				Expect(err).To(BeNil())
				Expect(request).NotTo(BeNil())
				Expect(writer.Code).To(Equal(http.StatusBadRequest))
				Expect(result).To(Equal(expected))
				Expect(result.RequestID).To(BeAssignableToTypeOf(expected.RequestID))
				Expect(result.StatusCode).To(Equal(expected.StatusCode))
				Expect(loggerMock.AssertCalled(GinkgoT(), logger_mock.Error, "Create Group Error", ([]any{"REQUEST_ID", "123", "CODE_ERROR", "DJ-401", "ORIGIN", "DecoderBodyRequest", "ERROR_MESSAGE", "json: cannot unmarshal number into Go struct field GroupInputDTO.code of type string"}))).To(BeTrue())
				Expect(loggerMock.ExpectedCalls).To(HaveLen(1))
			})

			It("Should create a new group without errors", func() {
				ReturnSuccessWhenInsertGroupCommandOfExecuteIsCalled(insertGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				data := group_dto.GroupInputDTO{
					Name:   "teste 1",
					Code:   "23",
					Status: "active",
				}
				groupController = group_controller.NewCreateGroupController(insertGroupCommandMock, prometheusMock, loggerMock, requestIDMock)
				body := `{
    						"name":"teste 1",
							"code": "23",
    						"status":"active"
						  }`
				request, err := http.NewRequest("POST", helpers.CREATE_GROUP_V1, strings.NewReader(body))
				writer := httptest.NewRecorder()

				expected := helpers.TypeResponseSuccess{
					RequestID:  "123",
					StatusCode: http.StatusCreated,
					Data: map[string]any{
						"created_at": "23/09/2021 10:00:00",
						"updated_at": "23/09/2021 10:00:00",
						"id":         "123",
						"name":       "test 1",
						"code":       "23",
						"status":     "active",
					},
				}

				result := helpers.TypeResponseSuccess{}

				groupController.CreateGroup(writer, request)

				json.Unmarshal(writer.Body.Bytes(), &result)

				Expect(err).To(BeNil())
				Expect(request).NotTo(BeNil())
				Expect(writer.Code).To(Equal(http.StatusCreated))
				Expect(result).To(Equal(expected))
				Expect(result.RequestID).To(BeAssignableToTypeOf(expected.RequestID))
				Expect(result.StatusCode).To(Equal(expected.StatusCode))
				Expect(insertGroupCommandMock.AssertCalled(GinkgoT(), group_command_mock.Execute, data)).To(BeTrue())
				Expect(insertGroupCommandMock.MethodCalled(group_command_mock.Execute, data)).To(ContainElements(group_dto.GroupOutputDTO{
					ID:        "123",
					Name:      "test 1",
					Code:      "23",
					Status:    "active",
					CreatedAt: "23/09/2021 10:00:00",
					UpdatedAt: "23/09/2021 10:00:00",
				}, (*util.ValidationError)(nil)))

			})
		})
	})
})
