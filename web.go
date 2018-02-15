package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Response is the format used to return an response to Slack.
type Response struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

// Here we have wonderful endpoints to convert between people and pizza.

// PeopleToPizzaHandler determines how many pizza for people.
func PeopleToPizzaHandler(w http.ResponseWriter, r *http.Request) {
	people := r.URL.Query().Get("text")
	if len(people) == 0 {
		log.Fatal("Misread the response. Couldn't find the text")
	}
	token := r.URL.Query().Get("token")
	if len(token) == 0 {
		log.Fatal("Misread the response. Couldn't find the token")
	}
	if token != Conf.SlackToken {
		return // Just drop the request. Don't tell them anything
	}

	peopleNum, err := strconv.Atoi(people)
	if err != nil {
		ret, err := json.Marshal(
			Response{ResponseType: "ephemeral",
				Text: "You entered something that's not a number."})
		if err != nil {
			log.Fatal("Couldn't return the response.")
		}
		// Return the response
		w.Write([]byte(ret))
		return
	}

	if peopleNum > Conf.MaxPeople {
		ret, err := json.Marshal(
			Response{ResponseType: "ephemeral",
				Text: "Too many people. Please enter a number less than" + strconv.Itoa(Conf.MaxPeople)})
		if err != nil {
			log.Fatal("Couldn't return the response.")
		}
		// Return the response
		w.Write([]byte(ret))
		return
	}

	v := GetData(InitClient())
	fmt.Println(v)
	if len(v) > 1 {
		result, err := PeopleToPizzas(peopleNum, v)
		if err != nil {
			w.Write([]byte("fail"))
		}
		ret, err := json.Marshal(
			Response{ResponseType: "in_channel",
				Text: "You should order " + strconv.FormatFloat(result, 'g', 5, 64) + " pizzas for " + people + " people."})
		if err != nil {
			log.Fatal("Couldn't return the response.")
		}
		// Return the response
		w.Write([]byte(ret))
		return
	} else {
		ret, err := json.Marshal(Response{ResponseType: "ephemeral", Text: "Not enough records yet. Add some more pizza to the spreadsheet."})
		if err != nil {
			log.Fatal("Couldn't return the response.")
		}
		// Return the response
		w.Write([]byte(ret))
		return
	}
}

// SetUpAPICalls creates handler functions for api calls
func SetUpAPICalls() {
	fmt.Println("Setting up handlers...")
	http.HandleFunc("/api/pizza/", PeopleToPizzaHandler)
}
