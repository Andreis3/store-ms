package stores_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test_StoresInterfaceSuite(t *testing.T) {
	// fetch the current config
	suiteConfig, reporterConfig := GinkgoConfiguration()

	// adjust it
	//suiteConfig.ParallelTotal = 3
	suiteConfig.SkipStrings = []string{"NEVER-RUN"}
	reporterConfig.Verbose = true
	reporterConfig.FullTrace = true

	RegisterFailHandler(Fail)
	RunSpecs(t, "Store Test Suite ", suiteConfig, reporterConfig)
}
