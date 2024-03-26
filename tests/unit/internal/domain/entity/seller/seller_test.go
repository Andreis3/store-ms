//go:build unit
// +build unit

package seller_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/domain/entity/seller"
	"github.com/andreis3/stores-ms/internal/domain/valueobject"
)

func Test_SellerEntitySuite(t *testing.T) {
	suiteConfig, reporterConfig := GinkgoConfiguration()

	suiteConfig.SkipStrings = []string{"SKIPPED", "PENDING", "NEVER-RUN", "SKIP"}
	reporterConfig.FullTrace = true
	reporterConfig.Succinct = true

	RegisterFailHandler(Fail)
	RunSpecs(t, "Seller Entity Test Suite ", suiteConfig, reporterConfig)
}

var _ = Describe("DOMAIN :: ENTITY :: SELLER", func() {
	Describe("#Validate", func() {
		Context("When I call the method Validate", func() {
			It("Should return a notification when seller name is empty", func() {
				status := valueobject.NewStatus("active")
				seller := entity_seller.NewSeller("", "123", status, true)

				notification := seller.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]).To(Equal("seller_name: is required"))
			})

			It("Should return a notification when code is empty", func() {
				status := valueobject.NewStatus("active")
				seller := entity_seller.NewSeller("seller", "", status, true)

				notification := seller.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]).To(Equal("code: is required"))
			})

			It("Should return a notification when status is empty", func() {
				status := valueobject.NewStatus("")
				seller := entity_seller.NewSeller("seller", "123", status, true)

				notification := seller.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]).To(Equal("status: is required"))
			})

			It("Should return a notification when seller name, code and status are empty", func() {
				status := valueobject.NewStatus("")
				seller := entity_seller.NewSeller("", "", status, true)

				notification := seller.Validate()

				Expect(notification).To(HaveLen(3))
				Expect(notification[0]).To(Equal("seller_name: is required"))
				Expect(notification[1]).To(Equal("code: is required"))
				Expect(notification[2]).To(Equal("status: is required"))
			})

			It("Should return a notification when status is invalid", func() {
				status := valueobject.NewStatus("invalid")
				seller := entity_seller.NewSeller("seller", "123", status, true)

				notification := seller.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]).To(Equal("status: is invalid, valid values are active or inactive"))
			})

			It("Should not return a notification when status is active", func() {
				status := valueobject.NewStatus("active")
				seller := entity_seller.NewSeller("seller", "123", status, true)

				notification := seller.Validate()

				Expect(notification).To(HaveLen(0))
			})

			It("Should not return a notification when status is inactive", func() {
				status := valueobject.NewStatus("inactive")
				seller := entity_seller.NewSeller("seller", "123", status, true)

				notification := seller.Validate()

				Expect(notification).To(HaveLen(0))
			})
		})
	})
})
