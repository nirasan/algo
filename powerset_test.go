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

func TestPowerset3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	result := [][]int{}
	Powerset3(nums, func(in []int){
		result = append(result, in)
	})
	if len(result) != int(math.Pow(2, float64(len(nums)))) {
		t.Errorf("invalid result: %v\n", result)
	} else {
		t.Logf("result: %v\n", result)
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
	b.Run("Powerset3", func(b *testing.B){
		f := func(in []int) {}
		for i := 0; i < b.N; i++ {
			Powerset3(nums, f)
		}
	})
}

func BenchmarkPowerset_len10(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
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
	b.Run("Powerset3", func(b *testing.B){
		f := func(in []int) {}
		for i := 0; i < b.N; i++ {
			Powerset3(nums, f)
		}
	})
}