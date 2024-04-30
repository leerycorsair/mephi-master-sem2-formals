package main

import (
	"fmt"
	"reflect"
	"sort"
)

type HashTableT map[string]int
type CustomErrorT string

type ResultT []string

func (e CustomErrorT) ParseError() ResultT {
	return ResultT{}
}

type byValue struct {
	keys   ResultT
	values HashTableT
}

func (bv byValue) Len() int {
	return len(bv.keys)
}

func (bv byValue) Less(i, j int) bool {
	if bv.values[bv.keys[i]] < bv.values[bv.keys[j]] {
		return true
	} else if bv.values[bv.keys[i]] == bv.values[bv.keys[j]] {
		return bv.keys[i] < bv.keys[j]
	} else {
		return false
	}
}

func (bv byValue) Swap(i, j int) {
	bv.keys[i], bv.keys[j] = bv.keys[j], bv.keys[i]
}

func (t HashTableT) ParseHash() ResultT {
	keys := make(ResultT, 0, len(t))
	for key := range t {
		keys = append(keys, key)
	}
	bv := byValue{keys: keys, values: t}
	sort.Sort(bv)
	return bv.keys
}

func leftFunc(l string) int {
	return 20
}

func rightFunc(r string) int {
	return 10
}

func main() {
	e := Right2[string]{"ew"}
	r := CaseFunction2(e, leftFunc, rightFunc)
	fmt.Println(r)
	fmt.Println(reflect.TypeOf(e))

	f := FakeType[string]{"ew"}
	r = CaseFunction2(f, leftFunc, rightFunc)
	fmt.Println(r)
	fmt.Println(reflect.TypeOf(e))
}
