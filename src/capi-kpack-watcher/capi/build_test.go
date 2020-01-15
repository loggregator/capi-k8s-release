package capi

import (
	"capi_kpack_watcher/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sclevine/spec"
)

func TestUpdateBuild(t *testing.T) {
	spec.Run(t, "TestUpdateBuild", func(t *testing.T, when spec.G, it spec.S) {
		const (
			guid = "guid"
		)
		var (
			client     *Client
			testServer *httptest.Server
		)

		it.Before(func() {
			testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))

			client = new(Client)
			client.host = ""
			client.restClient = new(mocks.REST)
			client.uaaClient = new(mocks.TokenFetcher)
		})

		//it.After(func() {
		//	mock.AssertExpectationsForObjects(t, mockCAPI)
		//})

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
