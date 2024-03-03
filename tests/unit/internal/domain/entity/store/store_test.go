package store_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/domain/entity/store"
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
				store := store.NewStore("", "Company Name", "active", "12345678901234", "domain.com", "groupCOD", []store.Contact{
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
				store := store.NewStore("storeKey", "", "active", "12345678901234", "domain.com", "groupCOD", []store.Contact{
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
				store := store.NewStore("storeKey", "Company Name", "", "12345678901234", "domain.com", "groupCOD", []store.Contact{
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
				store := store.NewStore("storeKey", "Company Name", "invalid", "12345678901234", "domain.com", "groupCOD", []store.Contact{
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
				store := store.NewStore("storeKey", "Company Name", "active", "", "domain.com", "groupCOD", []store.Contact{
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
				store := store.NewStore("storeKey", "Company Name", "active", "12345678901234", "", "groupCOD", []store.Contact{
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
				store := store.NewStore("storeKey", "Company Name", "active", "12345678901234", "domain.com", "", []store.Contact{
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
				store := store.NewStore("storeKey", "Company Name", "active", "12345678901234", "domain.com", "groupCOD", []store.Contact{})

				notifications := store.Validate()

				Expect(notifications).To(HaveLen(1))
				Expect(notifications[0]).To(HaveKeyWithValue("contacts", "min 1 contact is required"))
			})

			It("Should return a notifications when Contact.Name is empty", func() {
				store := store.NewStore("storeKey", "Company Name", "active", "12345678901234", "domain.com", "groupCOD", []store.Contact{
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
				store := store.NewStore("storeKey", "Company Name", "active", "12345678901234", "domain.com", "groupCOD", []store.Contact{
					{
						Name:  "Contact Name",
						Email: "",
						Phone: "1234567890",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications).To(HaveLen(1))
				Expect(notifications).To(ContainElement(map[string]string{"contacts[0].email": "is required"}))
			})

			It("Should return a notifications when Contact.Phone is empty", func() {
				store := store.NewStore("storeKey", "Company Name", "active", "12345678901234", "domain.com", "groupCOD", []store.Contact{
					{
						Name:  "Contact Name",
						Email: "email@.com.br",
						Phone: "",
						Ramal: "123",
					},
				})

				notifications := store.Validate()

				Expect(notifications).To(HaveLen(1))
				Expect(notifications).To(ContainElement(map[string]string{"contacts[0].phone": "is required"}))
			})

			It("Should return a notifications when Contact contains 2 elements with elements 1 empty name and 2 empty email", func() {
				store := store.NewStore("storeKey", "Company Name", "active", "12345678901234", "domain.com", "groupCOD", []store.Contact{
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
				Expect(notifications).To(ContainElement(map[string]string{"contacts[0].name": "is required"}))
				Expect(notifications).To(ContainElement(map[string]string{"contacts[1].email": "is required"}))
			})

			It("Should return a notifications empty when all fields are filled and status active", func() {
				store := store.NewStore("storeKey", "Company Name", "active", "12345678901234", "domain.com", "groupCOD", []store.Contact{
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
				store := store.NewStore("storeKey", "Company Name", "inactive", "12345678901234", "domain.com", "groupCOD", []store.Contact{
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
