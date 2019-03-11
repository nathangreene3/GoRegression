package main

import (
	"math"
)

type point []float64
type points []point

func (p point) magnitude() float64 {
	var r float64
	for _, v := range p {
		r += v * v
	}

	return math.Sqrt(r)
}

func equal(p, q point) bool {
	if p == nil {
		return q == nil
	}

	if q == nil {
		return false
	}

	if len(p) != len(q) {
		return false
	}

	for i := range p {
		if p[i] != q[i] {
			return false
		}
	}

	return true
}
