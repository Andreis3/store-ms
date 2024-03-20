//go:build unit
// +build unit

package saleschannel_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/domain/entity/saleschannel"
	"github.com/andreis3/stores-ms/internal/domain/valueobject"
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
				status := valueobject.NewStatus("active")
				salesChannel := entity_saleschannel.NewSalesChannel("", "123", status, true)

				notification := salesChannel.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["sales_channel"]).To(Equal("is required"))
			})

			It("Should return a notification when code is empty", func() {
				status := valueobject.NewStatus("active")
				salesChannel := entity_saleschannel.NewSalesChannel("salesChannel", "", status, true)

				notification := salesChannel.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["code"]).To(Equal("is required"))
			})

			It("Should return a notification when status is empty", func() {
				status := valueobject.NewStatus("")
				salesChannel := entity_saleschannel.NewSalesChannel("salesChannel", "123", status, true)

				notification := salesChannel.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["status"]).To(Equal("is required"))
			})

			It("Should return a notification when sales channel, code and status are empty", func() {
				status := valueobject.NewStatus("")
				salesChannel := entity_saleschannel.NewSalesChannel("", "", status, true)

				notification := salesChannel.Validate()

				Expect(notification).To(HaveLen(3))
				Expect(notification[0]["sales_channel"]).To(Equal("is required"))
			})

			It("Should return a notification when status is invalid", func() {
				status := valueobject.NewStatus("invalid")
				salesChannel := entity_saleschannel.NewSalesChannel("salesChannel", "123", status, true)

				notification := salesChannel.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["status"]).To(Equal("is invalid, valid values are active or inactive"))
			})

			It("Should not return a notification when status is active", func() {
				status := valueobject.NewStatus("active")
				salesChannel := entity_saleschannel.NewSalesChannel("salesChannel", "123", status, true)

				notification := salesChannel.Validate()

				Expect(notification).To(HaveLen(0))
			})

			It("Should not return a notification when status is inactive", func() {
				status := valueobject.NewStatus("inactive")
				salesChannel := entity_saleschannel.NewSalesChannel("salesChannel", "123", status, true)

				notification := salesChannel.Validate()

				Expect(notification).To(HaveLen(0))
			})
		})
	})
})
