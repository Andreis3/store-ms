//go:build unit
// +build unit

package group_service__test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/domain/service/group_service"
	repo_group "github.com/andreis3/stores-ms/internal/infra/repository/postgres/group"
	group_dto "github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
	"github.com/andreis3/stores-ms/internal/util"
	repo_group_mock "github.com/andreis3/stores-ms/tests/mock/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/tests/mock/infra/uow"
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
		var uow *uow_mock.UnitOfWorkMock
		var mapRegister []uow_mock.RegisterRepository
		var groupRepositoryMock *repo_group_mock.GroupRepositoryMock
		Context("When I call the method InsertGroup", func() {
			It("Should insert a new group not errors", func() {
				groupRepositoryMock = &repo_group_mock.GroupRepositoryMock{
					InsertGroupFunc: func(group repo_group.GroupModel) (string, *util.ValidationError) {
						return "1", nil
					},
				}

				mapRegister = []uow_mock.RegisterRepository{
					{
						Key:  util.GROUP_REPOSITORY_KEY,
						Repo: groupRepositoryMock,
					},
				}
				uow = uow_mock.NewProxyUnitOfWorkMock(mapRegister)
				service := group_service.NewInsertGroupService(uow)
				groupInputDTO := group_dto.GroupInputDTO{
					GroupName: "Group 1",
					Code:      "G1",
					Status:    "active",
				}

				groupOutputDTO, err := service.InsertGroup(groupInputDTO)
				Expect(err).To(BeNil())
				Expect(groupOutputDTO).ToNot(BeNil())

			})
		})
	})
})
