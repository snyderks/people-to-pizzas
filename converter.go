package main

import (
	"errors"
	"fmt"

	"github.com/ohheydom/linearregression"
)

// PizzaRecord is an event of pizza being served
// with the amount left over recorded
type PizzaRecord struct {
	Pizzas         int
	People         int
	TooLittle      bool
	PizzasLeftOver float64
}

func PeopleToPizzas(people int, past []PizzaRecord) (result float64, err error) {
	var x [][]float64
	var y []float64

	for _, rec := range past {
		baseVal := (float64(rec.Pizzas) - rec.PizzasLeftOver)
		if rec.TooLittle {
			baseVal *= Conf.TooLittleIncrease
		}
		x = append(x, []float64{(float64(rec.People))})
		y = append(y, baseVal)
	}

	fmt.Println(x)
	fmt.Println(y)

	if len(x) < 2 || len(y) < 2 {
		return 0, errors.New("Not enough past events.")
	}

	l := linearregression.LinearRegression{Method: "sls"}
	l.Fit(x, y)
	fmt.Println(l.Weights)

	predicted := l.Predict([][]float64{[]float64{float64(people)}})
	fmt.Println(predicted)
	fmt.Println(predicted[0])
	return predicted[0], nil
}
