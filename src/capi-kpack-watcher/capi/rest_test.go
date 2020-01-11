package capi

import (
	"capi_kpack_watcher/mocks"
	"github.com/sclevine/spec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRestClient_PATCH(t *testing.T) {
	spec.Run(t, "TestRestClient_PATCH", func(t *testing.T, when spec.G, it spec.S) {
		var (
			client     Client
			testServer *httptest.Server
		)

		it.Before(func() {
			testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			client = Client{
			}

		})

		it.After(func() {
			mock.AssertExpectationsForObjects(t, mockCAPI)
		})

		when("something useful", func() {
			it.Before(func() {

				testServer.Start()
			})

			it.After(func() {
				testServer.Close()
			})

			it("does useful things", func() {
				assert.Equal(t, 1, 1)
			})
		})
	})
}
