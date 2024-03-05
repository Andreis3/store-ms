//go:build unit
// +build unit

package store_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/domain/entity/store"
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
				store := store.NewStore("", "Company Name", "12345678901234", "domain.com", "groupCOD", status, []store.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications).To(HaveLen(1))
				Expect(notifications[0]).To(HaveKeyWithValue("store_key", "is required"))
			})

			It("Should return a notifications when CompanyName is empty", func() {
				status := valueobject.NewStatus("active")
				store := store.NewStore("storeKey", "", "12345678901234", "domain.com", "groupCOD", status, []store.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications).To(HaveLen(1))
				Expect(notifications[0]).To(HaveKeyWithValue("company_name", "is required"))
			})

			It("Should return a notifications when Status is empty", func() {
				status := valueobject.NewStatus("")
				store := store.NewStore("storeKey", "Company Name", "12345678901234", "domain.com", "groupCOD", status, []store.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications).To(HaveLen(1))
				Expect(notifications[0]).To(HaveKeyWithValue("status", "is required"))
			})

			It("Should return a notifications when Status is invalid", func() {
				status := valueobject.NewStatus("invalid")
				store := store.NewStore("storeKey", "Company Name", "12345678901234", "domain.com", "groupCOD", status, []store.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications).To(HaveLen(1))
				Expect(notifications[0]).To(HaveKeyWithValue("status", "is invalid, valid values are active or inactive"))
			})

			It("Should return a notifications when CNPJ is empty", func() {
				status := valueobject.NewStatus("active")
				store := store.NewStore("storeKey", "Company Name", "", "domain.com", "groupCOD", status, []store.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications).To(HaveLen(1))
				Expect(notifications[0]).To(HaveKeyWithValue("cnpj", "is required"))
			})

			It("Should return a notifications when Domain is empty", func() {
				status := valueobject.NewStatus("active")
				store := store.NewStore("storeKey", "Company Name", "12345678901234", "", "groupCOD", status, []store.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications).To(HaveLen(1))
				Expect(notifications[0]).To(HaveKeyWithValue("domain", "is required"))
			})

			It("Should return a notifications when GroupCOD is empty", func() {
				status := valueobject.NewStatus("active")
				store := store.NewStore("storeKey", "Company Name", "12345678901234", "domain.com", "", status, []store.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications).To(HaveLen(1))
				Expect(notifications[0]).To(HaveKeyWithValue("group_code", "is required"))
			})

			It("Should return a notifications when Contact is empty", func() {
				status := valueobject.NewStatus("active")
				store := store.NewStore("storeKey", "Company Name", "12345678901234", "domain.com", "groupCOD", status, []store.Contact{})

				notifications := store.Validate()

				Expect(notifications).To(HaveLen(1))
				Expect(notifications[0]).To(HaveKeyWithValue("contacts", "min 1 contact is required"))
			})

			It("Should return a notifications when Contact.Name is empty", func() {
				status := valueobject.NewStatus("active")
				store := store.NewStore("storeKey", "Company Name", "12345678901234", "domain.com", "groupCOD", status, []store.Contact{
					{
						Name:  "",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications).To(HaveLen(1))
				Expect(notifications[0]).To(HaveKeyWithValue("contacts[0].name", "is required"))
			})

			It("Should return a notifications when Contact.Email is empty", func() {
				status := valueobject.NewStatus("active")
				store := store.NewStore("storeKey", "Company Name", "12345678901234", "domain.com", "groupCOD", status, []store.Contact{
					{
						Name:  "Contact Name",
						Email: "",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications).To(HaveLen(1))
				Expect(notifications).To(ContainElement(map[string]any{"contacts[0].email": "is required"}))
			})

			It("Should return a notifications when Contact.Phone is empty", func() {
				status := valueobject.NewStatus("active")
				store := store.NewStore("storeKey", "Company Name", "12345678901234", "domain.com", "groupCOD", status, []store.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications).To(HaveLen(1))
				Expect(notifications).To(ContainElement(map[string]any{"contacts[0].phone": "is required"}))
			})

			It("Should return a notifications when Contact contains 2 elements with elements 1 empty name and 2 empty email", func() {
				status := valueobject.NewStatus("active")
				store := store.NewStore("storeKey", "Company Name", "12345678901234", "domain.com", "groupCOD", status, []store.Contact{
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

				Expect(notifications).To(HaveLen(2))
				Expect(notifications).To(ContainElement(map[string]any{"contacts[0].name": "is required"}))
				Expect(notifications).To(ContainElement(map[string]any{"contacts[1].email": "is required"}))
			})

			It("Should return a notifications empty when all fields are filled and status active", func() {
				status := valueobject.NewStatus("active")
				store := store.NewStore("storeKey", "Company Name", "12345678901234", "domain.com", "groupCOD", status, []store.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications).To(BeEmpty())
			})

			It("Should return a notifications empty when all fields are filled and status inactive", func() {
				status := valueobject.NewStatus("inactive")
				store := store.NewStore("storeKey", "Company Name", "12345678901234", "domain.com", "groupCOD", status, []store.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications).To(BeEmpty())
			})
		})
	})
})
