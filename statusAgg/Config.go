package statusAggregator

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"os"
	"time"
)

// Configuration struct to be read in from config file
type Configuration struct {
	Port           int
	CheckFrequency time.Duration
	Timeout        time.Duration
	Sites          []string
	JSONLogs       bool
}

// GetConfig Gets the configuration from env/{GO_ENV}.json
func GetConfig(configPath string) Configuration {
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalln("Can't open file at", configPath)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	configuration := Configuration{}

	err = decoder.Decode(&configuration)

	if err != nil {
		log.Fatalln("Unable to decode json")
	}

	return configuration
}
