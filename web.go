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
	r.ParseForm()
	log.Println(r.PostForm)
	people := r.PostForm.Get("text")
	if len(people) == 0 {
		log.Println("Misread the response. Couldn't find the text")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	token := r.PostForm.Get("token")
	if len(token) == 0 {
		log.Println("Misread the response. Couldn't find the token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if token != Conf.SlackToken {
		log.Println("Token mismatch.")
		return // Just drop the request. Don't tell them anything
	}

	peopleNum, err := strconv.Atoi(people)
	if err != nil {
		ret, err := json.Marshal(
			Response{ResponseType: "ephemeral",
				Text: "You entered something that's not a number."})
		if err != nil {
			log.Println("Couldn't return the response for not a number error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Return the response
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(ret))
		return
	}

	if peopleNum > Conf.MaxPeople {
		ret, err := json.Marshal(
			Response{ResponseType: "ephemeral",
				Text: "Too many people. Please enter a number less than" + strconv.Itoa(Conf.MaxPeople)})
		if err != nil {
			log.Println("Couldn't return the response for out of bounds error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Return the response
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(ret))
		return
	}

	v, err := GetData(InitClient())
	if err != nil {
		log.Println("Couldn't return the response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
			log.Println("Couldn't return the response for pizza advisement:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Return the response
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(ret))
		return
	} else {
		ret, err := json.Marshal(Response{ResponseType: "ephemeral", Text: "Not enough records yet. Add some more pizza to the spreadsheet."})
		if err != nil {
			log.Println("Couldn't return the response for not enough records:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Return the response
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(ret))
		return
	}
}

// SetUpAPICalls creates handler functions for api calls
func SetUpAPICalls() {
	fmt.Println("Setting up handlers...")
	http.HandleFunc("/api/pizza", PeopleToPizzaHandler)
}
