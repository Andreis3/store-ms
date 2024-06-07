//go:build unit
// +build unit

package store_test

import (
	"github.com/andreis3/stores-ms/internal/domain/entity"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/domain/valueobject"
)

func Test_StoreEntitySuite(t *testing.T) {
	suiteConfig, reporterConfig := GinkgoConfiguration()

	suiteConfig.SkipStrings = []string{"SKIPPED", "PENDING", "NEVER-RUN", "SKIP"}
	reporterConfig.FullTrace = true
	reporterConfig.Succinct = true

	RegisterFailHandler(Fail)
	RunSpecs(t, "Store Entity Test Suite ", suiteConfig, reporterConfig)
}

var _ = Describe("DOMAIN :: ENTITY :: STORE", func() {
	Describe("#Validate", func() {
		Context("When I call the method Validate", func() {
			It("Should return a notifications when StoreKey is empty", func() {
				status := valueobject.NewStatus("active")
				cnpj := valueobject.NewCNPJ("10.140.120/0001-48")
				store := entity.NewStore("", "Company Name", "domain.com", "groupCOD", cnpj, status, []entity.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(HaveLen(1))
				Expect(notifications.ReturnNotification()[0]).To(Equal("store_key: is required"))
			})

			It("Should return a notifications when CompanyName is empty", func() {
				status := valueobject.NewStatus("active")
				cnpj := valueobject.NewCNPJ("10.140.120/0001-48")
				store := entity.NewStore("storeKey", "", "domain.com", "groupCOD", cnpj, status, []entity.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(HaveLen(1))
				Expect(notifications.ReturnNotification()[0]).To(Equal("company_name: is required"))
			})

			It("Should return a notifications when Status is empty", func() {
				status := valueobject.NewStatus("")
				cnpj := valueobject.NewCNPJ("10.140.120/0001-48")
				store := entity.NewStore("storeKey", "Company Name", "domain.com", "groupCOD", cnpj, status, []entity.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(HaveLen(1))
				Expect(notifications.ReturnNotification()[0]).To(Equal("status: is required"))
			})

			It("Should return a notifications when Status is invalid", func() {
				status := valueobject.NewStatus("invalid")
				cnpj := valueobject.NewCNPJ("10.140.120/0001-48")
				store := entity.NewStore("storeKey", "Company Name", "domain.com", "groupCOD", cnpj, status, []entity.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(HaveLen(1))
				Expect(notifications.ReturnNotification()[0]).To(Equal("status: is invalid, valid values are active or inactive"))
			})

			It("Should return a notifications when CNPJ is empty", func() {
				status := valueobject.NewStatus("active")
				cnpj := valueobject.NewCNPJ("")
				store := entity.NewStore("storeKey", "Company Name", "domain.com", "groupCOD", cnpj, status, []entity.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(HaveLen(1))
				Expect(notifications.ReturnNotification()[0]).To(Equal("cnpj: is required"))
			})

			It("Should return a notifications when CNPJ with less than 14 characters", func() {
				status := valueobject.NewStatus("active")
				cnpj := valueobject.NewCNPJ("10.140.120/0001-4")
				store := entity.NewStore("storeKey", "Company Name", "domain.com", "groupCOD", cnpj, status, []entity.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(HaveLen(1))
				Expect(notifications.ReturnNotification()[0]).To(Equal("cnpj: is invalid, must have 14 characters"))
			})

			It("Should return a notifications when CNPJ is in the black list", func() {
				status := valueobject.NewStatus("active")
				cnpj := valueobject.NewCNPJ("00.000.000/0000-00")
				store := entity.NewStore("storeKey", "Company Name", "domain.com", "groupCOD", cnpj, status, []entity.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(HaveLen(1))
				Expect(notifications.ReturnNotification()[0]).To(Equal("cnpj: is invalid, must be a valid CNPJ number"))
			})

			It("Should return a notifications when CNPJ is invalid", func() {
				status := valueobject.NewStatus("active")
				cnpj := valueobject.NewCNPJ("10.140.120/0001-49")
				store := entity.NewStore("storeKey", "Company Name", "domain.com", "groupCOD", cnpj, status, []entity.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(HaveLen(1))
				Expect(notifications.ReturnNotification()[0]).To(Equal("cnpj: is invalid, must be a valid CNPJ number calculated with the module 11 algorithm"))
			})

			It("Should return a notifications when Domain is empty", func() {
				status := valueobject.NewStatus("active")
				cnpj := valueobject.NewCNPJ("10.140.120/0001-48")
				store := entity.NewStore("storeKey", "Company Name", "", "groupCOD", cnpj, status, []entity.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(HaveLen(1))
				Expect(notifications.ReturnNotification()[0]).To(Equal("domain: is required"))
			})

			It("Should return a notifications when GroupCOD is empty", func() {
				status := valueobject.NewStatus("active")
				cnpj := valueobject.NewCNPJ("10.140.120/0001-48")
				store := entity.NewStore("storeKey", "Company Name", "domain.com", "", cnpj, status, []entity.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(HaveLen(1))
				Expect(notifications.ReturnNotification()[0]).To(Equal("code: is required"))
			})

			It("Should return a notifications when Contact is empty", func() {
				status := valueobject.NewStatus("active")
				cnpj := valueobject.NewCNPJ("10.140.120/0001-48")
				store := entity.NewStore("storeKey", "Company Name", "domain.com", "groupCOD", cnpj, status, []entity.Contact{})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(HaveLen(1))
				Expect(notifications.ReturnNotification()[0]).To(Equal("contacts: min 1 contact is required"))
			})

			It("Should return a notifications when Contact.Name is empty", func() {
				status := valueobject.NewStatus("active")
				cnpj := valueobject.NewCNPJ("10.140.120/0001-48")
				store := entity.NewStore("storeKey", "Company Name", "domain.com", "groupCOD", cnpj, status, []entity.Contact{
					{
						Name:  "",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(HaveLen(1))
				Expect(notifications.ReturnNotification()[0]).To(Equal("contacts[0].name: is required"))
			})

			It("Should return a notifications when Contact.Email is empty", func() {
				status := valueobject.NewStatus("active")
				cnpj := valueobject.NewCNPJ("10.140.120/0001-48")
				store := entity.NewStore("storeKey", "Company Name", "domain.com", "groupCOD", cnpj, status, []entity.Contact{
					{
						Name:  "Contact Name",
						Email: "",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(HaveLen(1))
				Expect(notifications.ReturnNotification()).To(ContainElement("contacts[0].email: is required"))
			})

			It("Should return a notifications when Contact.Phone is empty", func() {
				status := valueobject.NewStatus("active")
				cnpj := valueobject.NewCNPJ("10.140.120/0001-48")
				store := entity.NewStore("storeKey", "Company Name", "domain.com", "groupCOD", cnpj, status, []entity.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(HaveLen(1))
				Expect(notifications.ReturnNotification()).To(ContainElement("contacts[0].phone: is required"))
			})

			It("Should return a notifications when Contact contains 2 elements with elements 1 empty name and 2 empty email", func() {
				status := valueobject.NewStatus("active")
				cnpj := valueobject.NewCNPJ("10.140.120/0001-48")
				store := entity.NewStore("storeKey", "Company Name", "domain.com", "groupCOD", cnpj, status, []entity.Contact{
					{
						Name:  "",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
					{
						Name:  "Contact Name",
						Email: "",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(HaveLen(2))
				Expect(notifications.ReturnNotification()).To(ContainElement("contacts[0].name: is required"))
				Expect(notifications.ReturnNotification()).To(ContainElement("contacts[1].email: is required"))
			})

			It("Should return a notifications empty when all fields are filled and status active", func() {
				status := valueobject.NewStatus("active")
				cnpj := valueobject.NewCNPJ("10.140.120/0001-48")
				store := entity.NewStore("storeKey", "Company Name", "domain.com", "groupCOD", cnpj, status, []entity.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(BeEmpty())
			})

			It("Should return a notifications empty when all fields are filled and status inactive", func() {
				status := valueobject.NewStatus("inactive")
				cnpj := valueobject.NewCNPJ("10.140.120/0001-48")
				store := entity.NewStore("storeKey", "Company Name", "domain.com", "groupCOD", cnpj, status, []entity.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications.ReturnNotification()).To(BeEmpty())
			})
		})
	})
})
