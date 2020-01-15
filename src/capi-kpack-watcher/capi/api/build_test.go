package api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"capi_kpack_watcher/model"

	"github.com/sclevine/spec"
)

func TestUpdateBuild(t *testing.T) {
	spec.Run(t, "TestUpdateBuild", func(t *testing.T, when spec.G, it spec.S) {
		//const (
		//	guid = "guid"
		//)
		//var (
		//	mockAPI     *mocks.APIRequester
		//	buildStatus model.BuildStatus
		//)
		//
		//it.Before(func() {
		//	mockAPI = new(mocks.APIRequester)
		//	buildStatus = model.BuildStatus{State: "STAGING"}
		//})
		//
		//it.After(func() {
		//	mock.AssertExpectationsForObjects(t, mockAPI)
		//})
		//
		//when("something useful", func() {
		//	it.Before(func() {
		//	})
		//
		//	it("does useful things", func() {
		//		buildStatus.State = "who cares"
		//		assert.Equal(t, 1, 1)
		//	})
		//})



		const (
			guid = "guid"
		)

		var (
			API     api
			buildStatus model.BuildStatus
			testServer *httptest.Server
			authToken string
		)

		it.Before(func() {
			buildStatus = model.BuildStatus{State: "SUCCESS"}

			testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, r.URL.String(), "/")

				//b, err := ioutil.ReadAll(r.Body)
				//if err != nil {
				//	panic(err)
				//}
				//assert.Equal(t, b, buildStatus.ToJSON())

				//assert.Equal(t, r.Header.Get("Authorization"), authToken)
				//assert.Equal(t, r.Header.Get("Content-Type"), "application/json")

				w.WriteHeader(http.StatusOK)
			}))
			authToken = "valid-auth-token-returned-by-uaa"
			API = api{host: testServer.URL , httpClient: testServer.Client()}
		})

		when("request is valid", func() {
			it.After(func() {
				testServer.Close()
			})

			it("responds with StatusOK returned from CAPI", func() {
				err := API.UpdateBuild(authToken, guid, buildStatus)
				//assert.Equal(t, http.StatusOK, response.StatusCode)
				assert.NoError(t, err)
			})
		})
	})
}
