package capi

import (
	"io"
	"log"
	"net/http"
)

func (r *restClient) PATCH(url, authToken string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(
		http.MethodPatch,
		url,
		body,
	)
	if err != nil {
		return nil, err
	}

	log.Printf("[CAPI/PATCH] Sending request PATCH %s", url)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authToken)

	resp, err := r.client.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	return resp, nil
}

type restClient struct {
	client *http.Client
}
