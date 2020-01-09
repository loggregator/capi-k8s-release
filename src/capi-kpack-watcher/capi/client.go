package capi

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"capi_kpack_watcher/model"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"
	uaaClient "code.cloudfoundry.org/uaa-go-client"
	uaaConfig "code.cloudfoundry.org/uaa-go-client/config"
)

// PATCHBuild send a PATCH request to CAPI with a guid and status about a build.
func (c *client) PATCHBuild(guid string, status model.BuildStatus) error {
	const endpoint string = "v3/internal/builds"

	url := fmt.Sprintf("https://api.%s/%s/%s", c.host, endpoint, guid)

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(status.ToJSON()))
	if err != nil {
		return err
	}

	// TODO: Ideally this header should be set on the client, rather than for
	// every request.
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", c.Token())

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Printf("[PATCHBuild] Sent payload: %s\n", status.ToJSON())
	log.Printf("[PATCHBuild] Response Status: %d\n", resp.StatusCode)

	return nil
}

func NewUAAClient() uaaClient.Client {
	logger := lager.NewLogger("test")
	clock := clock.NewClock()

	client, err := uaaClient.NewClient(
		logger,
		&uaaConfig.Config{
			ClientName:       os.Getenv("UAA_CLIENT_NAME"),
			ClientSecret:     os.Getenv("UAA_CLIENT_SECRET"),
			UaaEndpoint:      os.Getenv("UAA_ENDPOINT"),
			SkipVerification: true,
		},
		clock,
	)
	if err != nil {
		panic(err)
	}

	return client
}

func (c *client) Token() string {
	token, err := c.uaaClient.FetchToken(true)
	if err != nil {
		panic(err)
	}

	return "Bearer " + token.AccessToken
}

// NewCAPIClient creates a client to be used to communicate with CAPI. This
// client is a HTTP client for talking to CAPI over a REST API.
func NewCAPIClient() CAPI {
	// TODO: We may want to consider using cloudfoundry/tlsconfig for using
	// standard TLS configs in Golang.
	return &client{
		host: os.Getenv("CAPI_HOST"),
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		uaaClient: NewUAAClient(),
	}
}

type client struct {
	host       string
	httpClient *http.Client
	uaaClient  uaaClient.Client
}
