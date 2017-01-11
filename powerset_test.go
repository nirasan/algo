package main

import (
	"testing"
	"math"
)

func TestPowerset(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	samples := []func([]int)[][]int{ Powerset1, Powerset2 }
	for i, sample := range samples {
		result := sample(nums)
		if len(result) != int(math.Pow(2, float64(len(nums)))) {
			t.Errorf("[%d] invalid result: %v\n", i, result)
		} else {
			t.Logf("[%d] result: %v\n", i, result)
		}
	}
}
