//go:build unit
// +build unit

package group_controller_test

import (
	"context"
	"encoding/json"
	metric_prometheus_mock "github.com/andreis3/stores-ms/tests/mock/infra/common/metric/prometheus"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	group_controller "github.com/andreis3/stores-ms/internal/interface/http/controller/group"
	group_dto "github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/interface/http/helpers"
	"github.com/andreis3/stores-ms/internal/util"
	group_command_mock "github.com/andreis3/stores-ms/tests/mock/app/command/group"
	logger_mock "github.com/andreis3/stores-ms/tests/mock/infra/common/logger"
	helpers_mock "github.com/andreis3/stores-ms/tests/mock/interface/http/helpers"
)

func Test_GroupControllerSuite(t *testing.T) {
	suiteConfig, reporterConfig := GinkgoConfiguration()

	suiteConfig.SkipStrings = []string{"NEVER-RUN"}
	reporterConfig.FullTrace = true
	reporterConfig.Succinct = true

	RegisterFailHandler(Fail)
	RunSpecs(t, "Group Controller Test Suite ", suiteConfig, reporterConfig)
}

var _ = Describe("INTERFACE :: HTTP :: CONTROLLERS :: GROUP :: GROUP_CONTROLLER", func() {
	Describe("#CreateGroup", func() {
		var groupCommandMock *group_command_mock.InsertGroupCommandMock
		var prometheusMock *metric_prometheus_mock.PrometheusAdapterMock
		var loggerMock *logger_mock.LoggerMock
		var requestIDMock *helpers_mock.RequestIDMock
		var groupController *group_controller.Controller
		Context("When I call the method CreateGroup", func() {
			It("Should return a error when of method insertGroupCommand.Execute is call", func() {
				groupCommandMock = &group_command_mock.InsertGroupCommandMock{
					ExecuteFunc: func(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
						return group_dto.GroupOutputDTO{}, &util.ValidationError{
							Code:        "VBR-400",
							Status:      http.StatusBadRequest,
							ClientError: []string{"error test"},
							LogError:    []string{"error test"},
						}
					},
				}
				requestIDMock = &helpers_mock.RequestIDMock{
					GenerateFunc: func() string {
						return "123"
					},
				}
				loggerMock = &logger_mock.LoggerMock{
					ErrorFunc: func(message string, keysAndValues ...any) {
						return
					},
				}
				prometheusMock = &metric_prometheus_mock.PrometheusAdapterMock{
					CounterRequestHttpStatusCodeFunc: func(ctx context.Context, router string, statusCode int) {},
					HistogramRequestDurationFunc:     func(ctx context.Context, router string, statusCode int, duration float64) {},
				}

				groupController = group_controller.NewGroupController(groupCommandMock, prometheusMock, loggerMock, requestIDMock)
				body := `{
							"group_name":"teste 1",
							"code": "23",
							"status":"active"
						  }`
				request, err := http.NewRequest("POST", "/group", strings.NewReader(body))
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
				Expect(loggerMock.ErrorMSG).To(Equal("Create Group Error"))
				Expect(loggerMock.ErrorAny).To(Equal([]interface{}{"REQUEST_ID", "123", "CODE_ERROR", "VBR-400", "ERROR_MESSAGE", "error test"}))
			})

			It("Should return a error when of method DecoderBodyRequest is call with poorly formatted payload", func() {
				groupCommandMock = &group_command_mock.InsertGroupCommandMock{
					ExecuteFunc: func(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
						return group_dto.GroupOutputDTO{}, &util.ValidationError{}
					},
				}
				requestIDMock = &helpers_mock.RequestIDMock{
					GenerateFunc: func() string {
						return "123"
					},
				}
				loggerMock = &logger_mock.LoggerMock{
					ErrorFunc: func(message string, keysAndValues ...any) {
						return
					},
				}
				prometheusMock = &metric_prometheus_mock.PrometheusAdapterMock{
					CounterRequestHttpStatusCodeFunc: func(ctx context.Context, router string, statusCode int) {},
					HistogramRequestDurationFunc:     func(ctx context.Context, router string, statusCode int, duration float64) {},
				}

				groupController = group_controller.NewGroupController(groupCommandMock, prometheusMock, loggerMock, requestIDMock)
				body := `{
							"group_name":"teste 1",
							"code": "23",
							"status":"active"
						  `
				request, err := http.NewRequest("POST", "/group", strings.NewReader(body))
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
				Expect(loggerMock.ErrorMSG).To(Equal("Create Group Error"))
				Expect(loggerMock.ErrorAny).To(Equal([]interface{}{"REQUEST_ID", "123", "CODE_ERROR", "DJ-402", "ERROR_MESSAGE", "unexpected EOF"}))
			})

			It("Should return a error when of method DecoderBodyRequest is call with invalid json syntax", func() {
				groupCommandMock = &group_command_mock.InsertGroupCommandMock{
					ExecuteFunc: func(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
						return group_dto.GroupOutputDTO{}, &util.ValidationError{}
					},
				}
				requestIDMock = &helpers_mock.RequestIDMock{
					GenerateFunc: func() string {
						return "123"
					},
				}
				loggerMock = &logger_mock.LoggerMock{
					ErrorFunc: func(message string, keysAndValues ...any) {
						return
					},
				}
				prometheusMock = &metric_prometheus_mock.PrometheusAdapterMock{
					CounterRequestHttpStatusCodeFunc: func(ctx context.Context, router string, statusCode int) {},
					HistogramRequestDurationFunc:     func(ctx context.Context, router string, statusCode int, duration float64) {},
				}

				groupController = group_controller.NewGroupController(groupCommandMock, prometheusMock, loggerMock, requestIDMock)
				body := `{
							"group_name":"teste 1",
							"code": "23",
							"status":"active",
						  }`
				request, err := http.NewRequest("POST", "/group", strings.NewReader(body))
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
				Expect(loggerMock.ErrorMSG).To(Equal("Create Group Error"))
				Expect(loggerMock.ErrorAny).To(Equal([]interface{}{"REQUEST_ID", "123", "CODE_ERROR", "DJ-400", "ERROR_MESSAGE", "invalid character '}' looking for beginning of object key string"}))
			})

			It("Should return a error when of method DecoderBodyRequest is call with invalid json field type", func() {
				groupCommandMock = &group_command_mock.InsertGroupCommandMock{
					ExecuteFunc: func(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
						return group_dto.GroupOutputDTO{}, &util.ValidationError{}
					},
				}
				requestIDMock = &helpers_mock.RequestIDMock{
					GenerateFunc: func() string {
						return "123"
					},
				}
				loggerMock = &logger_mock.LoggerMock{
					ErrorFunc: func(message string, keysAndValues ...any) {
						return
					},
				}
				prometheusMock = &metric_prometheus_mock.PrometheusAdapterMock{
					CounterRequestHttpStatusCodeFunc: func(ctx context.Context, router string, statusCode int) {},
					HistogramRequestDurationFunc:     func(ctx context.Context, router string, statusCode int, duration float64) {},
				}

				groupController = group_controller.NewGroupController(groupCommandMock, prometheusMock, loggerMock, requestIDMock)
				body := `{
							"group_name":"teste 1",
							"code": 23,
							"status":"active"
						  }`
				request, err := http.NewRequest("POST", "/group", strings.NewReader(body))
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
				Expect(loggerMock.ErrorMSG).To(Equal("Create Group Error"))
				Expect(loggerMock.ErrorAny).To(Equal([]interface{}{"REQUEST_ID", "123", "CODE_ERROR", "DJ-401", "ERROR_MESSAGE", "json: cannot unmarshal number into Go struct field GroupInputDTO.code of type string"}))
			})

			It("Should create a new group without errors", func() {
				groupCommandMock = &group_command_mock.InsertGroupCommandMock{
					ExecuteFunc: func(data group_dto.GroupInputDTO) (group_dto.GroupOutputDTO, *util.ValidationError) {
						return group_dto.GroupOutputDTO{
							ID:        "123",
							GroupName: "test 1",
							Code:      "23",
							Status:    "active",
							CreatedAt: "23/09/2021 10:00:00",
							UpdatedAt: "23/09/2021 10:00:00",
						}, nil
					},
				}
				requestIDMock = &helpers_mock.RequestIDMock{
					GenerateFunc: func() string {
						return "123"
					},
				}
				prometheusMock = &metric_prometheus_mock.PrometheusAdapterMock{
					CounterRequestHttpStatusCodeFunc: func(ctx context.Context, router string, statusCode int) {},
					HistogramRequestDurationFunc:     func(ctx context.Context, router string, statusCode int, duration float64) {},
				}
				loggerMock = &logger_mock.LoggerMock{}

				groupController = group_controller.NewGroupController(groupCommandMock, prometheusMock, loggerMock, requestIDMock)
				body := `{
    						"group_name":"teste 1",
							"code": "23",
    						"status":"active"
						  }`
				request, err := http.NewRequest("POST", "/group", strings.NewReader(body))
				writer := httptest.NewRecorder()

				expected := helpers.TypeResponseSuccess{
					RequestID:  "123",
					StatusCode: http.StatusCreated,
					Data: map[string]any{
						"created_at": "23/09/2021 10:00:00",
						"updated_at": "23/09/2021 10:00:00",
						"id":         "123",
						"group_name": "test 1",
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
				Expect(groupCommandMock.FuncParamsInput).To(Equal([]any{group_dto.GroupInputDTO{
					GroupName: "teste 1",
					Code:      "23",
					Status:    "active",
				}}))
				Expect(groupCommandMock.FuncParamsOutput).To(Equal([]any{
					group_dto.GroupOutputDTO{
						ID:        "123",
						GroupName: "test 1",
						Code:      "23",
						Status:    "active",
						CreatedAt: "23/09/2021 10:00:00",
						UpdatedAt: "23/09/2021 10:00:00",
					},
					(*util.ValidationError)(nil),
				}))
			})
		})
	})
})
