package api

import (
	"crypto/tls"
	"net/http"

	"capi_kpack_watcher/model"
)

// API is the REST API for CAPI actions.
type APIRequester interface {
	UpdateBuild(token, guid string, status model.BuildStatus) error
}

func NewAPI(host string) *api {
	// TODO: We may want to consider using cloudfoundry/tlsconfig for using
	// standard TLS configs in Golang.
	return &api{
		host: host,
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}
}

type api struct {
	host       string
	httpClient *http.Client
}
