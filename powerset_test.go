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

func BenchmarkPowerset_len3(b *testing.B) {
	nums := []int{1, 2, 3}
	b.Run("Powerset1", func(b *testing.B){
		for i := 0; i < b.N; i++ {
			Powerset1(nums)
		}
	})
	b.Run("Powerset2", func(b *testing.B){
		for i := 0; i < b.N; i++ {
			Powerset2(nums)
		}
	})
}
