package main

import (
	"fmt"
	"math"
	"math/rand"
	// "github.com/guptarohit/asciigraph"
)

func main() {
	// rand.Seed(int64(time.Now().Nanosecond()))

	// n := 25
	// pnts := randomPointsNearCurve(n, func(x float64) float64 { return x })
	// y := make([]float64, 0, n)
	// for i := 0; i < n; i++ {
	// 	y = append(y, pnts[i][1])
	// }

	// fmt.Println(asciigraph.Plot(y))

	// pnts := points{
	// 	point{0, 1},
	// 	point{1, 5},
	// 	point{2, 3},
	// 	point{3, 9},
	// 	point{4, 7},
	// }

	// pnts := points{
	// 	point{0, 0},
	// 	point{1, 1},
	// 	point{2, 4},
	// }

	// fmt.Println(fit(1, pnts))
	// fmt.Println(polyCoefs{-1, 1, 2}.minAt(0, 0.001, 0.001))
	// fmt.Println(findBucketMaxVol())
	// fmt.Println(bucketVolume(0.4, 0.4))
	// fmt.Println(boxVolume(0.4, 0.4))
	fmt.Println(findMaxBoxVolume())
}

func randomPointsNearCurve(n int, f func(x float64) float64) points {
	pnts := make(points, 0, n)

	var v float64
	for i := 0; i < n; i++ {
		pnts = append(pnts, make(point, 0, 2))
		v += rand.Float64()
		pnts[i] = append(pnts[i], v)                       // Append x
		pnts[i] = append(pnts[i], f(v)+rand.NormFloat64()) // Append y
	}

	return pnts
}

func (w polyCoefs) minAt(x0, a, tol float64) float64 {
	var d float64

	for {
		d = w.derivAt(x0)
		if math.Abs(d) < tol {
			return x0
		}

		x0 -= a * d
	}
}

func (w polyCoefs) derivAt(x float64) float64 {
	var d float64
	k := 1.0

	for i := 1; i < len(w); i++ {
		d += k * w[i] * math.Pow(x, k-1)
		k++
	}

	return d
}

func bucketVolume(rt, rb float64) float64 {
	rt2, rb2 := rt*rt, rb*rb
	return (rt2 - rt*rb + rb2) / 6.0 * math.Sqrt(math.Pow((1.0-math.Pi*rb2)/(rt+rb), 2)-4.0*math.Pow(math.Pi*(rt-rb), 2))
}

func gradBucketVolume(rt, rb, h float64) (float64, float64) {
	return (bucketVolume(rt+h, rb) - bucketVolume(rt-h, rb)) / (2 * h), (bucketVolume(rt, rb+h) - bucketVolume(rt, rb-h)) / (2 * h)
}

func findBucketMaxVol() (float64, float64, float64) {
	var rt0, rb0, rt1, rb1 float64
	var V0, V1 float64
	tol := 0.0001

	rt1, rb1 = 0.1, 0.1
	for {
		V0 = V1
		V1 = bucketVolume(rt1, rb1)
		if math.Abs(V1-V0) < tol {
			return rt1, rb1, V1
		}

		rt0, rb0 = gradBucketVolume(rt1, rb1, tol)
		rt1 += tol * rt0
		rb1 += tol * rb0
	}
}

func findMaxBoxVolume() float64 {
	var dx, dy float64
	var V0, V1 float64
	tol := 0.000001 // Smallest difference in volume
	r := 0.1        // Learning rate (0 < r < 1)

	x, y := 1.0, 1.0
	for {
		V0 = V1
		V1 = boxVolume(x, y)
		if math.Abs(V1-V0) < tol {
			return V1
		}

		dx, dy = gradBoxVolume(x, y)
		x += r * dx
		y += r * dy
	}
}

// boxVolume returns the volume of a box having surface area of 1 given two sides.
func boxVolume(x, y float64) float64 {
	if x < 0 || y < 0 {
		panic("boxVolume: sides must be non-negative")
	}

	xy := x * y
	return (xy - 2*xy*xy) / (2 * (x + y))
}

// gradBoxVolume returns dx and dy of a box having surface area of 1 given two sides.
func gradBoxVolume(x, y float64) (float64, float64) {
	xx := x * x
	xy := x * y
	yy := y * y
	div := x + y
	div *= 2 * div
	return yy * (1 - 2*xx - 4*xy) / div, xx * (1 - 2*xx - 4*xy) / div
}
