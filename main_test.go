package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

const path = "testConfig.json"
const incorrectJSONPath = "badConfig.json"

// TestReadConfig attempts to read a valid configuration.
func TestReadConfig(t *testing.T) {
	config, err := ConfigRead(path)
	if err != nil {
		t.Errorf("ReadConfig returned an error: %s", err)
	}
	if len(config.SlackToken) == 0 {
		t.Error("SlackToken was not read.")
	}
	if len(config.HTTPPort) == 0 {
		t.Error("HTTPPort was not read.")
	}
	if len(config.Hostname) == 0 {
		t.Error("Hostname was not read.")
	}
	if len(config.SheetsKey) == 0 {
		t.Error("Google Sheets API token was not read.")
	}
	if len(config.SheetRange) == 0 {
		t.Error("Spreadsheet range selection was not read.")
	}
	if config.TooLittleIncrease == 0 {
		t.Error("Increase was not read.")
	}
	if config.MaxPeople == 0 {
		t.Error("Max people was not read.")
	}
}

// TestFailParseConfig tests correct failure if the config file could not be parsed.
func TestFailParseConfig(t *testing.T) {
	_, err := ConfigRead(incorrectJSONPath)
	if err == nil {
		t.Error("No error was returned")
	}
}

func TestSetUpHandlers(t *testing.T) {
	SetUpAPICalls()
}

func TestPeopleToPizzas(t *testing.T) {
	PeopleToPizzas(0, nil)
}

func TestPeopleToPizzaHandler(t *testing.T) {
	c, _ := ConfigRead("config.json")
	req := httptest.NewRequest("POST", "/api/pizza?token="+c.SlackToken+"&text=5", nil)
	w := httptest.NewRecorder()
	PeopleToPizzaHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var data Response
	err := json.Unmarshal(body, &data)
	if err != nil {
		t.Error("Response was the wrong format:", data)
		return
	}
	if len(data.Text) == 0 {
		t.Error("Response was empty.")
		return
	}
}
