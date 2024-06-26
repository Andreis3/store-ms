//go:build unit
// +build unit

package group_service_test

import (
	"net/http"

	"github.com/andreis3/stores-ms/internal/util"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/domain/service/group"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
)

var _ = Describe("DOMAIN :: SERVICE :: GROUP_SERVICE :: INSERT_GROUP_SERVICE", func() {
	Describe("#InsertGroup", func() {
		Context("When I call the method InsertGroup", func() {
			It("Should insert a new group not return errors", func() {
				groupServiceUowDependency := ContextInsertSuccess()
				service := group_service.NewInsertGroupService(groupServiceUowDependency)
				groupInputDTO := group_dto.GroupInputDTO{
					GroupName: "Group 1",
					Code:      "G1",
					Status:    "active",
				}

				groupOutputDTO, err := service.InsertGroup(groupInputDTO)

				Expect(err).To(BeNil())
				Expect(groupOutputDTO).ToNot(BeNil())
				Expect(groupOutputDTO.GroupName).To(Equal(groupInputDTO.GroupName))
				Expect(groupOutputDTO.Code).To(Equal(groupInputDTO.Code))
				Expect(groupOutputDTO.Status).To(Equal(groupInputDTO.Status))
				Expect(groupOutputDTO.ID).NotTo(BeEmpty())
				Expect(groupOutputDTO.CreatedAt).NotTo(BeEmpty())
				Expect(groupOutputDTO.UpdatedAt).NotTo(BeEmpty())
			})

			It("Should return an error when the method InsertGroup of the repository is call", func() {
				groupServiceUowDependency := ContextInsertReturnErrorGroupRepositoryInsertGroup()
				service := group_service.NewInsertGroupService(groupServiceUowDependency)
				groupInputDTO := group_dto.GroupInputDTO{
					GroupName: "Group 1",
					Code:      "G1",
					Status:    "active",
				}

				groupOutputDTO, err := service.InsertGroup(groupInputDTO)
				expectedError := &util.ValidationError{
					Code:        "PIDB-235",
					Status:      500,
					ClientError: []string{"Internal Server Error"},
					LogError:    []string{"Insert group error"},
				}

				Expect(err).ToNot(BeNil())
				Expect(groupOutputDTO).To(BeZero())
				Expect(err).To(Equal(expectedError))
			})

			It("Should return an error when the method CommitOrRollback of the UOW is call", func() {
				groupServiceUowDependency := ContextInsertReturnErrorWhenCommitCommandUow()
				service := group_service.NewInsertGroupService(groupServiceUowDependency)
				groupInputDTO := group_dto.GroupInputDTO{
					GroupName: "Group 1",
					Code:      "G1",
					Status:    "active",
				}

				groupOutputDTO, err := service.InsertGroup(groupInputDTO)
				expectedError := &util.ValidationError{
					Code:        "PIDB-235",
					Status:      http.StatusInternalServerError,
					LogError:    []string{"Commit error"},
					ClientError: []string{"Internal Server Error"},
				}

				Expect(err).ToNot(BeNil())
				Expect(groupOutputDTO).To(BeZero())
				Expect(err).To(Equal(expectedError))
			})

			It("Should return an error when the payload input is invalid", func() {
				groupServiceUowDependency := ContextInsertSuccess()
				service := group_service.NewInsertGroupService(groupServiceUowDependency)
				groupInputDTO := group_dto.GroupInputDTO{
					GroupName: "",
					Code:      "G1",
					Status:    "active",
				}

				groupOutputDTO, err := service.InsertGroup(groupInputDTO)
				expectedError := &util.ValidationError{
					Code:        "VBR-0001",
					Status:      http.StatusBadRequest,
					ClientError: []string{"group_name: is required"},
					LogError:    []string{"group_name: is required"},
				}

				Expect(err).ToNot(BeNil())
				Expect(groupOutputDTO).To(BeZero())
				Expect(err).To(Equal(expectedError))
			})
		})
	})
})
