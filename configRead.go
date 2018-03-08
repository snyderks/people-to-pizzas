package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Config holds the different configuration options.
type Config struct {
	SlackToken        string  `json:"slack-token"`
	SheetsKey         string  `json:"sheets-key"`
	SheetID           string  `json:"sheet-id"`
	SheetRange        string  `json:"sheet-range"`
	TooLittleIncrease float64 `json:"too-little-increase"`
	MaxPeople         int     `json:"max-people"`
	HTTPPort          string  `json:"http-port"`
	Hostname          string  `json:"hostname,omitempty"`
	Debug             bool    `json:"debug"`
}

// ConfigRead takes a path to a JSON file.
// If it fails to read the file, it falls back to environment variables.
// Returns an error if it can't parse the JSON file or if it can't read environment variables.
func ConfigRead(path string) (Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil { // not using json config. Try to get it from env vars
		tooLittleStr := os.Getenv("PIZZA_TOO_LITTLE_INCREASE")
		tooLittle, err := strconv.ParseFloat(tooLittleStr, 64)
		if err != nil {
			log.Fatal("Couldn't convert PIZZA_TOO_LITTLE_INCREASE env var to float64.")
		}
		maxPeopleStr := os.Getenv("PIZZA_MAX_PEOPLE")
		maxPeople, err := strconv.Atoi(maxPeopleStr)
		if err != nil {
			log.Fatal("Couldn't convert PIZZA_MAX_PIZZAS env var to int.")
		}
		config := Config{
			SlackToken:        os.Getenv("PIZZA_SLACK_TOKEN"),
			SheetsKey:         os.Getenv("PIZZA_SHEETS_KEY"),
			SheetID:           os.Getenv("PIZZA_SHEET_ID"),
			SheetRange:        os.Getenv("PIZZA_SHEET_RANGE"),
			TooLittleIncrease: tooLittle,
			MaxPeople:         maxPeople,
			HTTPPort:          os.Getenv("PORT"),
			Hostname:          os.Getenv("HOSTNAME"),
			Debug:             os.Getenv("DEBUG") == "1",
		}
		if !strings.Contains(config.HTTPPort, ":") {
			config.HTTPPort = ":" + config.HTTPPort
		}
		return config, nil
	}
	config := Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return Config{}, err
	}
	return config, err
}
