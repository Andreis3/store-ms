package saleschannel_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/domain/entity/saleschannel"
)

func Test_SalesChannelEntitySuite(t *testing.T) {
	suiteConfig, reporterConfig := GinkgoConfiguration()

	suiteConfig.SkipStrings = []string{"SKIPPED", "PENDING", "NEVER-RUN", "SKIP"}
	reporterConfig.FullTrace = true
	reporterConfig.Succinct = true

	RegisterFailHandler(Fail)
	RunSpecs(t, "SalesChannel Entity Test Suite ", suiteConfig, reporterConfig)
}

var _ = Describe("DOMAIN :: ENTITY :: SALES_CHANNEL", func() {
	Describe("#Validate", func() {
		Context("When I call the method Validate", func() {
			It("Should return a notification when sales channel is empty", func() {
				salesChannel := saleschannel.NewSalesChannel("", "123", "active", true)

				notification := salesChannel.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["sales_channel"]).To(Equal("is required"))
			})

			It("Should return a notification when code is empty", func() {
				salesChannel := saleschannel.NewSalesChannel("salesChannel", "", "active", true)

				notification := salesChannel.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["code"]).To(Equal("is required"))
			})

			It("Should return a notification when status is empty", func() {
				salesChannel := saleschannel.NewSalesChannel("salesChannel", "123", "", true)

				notification := salesChannel.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["status"]).To(Equal("is required"))
			})

			It("Should return a notification when sales channel, code and status are empty", func() {
				salesChannel := saleschannel.NewSalesChannel("", "", "", true)

				notification := salesChannel.Validate()

				Expect(notification).To(HaveLen(3))
				Expect(notification[0]["sales_channel"]).To(Equal("is required"))
			})

			It("Should return a notification when status is invalid", func() {
				salesChannel := saleschannel.NewSalesChannel("salesChannel", "123", "invalid", true)

				notification := salesChannel.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["status"]).To(Equal("is invalid, valid values are active or inactive"))
			})

			It("Should not return a notification when status is active", func() {
				salesChannel := saleschannel.NewSalesChannel("salesChannel", "123", "active", true)

				notification := salesChannel.Validate()

				Expect(notification).To(HaveLen(0))
			})

			It("Should not return a notification when status is inactive", func() {
				salesChannel := saleschannel.NewSalesChannel("salesChannel", "123", "inactive", true)

				notification := salesChannel.Validate()

				Expect(notification).To(HaveLen(0))
			})
		})
	})
})
