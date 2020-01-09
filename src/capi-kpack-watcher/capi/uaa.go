package capi

import (
	"os"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"
	uaaClient "code.cloudfoundry.org/uaa-go-client"
	uaaConfig "code.cloudfoundry.org/uaa-go-client/config"
)

type TokenFetcher interface {
	FetchToken() (string, error)
}

func NewUAAClient() uaaClient.Client {
	logger := lager.NewLogger("")
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
