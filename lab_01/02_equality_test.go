package main

import (
	"testing"

	"pgregory.net/rapid"
)

func TestLinearAndBinarySearch(t *testing.T) {
	property := PropCmpSearches(LinearSearch, BinarySearch)
	rapid.Check(t, property)
}

func TestInterpolationAndBinarySearch(t *testing.T) {
	property := PropCmpSearches(InterpolationSearch, BinarySearch)
	rapid.Check(t, property)
}

func TestInterpolationAndLinearSearch(t *testing.T) {
	property := PropCmpSearches(InterpolationSearch, LinearSearch)
	rapid.Check(t, property)
}

func TestBubbleAndInsertSort(t *testing.T) {
	property := PropCmpSorts(BubbleSort, InsertSort)
	rapid.Check(t, property)
}
