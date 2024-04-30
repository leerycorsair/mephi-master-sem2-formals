package main

import (
	"math"
	"testing"

	"pgregory.net/rapid"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func TestPropProd(t *testing.T) {
	gen := rapid.Custom(func(t *rapid.T) Matrix {
		size := rapid.Uint16Range(1, 10).Draw(t, "mtr")
		mtr := make(Matrix, size)
		for i := 0; i < len(mtr); i++ {
			mtr[i] = make([]float64, size)
			for j := 0; j < len(mtr[i]); j++ {
				mtr[i][j] = roundFloat(rapid.Float64Range(-10, 10).Draw(t, "mtr value"), 9)
			}
		}
		return mtr
	})
	cmp1 := func(a, b float64) bool {
		return math.Abs(a-b) < PRECISION
	}
	cmp2 := func(a, b int) bool {
		return a == b
	}

	property := PropProd(Determinant, Rank, cmp1, cmp2, gen)
	rapid.Check(t, property)
}
