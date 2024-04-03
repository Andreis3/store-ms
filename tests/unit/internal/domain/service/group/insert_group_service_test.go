//go:build unit
// +build unit

package group_service_test

import (
	"github.com/andreis3/stores-ms/internal/util"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/domain/service/group_service"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
)

var _ = Describe("DOMAIN :: SERVICE :: GROUP_SERVICE :: INSERT_GROUP_SERVICE", func() {
	Describe("#InsertGroup", func() {
		Context("When I call the method InsertGroup", func() {
			It("Should insert a new group not errors", func() {
				groupServiceUowDependency := Success()
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
				groupServiceUowDependency := ReturnErrorGroupRepositoryInsertGroup()
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
		})
	})
})
