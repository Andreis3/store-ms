//go:build unit
// +build unit

package group_command_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test_GroupCommandSuite(t *testing.T) {
	suiteConfig, reporterConfig := GinkgoConfiguration()
	suiteConfig.SkipStrings = []string{"SKIPPED", "PENDING", "NEVER-RUN", "SKIP"}
	reporterConfig.FullTrace = true
	reporterConfig.Succinct = true
	RegisterFailHandler(Fail)
	RunSpecs(t, "Group Command Insert Test Suite ", suiteConfig, reporterConfig)
}
