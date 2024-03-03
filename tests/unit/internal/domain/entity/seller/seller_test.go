//go:build unit
// +build unit

package seller_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/domain/entity/seller"
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
				seller := seller.NewSeller("", "123", "active", true)

				notification := seller.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["seller_name"]).To(Equal("is required"))
			})

			It("Should return a notification when code is empty", func() {
				seller := seller.NewSeller("seller", "", "active", true)

				notification := seller.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["code"]).To(Equal("is required"))
			})

			It("Should return a notification when status is empty", func() {
				seller := seller.NewSeller("seller", "123", "", true)

				notification := seller.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["status"]).To(Equal("is required"))
			})

			It("Should return a notification when seller name, code and status are empty", func() {
				seller := seller.NewSeller("", "", "", true)

				notification := seller.Validate()

				Expect(notification).To(HaveLen(3))
				Expect(notification[0]["seller_name"]).To(Equal("is required"))
				Expect(notification[1]["code"]).To(Equal("is required"))
				Expect(notification[2]["status"]).To(Equal("is required"))
			})

			It("Should return a notification when status is invalid", func() {
				seller := seller.NewSeller("seller", "123", "invalid", true)

				notification := seller.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["status"]).To(Equal("is invalid, valid values are active or inactive"))
			})

			It("Should not return a notification when status is active", func() {
				seller := seller.NewSeller("seller", "123", "active", true)

				notification := seller.Validate()

				Expect(notification).To(HaveLen(0))
			})

			It("Should not return a notification when status is inactive", func() {
				seller := seller.NewSeller("seller", "123", "inactive", true)

				notification := seller.Validate()

				Expect(notification).To(HaveLen(0))
			})
		})
	})
})
