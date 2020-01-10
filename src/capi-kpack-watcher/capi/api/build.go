package api

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"capi_kpack_watcher/model"
)

func (a *api) UpdateBuild(token, guid string, status model.BuildStatus) error {
	url := fmt.Sprintf("https://api.%s/v3/internal/builds/%s", a.host, guid)

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(status.ToJSON()))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", token)

	// TODO: Ideally this header should be set on the client, rather than for
	// every request.
	req.Header.Add("Content-Type", "application/json")

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Printf("[API/UpdateBuild] Sent payload: %s\n", status.ToJSON())
	log.Printf("[API/UpdateBuild] Response Status: %d\n", resp.StatusCode)

	return nil
}
