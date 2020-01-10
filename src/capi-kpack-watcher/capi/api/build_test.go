package api

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"capi_kpack_watcher/mocks"
	"capi_kpack_watcher/model"

	"github.com/sclevine/spec"
	"github.com/stretchr/testify/mock"
)

func TestUpdateBuild(t *testing.T) {
	spec.Run(t, "TestUpdateBuild", func(t *testing.T, when spec.G, it spec.S) {
		const (
			guid = "guid"
		)
		var (
			mockAPI *mocks.APIRequester
			status  model.BuildStatus
		)

		it.Before(func() {
			mockAPI = new(mocks.APIRequester)
			status = model.BuildStatus{State: "STAGING"}
		})

		it.After(func() {
			mock.AssertExpectationsForObjects(t, mockAPI)
		})

		when("something useful", func() {
			it.Before(func() {
			})

			it("does useful things", func() {
				status.State = "who cares"
				assert.Equal(t, 1, 1)
			})
		})
	})
}
