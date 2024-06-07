//go:build unit
// +build unit

package group_service_test

import (
	entity_group "github.com/andreis3/stores-ms/internal/domain/entity"
	"github.com/andreis3/stores-ms/internal/domain/services"
	"net/http"

	"github.com/stretchr/testify/mock"

	"github.com/andreis3/stores-ms/internal/domain/valueobject"
	"github.com/andreis3/stores-ms/internal/util"
	"github.com/andreis3/stores-ms/tests/mock/infra/common/uuid_mock"
	"github.com/andreis3/stores-ms/tests/mock/infra/repository/postgres/group"
	"github.com/andreis3/stores-ms/tests/mock/infra/uow"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DOMAIN :: SERVICE :: GROUP_SERVICE :: INSERT_GROUP_SERVICE", func() {
	Describe("#CreateGroup", func() {
		Context("When I call the method CreateGroup", func() {
			It("Should insert a new group not return errors", func() {
				groupRepositoryMock := new(repo_group_mock.GroupRepositoryMock)
				uuidMock := new(uuid_mock.UUIDMock)
				uowMock := ContextInsertSuccess(groupRepositoryMock, uuidMock)
				service := services.NewCreateGroupService(uowMock, uuidMock)
				status := valueobject.Status{
					Status: "active",
				}
				groupEntity := &entity_group.Group{
					Name:   "Group 1",
					Code:   "G1",
					Status: status,
				}

				groupOutputDTO, err := service.CreateGroup(*groupEntity)

				Expect(err).To(BeNil())
				Expect(groupOutputDTO).ToNot(BeNil())
				Expect(groupOutputDTO.Name).To(Equal(groupEntity.Name))
				Expect(groupOutputDTO.Code).To(Equal(groupEntity.Code))
				Expect(groupOutputDTO.Status).To(Equal(groupEntity.Status.Status))
				Expect(groupOutputDTO.ID).NotTo(BeEmpty())
				Expect(groupOutputDTO.CreatedAt).NotTo(BeEmpty())
				Expect(groupOutputDTO.UpdatedAt).NotTo(BeEmpty())
				Expect(groupRepositoryMock.ExpectedCalls).To(HaveLen(2))
				Expect(groupRepositoryMock.AssertCalled(GinkgoT(), repo_group_mock.SelectOneGroupByNameAndCode, groupEntity.Name, groupEntity.Code)).To(BeTrue())
				Expect(groupRepositoryMock.AssertNumberOfCalls(GinkgoT(), repo_group_mock.SelectOneGroupByNameAndCode, 1)).To(BeTrue())
				Expect(groupRepositoryMock.AssertNumberOfCalls(GinkgoT(), repo_group_mock.InsertGroup, 1)).To(BeTrue())
				Expect(groupRepositoryMock.AssertExpectations(GinkgoT())).To(Equal(true))
				Expect(uowMock.ExpectedCalls).To(HaveLen(2))
				Expect(uowMock.AssertCalled(GinkgoT(), uow_mock.GetRepository, util.GROUP_REPOSITORY_KEY)).To(BeTrue())
				Expect(uowMock.AssertCalled(GinkgoT(), uow_mock.Do, mock.AnythingOfType(uow_mock.DoParamFunc))).To(BeTrue())
				Expect(uowMock.AssertNumberOfCalls(GinkgoT(), uow_mock.Do, 1)).To(BeTrue())
				Expect(uowMock.AssertCalled(GinkgoT(), uow_mock.GetRepository, util.GROUP_REPOSITORY_KEY)).To(BeTrue())
				Expect(uowMock.AssertNumberOfCalls(GinkgoT(), uow_mock.GetRepository, 1)).To(BeTrue())
			})

			It("Should return an error when the method CreateGroup of the repository is call", func() {
				groupRepositoryMock := new(repo_group_mock.GroupRepositoryMock)
				uuidMock := new(uuid_mock.UUIDMock)
				uowMock := ContextInsertReturnErrorGroupRepositoryInsertGroup(groupRepositoryMock, uuidMock)
				service := services.NewCreateGroupService(uowMock, uuidMock)
				status := valueobject.Status{
					Status: "active",
				}
				groupEntity := &entity_group.Group{
					Name:   "Group 1",
					Code:   "G1",
					Status: status,
				}

				groupOutputDTO, err := service.CreateGroup(*groupEntity)
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
				uuidMock := new(uuid_mock.UUIDMock)
				uowMock := ContextInsertReturnErrorWhenCommitCommandUow(groupRepositoryMock, uuidMock)
				service := services.NewCreateGroupService(uowMock, uuidMock)
				status := valueobject.Status{
					Status: "active",
				}
				groupEntity := &entity_group.Group{
					Name:   "Group 1",
					Code:   "G1",
					Status: status,
				}

				groupOutputDTO, err := service.CreateGroup(*groupEntity)
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
				uuidMock := new(uuid_mock.UUIDMock)
				uowMock := ContextInsertSuccess(groupRepositoryMock, uuidMock)
				service := services.NewCreateGroupService(uowMock, uuidMock)
				status := valueobject.Status{
					Status: "active",
				}
				groupEntity := &entity_group.Group{
					Code:   "G1",
					Status: status,
				}

				groupOutputDTO, err := service.CreateGroup(*groupEntity)
				expectedError := &util.ValidationError{
					Code:        "VBR-0001",
					Origin:      "CreateGroupService.CreateGroup",
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
				uuidMock := new(uuid_mock.UUIDMock)
				uowMock := ContextInsertReturnErrorWhenSelectOneGroupByNameAndCode(groupRepositoryMock, uuidMock)
				service := services.NewCreateGroupService(uowMock, uuidMock)
				status := valueobject.Status{
					Status: "active",
				}
				groupEntity := &entity_group.Group{
					Name:   "Group 1",
					Code:   "G1",
					Status: status,
				}

				groupOutputDTO, err := service.CreateGroup(*groupEntity)
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
				uuidMock := new(uuid_mock.UUIDMock)
				uowMock := ContextInsertReturnErrorWhenSelectOneGroupByNameAndCodeReturnGroup(groupRepositoryMock, uuidMock)
				service := services.NewCreateGroupService(uowMock, uuidMock)
				status := valueobject.Status{
					Status: "active",
				}
				groupEntity := &entity_group.Group{
					Name:   "Group 1",
					Code:   "G1",
					Status: status,
				}

				groupOutputDTO, err := service.CreateGroup(*groupEntity)
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
