//go:build unit
// +build unit

package group_command_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/app/command/group"
	"github.com/andreis3/stores-ms/internal/util"
)

var _ = Describe("APP :: COMMAND :: GROUP :: SEARCH_GROUP_COMMAND", func() {
	Describe("#Execute", func() {
		Context("When I call the method SearchOneGroup of the search_group_service", func() {
			It("Should search a new group not return errors", func() {
				searchServiceMock := ContextSearchSuccess()
				command := group_command.NewSearchGroupCommand(searchServiceMock)

				id := "1"

				groupOutputDTO, err := command.Execute(id)

				Expect(err).To(BeNil())
				Expect(groupOutputDTO).ToNot(BeNil())
				Expect(groupOutputDTO.Name).To(Equal("Group 1"))
				Expect(groupOutputDTO.Code).To(Equal("G1"))
				Expect(groupOutputDTO.Status).To(Equal("active"))
				Expect(groupOutputDTO.ID).NotTo(BeEmpty())
				Expect(groupOutputDTO.CreatedAt).NotTo(BeEmpty())
				Expect(groupOutputDTO.UpdatedAt).NotTo(BeEmpty())
			})

			It("Should return an error when the method SearchOneGroup of the search_group_service is call", func() {
				searchServiceMock := ContextSearchReturnErrorGroupServiceInsertGroup()
				command := group_command.NewSearchGroupCommand(searchServiceMock)

				id := "1"

				groupOutputDTO, err := command.Execute(id)
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
