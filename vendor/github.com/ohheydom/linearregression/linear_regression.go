package linearregression

import (
	"log"
)

// LinearRegression stores the parameters for creating a linear regression model
type LinearRegression struct {
	Weights []float64
	NIter   int
	Method  string
}

// Fit will create the model and store the appropriate weights into the Weight field.
func (l *LinearRegression) Fit(x [][]float64, y []float64) {
	var nIter int
	if l.NIter <= 0 {
		nIter = 200
	} else {
		nIter = l.NIter
	}

	if l.Method == "gd" || l.Method == "" {
		l.Weights = gdSolver(x, y, nIter, 0.01)
	} else if l.Method == "sls" {
		l.Weights = slsSolver(x, y)
	}
}

// gdSolver returns the weights using a Gradient Descent algorithm.
// nIter is the number of iterations to run through before convergence, gamma is the step size.
func gdSolver(x [][]float64, y []float64, nIter int, gamma float64) []float64 {
	n := len(x)
	w := make([]float64, len(x[0])+1)
	errors := make([]float64, n)
	for i := 0; i < nIter; i++ {
		predY := predY(x, w)
		errorSum := 0.0
		for j := 0; j < n; j++ {
			errors[j] = y[j] - predY[j]
			errorSum += errors[j]
		}
		for k := 0; k < n; k++ {
			for l := 1; l < len(w); l++ {
				w[l] += gamma * x[k][l-1] * errors[k]
			}
		}
		w[0] += gamma * errorSum
	}
	return w
}

// slsSolver is a Simple Least Squares solver. It requires only one value for each sample and will return the intercept (w[0]) and slope (w[1]).
func slsSolver(x [][]float64, y []float64) []float64 {
	n := len(x)
	if len(x[0]) > 1 {
		log.Fatal("Simple Least Squares Solver calculates the weights given samples with only one feature.")
	}
	w := make([]float64, 2)
	var xy, xSum, ySum, xSquaredSum float64
	for i := 0; i < n; i++ {
		xy += x[i][0] * y[i]
		xSum += x[i][0]
		ySum += y[i]
		xSquaredSum += x[i][0] * x[i][0]
	}
	w[1] = ((float64(n) * xy) - (xSum * ySum)) / ((float64(n) * xSquaredSum) - (xSum * xSum))
	w[0] = (ySum - (w[1] * xSum)) / float64(n)
	return w
}

// Predict will predict values using the model created with the Fit function
func (l *LinearRegression) Predict(x [][]float64) []float64 {
	return predY(x, l.Weights)
}

// predY uses the given weights to calculate each sample's label.
func predY(x [][]float64, w []float64) []float64 {
	n, nFeatures := len(x), len(x[0])
	predY := make([]float64, n)
	for i := 0; i < n; i++ {
		for j := 1; j <= nFeatures; j++ {
			predY[i] += x[i][j-1] * w[j]
		}
		predY[i] += w[0]
	}
	return predY
}

// Mean Squared Error. Lower is better.
func MeanSquaredError(y, predY []float64) float64 {
	n := len(y)
	var total float64
	for i := 0; i < n; i++ {
		total += (y[i] - predY[i]) * (y[i] - predY[i])
	}
	return total / float64(n)
}
