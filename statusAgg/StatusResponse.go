package statusAggregator

import (
	"fmt"
	"time"
)

// StatusResponse Maintain state of last status code for each site
type StatusResponse struct {
	Key        string
	StatusCode int
	LastCheck  time.Time
}

func (sr StatusResponse) String() string {
	return fmt.Sprintf("%s %d %s", sr.LastCheck.Format("2006-01-02T15:04:05.999 MST"), sr.StatusCode, sr.Key)
}
