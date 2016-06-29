package main

import (
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	sa "eLabSolutions/statusAggregator/statusAgg"

	log "github.com/Sirupsen/logrus"
)

var wg sync.WaitGroup
var config sa.Configuration

func init() {
	log.SetLevel(log.DebugLevel)

	config = sa.GetConfig(os.Getenv("SACONFIGFILE"))

	if config.JSONLogs {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

// Check response code of site specified every time message is
// received from ticker channel , then sent StatusResponse to msgs channel
func monitorSite(site string, msgs chan sa.StatusResponse, ticker <-chan time.Time) {
	log.Info("Monitoring ", site)
	timeout := time.Duration(config.Timeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	for t := range ticker {
		resp, err := client.Get(site)

		var sr sa.StatusResponse

		if err != nil {
			log.Error("Error ", err)
			sr = sa.StatusResponse{Key: site, StatusCode: 503, LastCheck: t}
		} else {
			sr = sa.StatusResponse{Key: site, StatusCode: resp.StatusCode, LastCheck: t}
			resp.Body.Close()
		}

		msgs <- sr
	}

	// notify wait group that we're exiting the function, which should
	// never happen
	wg.Done()
}

func main() {
	wg.Add(1) // it is a failure for any of them to return
	log.Info("App startup")

	messages := make(chan sa.StatusResponse, 2)

	// Create goroutines for each site
	for _, s := range config.Sites {
		myTicker := time.NewTicker(config.CheckFrequency * time.Second)

		sa.StatusMap[s] = sa.StatusResponse{Key: s, StatusCode: 0, LastCheck: time.Unix(0, 0)}

		go monitorSite(s, messages, myTicker.C)
	}

	// Goroutine to update status map
	go sa.StatusMapUpdater(messages)

	http.HandleFunc("/", sa.WebHandler)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), nil))

	// Should never happen; app is not meant to exit any goroutine
	wg.Wait()

	log.Fatal("App exiting (not supposed to happen)")
}
