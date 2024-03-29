//go:build unit
// +build unit

package stores_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/andreis3/stores-ms/internal/interface/http/stores"
)

var _ = Describe("INTERFACE :: HTTP :: STORES :: STORES_CONTROLLER", func() {
	Describe("#CreateStores", func() {
		Context("When I call the method CreateStores", func() {
			It("Should create a new store", func() {
				storeController := stores.NewStoresController()

				req, err := http.NewRequest("POST", "/stores", nil)
				rr := httptest.NewRecorder()

				expected := `{"id": "123"}`

				storeController.CreateStores(rr, req)

				Expect(err).To(BeNil())
				Expect(req).NotTo(BeNil())
				Expect(rr.Code).To(Equal(http.StatusCreated))
				Expect(rr.Body.String()).To(Equal(expected))

			})
		})
	})
})
