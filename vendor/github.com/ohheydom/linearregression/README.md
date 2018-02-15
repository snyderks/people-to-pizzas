# Linear Regression

Linear regression written in golang. It utilizes Gradient Descent to calculate the weights.

## Usage

### Import

Import the package

```golang
import (
  "github.com/ohheydom/linearregression"
)
```

### Create a sample with labels

Samples need to be slices of float64s. The training set x should be a multidimensional slice. y should contain the labels for each corresponding x sample.

```golang
X := [][]float64{[]float64{1, 1}, []float64{2, 2}, []float64{3, 3}, []float64{4, 4}}
y := []float64{1, 2, 3, 4}
```

### Create a Linear Regression struct and Fit

```golang
lr := linearregression.LinearRegression{}
lr.fit(x, y)
```

### Show weights

The first weight listed will be the intercept value. The following weights will be the coefficients for each feature.

```golang
fmt.Println(lr.Weights)
```

Examples folder will have more examples soon.
