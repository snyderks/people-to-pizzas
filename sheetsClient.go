package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/api/googleapi/transport"
	sheets "google.golang.org/api/sheets/v4"
)

// Ctx is the context used for the sheets client
var Ctx context.Context

// Conf is the application's configuration
var Conf Config

func init() {
	Ctx = context.Background()

	c, err := ConfigRead("config.json")
	if err != nil {
		log.Fatalf("Could not get config %v", err)
	}
	Conf = c
}

// InitClient create a new Sheets Client.
func InitClient() *sheets.Service {
	client := &http.Client{Transport: &transport.APIKey{Key: Conf.SheetsKey}}

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}

	return srv
}

// GetData retrieves spreadsheet data.
func GetData(s *sheets.Service) []PizzaRecord {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Failed to parse record", r)
		}
	}()

	resp, err := s.Spreadsheets.Values.Get(Conf.SheetID, Conf.SheetRange).Context(Ctx).Do()
	if err != nil {
		log.Fatalf("Failed to get spreadsheet.")
	}
	recs := make([]PizzaRecord, 0)
	if len(resp.Values) > 0 {
		for _, row := range resp.Values {
			pizzas, _ := strconv.Atoi(row[0].(string))
			people, _ := strconv.Atoi(row[1].(string))
			var tooLittle bool
			if row[2].(string) == "FALSE" {
				tooLittle = false
			} else {
				tooLittle = true
			}
			leftOver, _ := strconv.ParseFloat(row[3].(string), 64)
			recs = append(recs, PizzaRecord{
				Pizzas:         pizzas,
				People:         people,
				TooLittle:      tooLittle,
				PizzasLeftOver: leftOver,
			})
		}
	}

	return recs
}
