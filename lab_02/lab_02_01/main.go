package main

import (
	"errors"
	"math"
)

type Matrix [][]float64

func Determinant(mtr Matrix) float64 {
	if len(mtr) != len(mtr[0]) {
		panic(errors.New("Matrix is not square"))
	}

	if len(mtr) == 1 {
		return mtr[0][0]
	}

	var det float64
	for i, coeff := range mtr[0] {
		minor := minorMtr(mtr, 0, i)
		det += coeff * math.Pow(-1, float64(i)) * Determinant(minor)
	}
	return det
}

func minorMtr(mtr Matrix, row, col int) Matrix {
	size := len(mtr)
	minor := make(Matrix, size-1)
	for i := range minor {
		minor[i] = make([]float64, size-1)
	}

	for i := 0; i < size; i++ {
		if i == row {
			continue
		}
		for j := 0; j < size; j++ {
			if j == col {
				continue
			}
			minorRow := i
			if i > row {
				minorRow--
			}
			minorCol := j
			if j > col {
				minorCol--
			}
			minor[minorRow][minorCol] = mtr[i][j]
		}
	}
	return minor
}

func Rank(mtr Matrix) int {
	var tmpMtr Matrix = copyMatrix(mtr)
	rows := len(tmpMtr)
	if rows == 0 {
		return 0
	}
	cols := len(tmpMtr[0])
	rank := 0
	for i := 0; i < rows; i++ {
		if allZero(tmpMtr[i]) {
			continue
		}
		pivotRow := i
		pivotCol := 0
		for pivotCol < cols && tmpMtr[pivotRow][pivotCol] == 0 {
			pivotCol++
		}
		if pivotCol == cols {
			continue
		}
		tmpMtr[pivotRow], tmpMtr[i] = tmpMtr[i], tmpMtr[pivotRow]
		for j := 0; j < rows; j++ {
			if j != i && tmpMtr[j][pivotCol] != 0 {
				scale := tmpMtr[j][pivotCol] / tmpMtr[i][pivotCol]
				for k := 0; k < cols; k++ {
					tmpMtr[j][k] -= scale * tmpMtr[i][k]
				}
			}
		}
		rank++
	}
	return rank
}

func allZero(row []float64) bool {
	for _, val := range row {
		if val < PRECISION {
			return false
		}
	}
	return true
}

func copyMatrix(src Matrix) Matrix {
	dst := make(Matrix, len(src))
	for i := range dst {
		dst[i] = make([]float64, len(src[i]))
	}
	for i := range src {
		copy(dst[i], src[i])
	}

	return dst
}
