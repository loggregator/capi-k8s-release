package capi

import (
	"capi_kpack_watcher/model"
)

// UpdateBuild updates the status of a build designated by guid.
func UpdateBuild(client CAPI, guid string, status model.BuildStatus) error {
	return client.UpdateBuild(guid, status)
}
