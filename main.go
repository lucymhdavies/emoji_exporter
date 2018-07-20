// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// A minimal example of how to include Prometheus instrumentation.
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"net/http"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/alexsasharegan/dotenv"
	"github.com/prometheus/client_golang/prometheus"
)

// Global vars for metrics
var (
	emojiScore = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "lmhd",
			Subsystem: "emoji",
			Name:      "twitter_ranking",
			Help:      "Number of uses of this emoji on twitter",
		},
		[]string{
			// Which emoji?
			"emoji",
			"name",
			"id",
		},
	)

	// TODO: emoji ranking
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

type Config struct {
	LogLevel string
}

var cfg Config

func main() {
	// Load env vars from .env file, if present
	// Ignore errors caused by file not existing
	_ = dotenv.Load()

	// Setup logging before anything else
	if len(os.Getenv("LOG_LEVEL")) == 0 {
		cfg.LogLevel = "info"
	} else {
		cfg.LogLevel = os.Getenv("LOG_LEVEL")
	}
	switch cfg.LogLevel {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	}

	//
	// Register Prometheus Metrics
	//

	prometheus.MustRegister(emojiScore)

	//
	// Loops
	//

	go metricsHandler()

	metricsUpdate()
}

func metricsHandler() {

	flag.Parse()
	http.Handle("/metrics", prometheus.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func metricsUpdate() {

	// Init with rest API
	rankings, err := Rankings()
	if err != nil {
		log.Fatalf("%s", err)
	}

	for _, emoji := range rankings {
		emojiScore.With(prometheus.Labels{
			"emoji": emoji.Char,
			"name":  emoji.Name,
			"id":    emoji.ID,
		}).Set(float64(emoji.Score))
	}

	// TODO: store an ID:Char lookup map

	// TODO: move a bunch of this out into emoji.go, with helper functions and all that jazz
	resp, _ := http.Get("https://stream.emojitracker.com/subscribe/eps")

	reader := bufio.NewReader(resp.Body)
	for {
		line, _ := reader.ReadBytes('\n')
		lineString := string(line)

		// Lines look like
		// data:{"1F449":1,"1F44D":1,"1F60F":1,"26F3":1}

		if strings.HasPrefix(lineString, "data:") {

			data := []byte(strings.TrimPrefix(lineString, "data:"))

			jsonMap := make(map[string]int)
			err = json.Unmarshal(data, &jsonMap)
			if err != nil {
				panic(err)
			}

			for key, val := range jsonMap {
				for _, emoji := range rankings {
					if emoji.ID == key {
						emojiScore.With(prometheus.Labels{
							"emoji": emoji.Char,
							"name":  emoji.Name,
							"id":    emoji.ID,
						}).Add(float64(val))
						log.Debugf("Char: %s (%s) : %d", key, emoji.Name, val)
					}
				}

			}
		}

	}
}
