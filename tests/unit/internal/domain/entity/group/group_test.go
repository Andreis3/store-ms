//go:build unit
// +build unit

package group_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/domain/entity/group"
	"github.com/andreis3/stores-ms/internal/domain/valueobject"
)

func Test_GroupEntitySuite(t *testing.T) {
	suiteConfig, reporterConfig := GinkgoConfiguration()

	suiteConfig.SkipStrings = []string{"SKIPPED", "PENDING", "NEVER-RUN", "SKIP"}
	reporterConfig.FullTrace = true
	reporterConfig.Succinct = true

	RegisterFailHandler(Fail)
	RunSpecs(t, "Group Entity Test Suite ", suiteConfig, reporterConfig)
}

var _ = Describe("DOMAIN :: ENTITY :: GROUP", func() {
	Describe("#Validate", func() {
		Context("When I call the method Validate", func() {
			It("Should return a notification when group name is empty", func() {
				status := valueobject.NewStatus("active")
				group := entity_group.NewGroup("", "123", status)

				notification := group.Validate()

				Expect(notification.ReturnNotification()).To(HaveLen(1))
				Expect(notification.ReturnNotification()[0]).To(Equal("name: is required"))
			})

			It("Should return a notification when code is empty", func() {
				status := valueobject.NewStatus("active")
				group := entity_group.NewGroup("group", "", status)

				notification := group.Validate()

				Expect(notification.ReturnNotification()).To(HaveLen(1))
				Expect(notification.ReturnNotification()[0]).To(Equal("code: is required"))
			})

			It("Should return a notification when status is empty", func() {
				status := valueobject.NewStatus("")
				group := entity_group.NewGroup("group", "123", status)

				notification := group.Validate()

				Expect(notification.ReturnNotification()).To(HaveLen(1))
				Expect(notification.ReturnNotification()[0]).To(Equal("status: is required"))
			})

			It("Should return a notification when group name, code and status are empty", func() {
				status := valueobject.NewStatus("")
				group := entity_group.NewGroup("", "", status)

				notification := group.Validate()

				Expect(notification.ReturnNotification()).To(HaveLen(3))
				Expect(notification.ReturnNotification()[0]).To(Equal("name: is required"))
				Expect(notification.ReturnNotification()[1]).To(Equal("code: is required"))
				Expect(notification.ReturnNotification()[2]).To(Equal("status: is required"))
			})

			It("Should return a notification when status is invalid", func() {
				status := valueobject.NewStatus("invalid")
				group := entity_group.NewGroup("group", "123", status)

				notification := group.Validate()

				Expect(notification.ReturnNotification()).To(HaveLen(1))
				Expect(notification.ReturnNotification()[0]).To(Equal("status: is invalid, valid values are active or inactive"))
			})

			It("Should return not return a notification when status is active", func() {
				status := valueobject.NewStatus("active")
				group := entity_group.NewGroup("group", "123", status)

				notification := group.Validate()

				Expect(notification.ReturnNotification()).To(HaveLen(0))
			})

			It("Should return not return a notification when status is inactive", func() {
				status := valueobject.NewStatus("inactive")
				group := entity_group.NewGroup("group", "123", status)

				notification := group.Validate()

				Expect(notification.ReturnNotification()).To(HaveLen(0))
			})
		})
	})
})
