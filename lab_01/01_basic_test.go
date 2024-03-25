package main

import (
	"sort"
	"testing"

	"pgregory.net/rapid"
)

func TestBubbleSort(t *testing.T) {
	property := PropCmpSorts(BubbleSort, sort.Ints)
	rapid.Check(t, property)
}
