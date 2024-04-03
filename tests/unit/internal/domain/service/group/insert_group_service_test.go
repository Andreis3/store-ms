//go:build unit
// +build unit

package group_service__test

import (
	"context"
	iuow "github.com/andreis3/stores-ms/internal/infra/uow/interfaces"
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
		var unitOfWork *uow_mock.UnitOfWorkMock
		var mapRegister []uow_mock.RegisterRepository
		var groupRepositoryMock *repo_group_mock.GroupRepositoryMock
		Context("When I call the method InsertGroup", func() {
			It("Should insert a new group not errors", func() {
				mapRegister = make([]uow_mock.RegisterRepository, 0)
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
				unitOfWork = &uow_mock.UnitOfWorkMock{
					DoFunc: func(ctx context.Context, callback func(iuow.IUnitOfWork) *util.ValidationError) *util.ValidationError {
						return nil
					},
					RegisterFunc: func(name string, callback iuow.RepositoryFactory) {
						unitOfWork.RepositoryMocks[name] = callback
					},
					GetRepositoryFunc: func(name string) any {
						repo := unitOfWork.RepositoryMocks[name]
						return repo
					},
					CommitOrRollbackFunc: func() *util.ValidationError {
						return nil
					},
					RollbackFunc: func() *util.ValidationError {
						return nil
					},
				}
				uow := uow_mock.NewProxyUnitOfWorkMock(unitOfWork, mapRegister)
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
