package main

import "testing"

func TestPowerset(t *testing.T) {
	nums := []int{1, 2, 3}
	samples := []func([]int)[][]int{ Powerset1 }
	for i, sample := range samples {
		result := sample(nums)
		if len(result) != 8 {
			t.Errorf("[%d] invalid result: %v\n", i, result)
		} else {
			t.Logf("[%d] result: %v\n", i, result)
		}
	}
}
