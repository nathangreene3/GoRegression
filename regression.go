package main

import (
	"log"
	"math"
)

// polyCoefs represent polynomial coefficients.
type polyCoefs []float64

// fit returns a set of weights representing a polynomial of a given degree trained
// on a set of two-dimensional data points.
func fit(degrees int, pnts points) polyCoefs {
	if degrees < 1 {
		panic("fit: degrees must be positive")
	}

	if len(pnts) < 2 {
		panic("fit: must provide at least two points")
	}

	w := make(polyCoefs, degrees+1)
	w.train(pnts)
	return w
}

// output returns w0 + w1*x + ... + wn-1*x^(n-1).
func (w polyCoefs) output(x float64) float64 {
	var y, n float64 // y = w0 + w1*x + ... + wn-1*x^(n-1); n = float64(i)
	for i := range w {
		y += w[i] * math.Pow(x, n)
		n++
	}

	return y
}

// loss returns the sum of the squared error of each y to the output of x over the
// set of weights.
func (w polyCoefs) loss(pnts points) float64 {
	var e float64
	for i := range pnts {
		e += math.Pow(pnts[i][1]-w.output(pnts[i][0]), 2)
	}

	return e
}

// gradLoss returns the gradient of the loss function over the set of weights'
// output over the two-dimensional training points.
func (w polyCoefs) gradLoss(pnts points) polyCoefs {
	n := len(w)
	derivWeights := make(polyCoefs, 0, n) // derivatives of weights
	var x, v, k float64                   // temp; deriv of ith weight; float of i

	for i := 0; i < n; i++ {
		for j := range pnts {
			x = pnts[j][0]
			v += math.Pow(x, k) * (pnts[j][1] - w.output(x))
		}

		derivWeights = append(derivWeights, -2*v)
		v = 0
		k++
	}

	return derivWeights
}

// train updates the weights using gradient descent given a set of two-dimensional
// points and a learning rate.
func (w polyCoefs) train(pnts points) {
	var derivWeights []float64
	var oldWeight float64
	a := 0.01
	changed := true

	for changed {
		changed = false
		derivWeights = w.gradLoss(pnts)

		for i := range w {
			oldWeight = w[i]
			w[i] -= a * derivWeights[i]

			if w[i] != oldWeight {
				changed = true
			}

			if math.IsNaN(w[i]) || math.IsInf(w[i], 0) {
				log.Fatal(w)
			}
		}
	}
}
