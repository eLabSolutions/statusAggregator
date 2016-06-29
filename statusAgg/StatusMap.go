package statusAggregator

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
)

// StatusMap Map of [siteKey] => StatusResponse
var StatusMap = make(map[string]StatusResponse)

// StatusMapUpdater update the map
func StatusMapUpdater(msgs chan StatusResponse) {
	log.Info("Waiting for results")

	// Read messages from channel and populate StatusMap
	for resp := range msgs {
		StatusMap[resp.Key] = resp

		log.Debug(StatusMap[resp.Key])
	}
}

// WebHandler returns the maximum status code; under normal conditions
// this should always be 200
func WebHandler(w http.ResponseWriter, r *http.Request) {
	maxStatusCode := 0

	for _, value := range StatusMap {
		if value.StatusCode > maxStatusCode {
			maxStatusCode = value.StatusCode
		}
	}

	if maxStatusCode == 0 {
		maxStatusCode = http.StatusInternalServerError
	}

	w.WriteHeader(maxStatusCode)
}
