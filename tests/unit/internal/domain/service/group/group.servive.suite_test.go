//go:build unit
// +build unit

package group_service_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func Test_GroupServiceSuite(t *testing.T) {
	suiteConfig, reporterConfig := GinkgoConfiguration()

	suiteConfig.SkipStrings = []string{"SKIPPED", "PENDING", "NEVER-RUN", "SKIP"}
	reporterConfig.FullTrace = true
	reporterConfig.Succinct = true

	RegisterFailHandler(Fail)
	RunSpecs(t, "Group Service Insert Test Suite ", suiteConfig, reporterConfig)
}
