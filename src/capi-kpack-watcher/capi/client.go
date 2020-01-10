package capi

import (
	"os"

	"capi_kpack_watcher/capi/api"
	"capi_kpack_watcher/model"

	uaaClient "code.cloudfoundry.org/uaa-go-client"
)

// UpdateBuild updates the status of a build designated by guid.
func (c *client) UpdateBuild(guid string, status model.BuildStatus) error {
	//TODO: fetch token, refresh token if needed
	token, err := c.FetchToken()
	if err != nil {
		return err
	}

	return c.api.UpdateBuild(token, guid, status)
}

func (c *client) FetchToken() (string, error) {
	// Always force update of token.
	const forceUpdate = true
	token, err := c.uaaClient.FetchToken(forceUpdate)
	if err != nil {
		return "", err
	}

	return "Bearer " + token.AccessToken, nil
}

// NewCAPIClient creates a client to be used to communicate with CAPI.
func NewCAPIClient() CAPI {
	return &client{
		api:       api.NewAPI(os.Getenv("CAPI_HOST")),
		uaaClient: NewUAAClient(),
	}
}

type client struct {
	api       api.APIRequester
	uaaClient uaaClient.Client
}
