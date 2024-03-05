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
				group := group.NewGroup("", "123", status)

				notification := group.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["group_name"]).To(Equal("is required"))
			})

			It("Should return a notification when code is empty", func() {
				status := valueobject.NewStatus("active")
				group := group.NewGroup("group", "", status)

				notification := group.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["code"]).To(Equal("is required"))
			})

			It("Should return a notification when status is empty", func() {
				status := valueobject.NewStatus("")
				group := group.NewGroup("group", "123", status)

				notification := group.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["status"]).To(Equal("is required"))
			})

			It("Should return a notification when group name, code and status are empty", func() {
				status := valueobject.NewStatus("")
				group := group.NewGroup("", "", status)

				notification := group.Validate()

				Expect(notification).To(HaveLen(3))
				Expect(notification[0]["group_name"]).To(Equal("is required"))
				Expect(notification[1]["code"]).To(Equal("is required"))
				Expect(notification[2]["status"]).To(Equal("is required"))
			})

			It("Should return a notification when status is invalid", func() {
				status := valueobject.NewStatus("invalid")
				group := group.NewGroup("group", "123", status)

				notification := group.Validate()

				Expect(notification).To(HaveLen(1))
				Expect(notification[0]["status"]).To(Equal("is invalid, valid values are active or inactive"))
			})

			It("Should return not return a notification when status is active", func() {
				status := valueobject.NewStatus("active")
				group := group.NewGroup("group", "123", status)

				notification := group.Validate()

				Expect(notification).To(HaveLen(0))
			})

			It("Should return not return a notification when status is inactive", func() {
				status := valueobject.NewStatus("inactive")
				group := group.NewGroup("group", "123", status)

				notification := group.Validate()

				Expect(notification).To(HaveLen(0))
			})
		})
	})
})
