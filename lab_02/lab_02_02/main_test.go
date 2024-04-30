package main

import (
	"testing"

	"pgregory.net/rapid"
)

func TestPropSum(t *testing.T) {
	gen1 := rapid.Custom(func(t *rapid.T) HashTableT {
		ht := rapid.MapOfN(rapid.StringMatching(`[a-z]+`), rapid.IntRange(1, 10), 1, 100).Draw(t, "hashTable")
		return HashTableT(ht)
	})
	gen2 := rapid.Custom(func(t *rapid.T) CustomErrorT {
		ce := rapid.StringMatching(`[a-z]+`).Draw(t, "customError")
		return CustomErrorT(ce)
	})
	resCmp := func(a, b ResultT) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	property := PropSum(HashTableT.ParseHash,
		CustomErrorT.ParseError, gen1, gen2, resCmp)
	rapid.Check(t, property)
}

func TestPropSum2(t *testing.T) {
	gen1 := rapid.Custom(func(t *rapid.T) HashTableT {
		ht := rapid.MapOfN(rapid.StringMatching(`[a-z]+`), rapid.IntRange(1, 10), 1, 100).Draw(t, "hashTable")
		return HashTableT(ht)
	})
	gen2 := rapid.Custom(func(t *rapid.T) CustomErrorT {
		ce := rapid.StringMatching(`[a-z]+`).Draw(t, "customError")
		return CustomErrorT(ce)
	})
	resCmp := func(a, b ResultT) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	property := PropSum2(HashTableT.ParseHash,
		CustomErrorT.ParseError, gen1, gen2, resCmp)
	rapid.Check(t, property)
}
