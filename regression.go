package main

import (
	"math"
)

type weights []float64

// fit returns a set of weights trained on a set of two-dimensional data points.
func fit(degrees int, x, y []float64) weights {
	if degrees < 1 {
		panic("fit: degrees must be positive")
	}

	w := make(weights, degrees+1)
	w.train(x, y, 0.001)
	return w
}

// output returns w0 + w1*x + ... + wn-1*x^(n-1).
func (w weights) output(x float64) float64 {
	var y, n float64 // y = w0 + w1*x + ... + wn-1*x^(n-1); n = float64(i)
	for i := range w {
		y += w[i] * math.Pow(x, n)
		n++
	}

	return y
}

// loss returns the sum of the squared error of each y to the output of x over the set of weights.
func (w weights) loss(x, y []float64) float64 {
	n := len(x)
	if n != len(y) {
		panic("loss: dimension mismatch")
	}

	var e float64
	for i := 0; i < n; i++ {
		e += math.Pow(y[i]-w.output(x[i]), 2)
	}

	return e
}

// gradLoss returns the gradient of the set of weights' output over the training data x and y.
func (w weights) gradLoss(x, y []float64) []float64 {
	n := len(x)
	if n != len(y) {
		panic("gradLoss: dimension mismatch")
	}

	N := len(w)
	derivWeights := make([]float64, 0, N) // derivatives of weights
	var t, v, k float64                   // temp, deriv of ith weight, float of i

	for i := 0; i < N; i++ {
		for j := 0; j < n; j++ {
			t = x[j]
			v += math.Pow(t, k) * (y[j] - w.output(t))
		}

		derivWeights = append(derivWeights, -2*v)
		v = 0
		k++
	}

	return derivWeights
}

func (w weights) train(x, y []float64, a float64) {
	if a <= 0 || 1 <= a {
		panic("train: learning rate must be on range (0,1)")
	}

	changed := true // Indicates if a weight updates to new value
	var derivWeights []float64
	var oldWeight float64
	for changed {
		changed = false
		derivWeights = w.gradLoss(x, y)
		for i := range w {
			oldWeight = w[i]
			w[i] -= a * derivWeights[i]
			if w[i] != oldWeight {
				changed = true
			}
		}
	}
}
