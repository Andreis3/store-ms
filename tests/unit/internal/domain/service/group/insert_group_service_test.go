//go:build unit
// +build unit

package group_service_test

import (
	"net/http"

	"github.com/stretchr/testify/mock"

	"github.com/andreis3/stores-ms/internal/util"
	"github.com/andreis3/stores-ms/tests/mock/infra/repository/postgres/group"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/domain/service/group"
	"github.com/andreis3/stores-ms/internal/interface/http/controller/group/dto"
)

var _ = Describe("DOMAIN :: SERVICE :: GROUP_SERVICE :: INSERT_GROUP_SERVICE", func() {
	Describe("#InsertGroup", func() {
		Context("When I call the method InsertGroup", func() {
			It("Should insert a new group not return errors", func() {
				groupRepositoryMock := new(repo_group_mock.GroupRepositoryMock)
				groupServiceUowDependency := ContextInsertSuccess(groupRepositoryMock)
				service := group_service.NewInsertGroupService(groupServiceUowDependency)
				groupInputDTO := group_dto.GroupInputDTO{
					Name:   "Group 1",
					Code:   "G1",
					Status: "active",
				}

				groupOutputDTO, err := service.InsertGroup(groupInputDTO)

				Expect(err).To(BeNil())
				Expect(groupOutputDTO).ToNot(BeNil())
				Expect(groupOutputDTO.Name).To(Equal(groupInputDTO.Name))
				Expect(groupOutputDTO.Code).To(Equal(groupInputDTO.Code))
				Expect(groupOutputDTO.Status).To(Equal(groupInputDTO.Status))
				Expect(groupOutputDTO.ID).NotTo(BeEmpty())
				Expect(groupOutputDTO.CreatedAt).NotTo(BeEmpty())
				Expect(groupOutputDTO.UpdatedAt).NotTo(BeEmpty())
				Expect(groupRepositoryMock.ExpectedCalls).To(HaveLen(2))
				Expect(groupRepositoryMock.AssertCalled(GinkgoT(), "SelectOneGroupByNameAndCode", groupInputDTO.Name, groupInputDTO.Code)).To(BeTrue())
				Expect(groupRepositoryMock.AssertNumberOfCalls(GinkgoT(), "SelectOneGroupByNameAndCode", 1)).To(BeTrue())
				Expect(groupRepositoryMock.AssertNumberOfCalls(GinkgoT(), "InsertGroup", 1)).To(BeTrue())
				Expect(groupRepositoryMock.AssertExpectations(GinkgoT())).To(Equal(true))
				Expect(groupServiceUowDependency.ExpectedCalls).To(HaveLen(2))
				Expect(groupServiceUowDependency.AssertCalled(GinkgoT(), "Do", mock.AnythingOfType("func(iuow.IUnitOfWork) *util.ValidationError"))).To(BeTrue())
				Expect(groupServiceUowDependency.AssertNumberOfCalls(GinkgoT(), "Do", 1)).To(BeTrue())
				Expect(groupServiceUowDependency.AssertCalled(GinkgoT(), "GetRepository", util.GROUP_REPOSITORY_KEY)).To(BeTrue())
				Expect(groupServiceUowDependency.AssertNumberOfCalls(GinkgoT(), "GetRepository", 1)).To(BeTrue())
			})

			It("Should return an error when the method InsertGroup of the repository is call", func() {
				groupRepositoryMock := new(repo_group_mock.GroupRepositoryMock)
				groupServiceUowDependency := ContextInsertReturnErrorGroupRepositoryInsertGroup(groupRepositoryMock)
				service := group_service.NewInsertGroupService(groupServiceUowDependency)
				groupInputDTO := group_dto.GroupInputDTO{
					Name:   "Group 1",
					Code:   "G1",
					Status: "active",
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
				groupRepositoryMock := new(repo_group_mock.GroupRepositoryMock)
				groupServiceUowDependency := ContextInsertReturnErrorWhenCommitCommandUow(groupRepositoryMock)
				service := group_service.NewInsertGroupService(groupServiceUowDependency)
				groupInputDTO := group_dto.GroupInputDTO{
					Name:   "Group 1",
					Code:   "G1",
					Status: "active",
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
				groupRepositoryMock := new(repo_group_mock.GroupRepositoryMock)
				groupServiceUowDependency := ContextInsertSuccess(groupRepositoryMock)
				service := group_service.NewInsertGroupService(groupServiceUowDependency)
				groupInputDTO := group_dto.GroupInputDTO{
					Name:   "",
					Code:   "G1",
					Status: "active",
				}

				groupOutputDTO, err := service.InsertGroup(groupInputDTO)
				expectedError := &util.ValidationError{
					Code:        "VBR-0001",
					Status:      http.StatusBadRequest,
					ClientError: []string{"name: is required"},
					LogError:    []string{"name: is required"},
				}

				Expect(err).ToNot(BeNil())
				Expect(groupOutputDTO).To(BeZero())
				Expect(err).To(Equal(expectedError))
			})

			It("Should return an error when the select group by name and code return an error", func() {
				groupRepositoryMock := new(repo_group_mock.GroupRepositoryMock)
				groupServiceUowDependency := ContextInsertReturnErrorWhenSelectOneGroupByNameAndCode(groupRepositoryMock)
				service := group_service.NewInsertGroupService(groupServiceUowDependency)
				groupInputDTO := group_dto.GroupInputDTO{
					Name:   "Group 1",
					Code:   "G1",
					Status: "active",
				}

				groupOutputDTO, err := service.InsertGroup(groupInputDTO)
				expectedError := &util.ValidationError{
					Code:        "PIDB-235",
					Status:      http.StatusInternalServerError,
					LogError:    []string{"Select group error"},
					ClientError: []string{"Internal Server Error"},
				}

				Expect(err).ToNot(BeNil())
				Expect(groupOutputDTO).To(BeZero())
				Expect(err).To(Equal(expectedError))
			})

			It("Should return an error when the group already exists", func() {
				groupRepositoryMock := new(repo_group_mock.GroupRepositoryMock)
				groupServiceUowDependency := ContextInsertReturnErrorWhenSelectOneGroupByNameAndCodeReturnGroup(groupRepositoryMock)
				service := group_service.NewInsertGroupService(groupServiceUowDependency)
				groupInputDTO := group_dto.GroupInputDTO{
					Name:   "Group 1",
					Code:   "G1",
					Status: "active",
				}

				groupOutputDTO, err := service.InsertGroup(groupInputDTO)
				expectedError := &util.ValidationError{
					Code:        "VBR-0002",
					Status:      http.StatusBadRequest,
					ClientError: []string{"Group already exists"},
					LogError:    []string{"Group already exists"},
				}

				Expect(err).ToNot(BeNil())
				Expect(groupOutputDTO).To(BeZero())
				Expect(err).To(Equal(expectedError))
			})
		})
	})
})
