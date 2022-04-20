package router

import (
	"encoding/json"
	"net/http"
)

type VersionResponse struct {
	Version string `json:"version"`
}

func (router *Router) HandleVersion(w http.ResponseWriter, r *http.Request) {
	resPayload, err := json.Marshal(VersionResponse{
		Version: router.Settings.Version,
	})
	if err != nil {
		SendError("error while marshalling response", w)
		return
	}

	w.Write(resPayload)
}
