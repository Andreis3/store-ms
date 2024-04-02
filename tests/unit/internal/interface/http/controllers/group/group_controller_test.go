//go:build unit
// +build unit

package group_controller_test

import (
	"encoding/json"
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

				groupController = group_controller.NewGroupController(groupCommandMock, loggerMock, requestIDMock)
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
				loggerMock = &logger_mock.LoggerMock{}

				groupController = group_controller.NewGroupController(groupCommandMock, loggerMock, requestIDMock)
				body := `{
    						"group_name":"teste 1",
							"code": "23",
    						"status":"active"
						  }`
				request, err := http.NewRequest("POST", "/group", strings.NewReader(body))
				writer := httptest.NewRecorder()

				expected := helpers.TypeResponseSuccess{
					RequestID:  "123",
					StatusCode: 201,
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
			})
		})
	})
})
