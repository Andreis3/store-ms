//go:build unit
// +build unit

package group_command_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/app/command/group"
	"github.com/andreis3/stores-ms/internal/interfaces/http/controllers/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
)

var _ = Describe("APP :: COMMAND :: GROUP :: CREATE_GROUP_COMMAND", func() {
	Describe("#Execute", func() {
		Context("When I call the method CreateGroup of the create_group_service", func() {
			It("Should insert a new group not return errors", func() {
				insertGroupService := ContextCreateSuccess()
				command := group_command.NewCreateGroupCommand(insertGroupService)

				groupInputDTO := group_dto.GroupInputDTO{
					Name:   "Group 1",
					Code:   "G1",
					Status: "active",
				}

				groupOutputDTO, err := command.Execute(groupInputDTO)

				Expect(err).To(BeNil())
				Expect(groupOutputDTO).ToNot(BeNil())
				Expect(groupOutputDTO.Name).To(Equal(groupInputDTO.Name))
				Expect(groupOutputDTO.Code).To(Equal(groupInputDTO.Code))
				Expect(groupOutputDTO.Status).To(Equal(groupInputDTO.Status))
				Expect(groupOutputDTO.ID).NotTo(BeEmpty())
				Expect(groupOutputDTO.CreatedAt).NotTo(BeEmpty())
				Expect(groupOutputDTO.UpdatedAt).NotTo(BeEmpty())
			})

			It("Should return an error when the method CreateGroup of the create_group_service is call", func() {
				insertGroupService := ContextCreateReturnErrorGroupServiceCreateGroup()
				command := group_command.NewCreateGroupCommand(insertGroupService)

				groupInputDTO := group_dto.GroupInputDTO{
					Name:   "Group 1",
					Code:   "G1",
					Status: "active",
				}

				groupOutputDTO, err := command.Execute(groupInputDTO)
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