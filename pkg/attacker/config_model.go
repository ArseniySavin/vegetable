package attacker

import (
	"log"
	"time"
)

var (
	templateBody []byte
)

const (
	ResponseFileName = "target.response"
	targetFilename   = "target.body.json"
	dataFilename     = "target.data.json"
	example          = `[
		{	
			"key1": "data1",
			"key2": "data2"
		}
	]`
)

type (
	LoadCfg struct {
		DataPath       string
		Name           string
		Duration       time.Duration
		Per            time.Duration
		Timeout        time.Duration
		Freq           int
		MaxConnections int
		ResponseSave   bool
	}
)

// Use string as 1s, 1m etc
func SetDuration(val string) time.Duration {
	duration, err := time.ParseDuration(val)
	if err != nil {
		log.Fatal(err)
	}
	return duration
}

// Use string as 1s, 1m etc
func SetTimeout(val string) time.Duration {
	timeout, err := time.ParseDuration(val)
	if err != nil {
		log.Fatal(err)
	}
	return timeout
}

// Use string as 1s, 1m etc
func SetPerSeconds(val string) time.Duration {
	timeout, err := time.ParseDuration(val)
	if err != nil {
		log.Fatal(err)
	}
	return timeout
}
