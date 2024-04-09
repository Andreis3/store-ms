//go:build integration
// +build integration

package create_group_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func Test_IntegrationCreateGroup(t *testing.T) {
	suiteConfig, reporterConfig := GinkgoConfiguration()

	suiteConfig.SkipStrings = []string{"NEVER-RUN"}
	reporterConfig.FullTrace = true
	reporterConfig.Succinct = true

	RegisterFailHandler(Fail)
	RunSpecs(t, "Group Integration Create Group Test Suite ", suiteConfig, reporterConfig)
}

var _ = Describe("INTEGRATION :: TEST :: CREATE :: NEW :: GROUP", func() {
	Describe("#POST", func() {
		Context("When send request to route /api/v1/groups", func() {
			When("When send request with valid data", func() {
				var client = &http.Client{}
				BeforeEach(func() {
					TruncateTable(ConnectionPostgres())
				})
				AfterEach(func() {
					TruncateTable(ConnectionPostgres())
				})
				It("Should return status 201", func() {
					body, _ := json.Marshal(map[string]interface{}{
						"group_name": "test 1",
						"code":       "test 1",
						"status":     "active",
					})
					payload := bytes.NewBuffer(body)

					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/groups", payload)
					defer req.Body.Close()
					req.Header.Set("Content-Type", "application/json")

					res, err := client.Do(req)

					response := map[string]any{}

					json.NewDecoder(res.Body).Decode(&response)

					Expect(err).To(BeNil())
					Expect(res.StatusCode).To(Equal(http.StatusCreated))
					Expect(response["request_id"]).NotTo(BeNil())
					Expect(response["status_code"].(float64)).To(Equal(float64(http.StatusCreated)))
					Expect(response["data"].(map[string]any)["group_name"]).To(Equal("test 1"))
					Expect(response["data"].(map[string]any)["code"]).To(Equal("test 1"))
					Expect(response["data"].(map[string]any)["status"]).To(Equal("active"))
					Expect(response["data"].(map[string]any)["id"]).NotTo(BeNil())
					Expect(response["data"].(map[string]any)["created_at"]).NotTo(BeNil())
					Expect(response["data"].(map[string]any)["updated_at"]).NotTo(BeNil())

				})
			})

			When("When send two request with same data", func() {
				var client = &http.Client{}
				BeforeEach(func() {
					TruncateTable(ConnectionPostgres())
				})
				AfterEach(func() {
					TruncateTable(ConnectionPostgres())
				})
				It("Should return status 500", func() {
					body, _ := json.Marshal(map[string]interface{}{
						"group_name": "test 1",
						"code":       "test 1",
						"status":     "active",
					})
					body2, _ := json.Marshal(map[string]interface{}{
						"group_name": "test 1",
						"code":       "test 1",
						"status":     "active",
					})
					payload := bytes.NewBuffer(body)
					payload2 := bytes.NewBuffer(body2)

					req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/groups", payload)
					req.Header.Set("Content-Type", "application/json")

					client.Do(req)

					req, _ = http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/groups", payload2)
					defer req.Body.Close()
					req.Header.Set("Content-Type", "application/json")

					res, err := client.Do(req)

					response := map[string]any{}

					json.NewDecoder(res.Body).Decode(&response)

					Expect(err).To(BeNil())

					Expect(res.StatusCode).To(Equal(http.StatusInternalServerError))
					Expect(response["request_id"]).NotTo(BeNil())
					Expect(response["status_code"].(float64)).To(Equal(float64(http.StatusInternalServerError)))
					Expect(response["error_message"].([]any)).To(ContainElement("Internal Server Error"))
				})
			})
		})
	})
})
