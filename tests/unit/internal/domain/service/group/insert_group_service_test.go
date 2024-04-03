//go:build unit
// +build unit

package group_service__test

import (
	context_group_service_test "github.com/andreis3/stores-ms/tests/unit/internal/domain/service/group/context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/domain/service/group_service"
	group_dto "github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
)

func Test_GroupServiceInsertSuite(t *testing.T) {
	suiteConfig, reporterConfig := GinkgoConfiguration()

	suiteConfig.SkipStrings = []string{"SKIPPED", "PENDING", "NEVER-RUN", "SKIP"}
	reporterConfig.FullTrace = true
	reporterConfig.Succinct = true

	RegisterFailHandler(Fail)
	RunSpecs(t, "Group Service Insert Test Suite ", suiteConfig, reporterConfig)
}

var _ = Describe("DOMAIN :: SERVICE :: GROUP_SERVICE :: INSERT_GROUP_SERVICE", func() {
	Describe("#InsertGroup", func() {
		Context("When I call the method InsertGroup", func() {
			It("Should insert a new group not errors", func() {
				uow := context_group_service_test.ContextGroupServiceSuccess()
				service := group_service.NewInsertGroupService(uow)
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
		})
	})
})
