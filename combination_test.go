package main

import (
	"testing"
)

func TestCombination(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	samples := []func([]int, int)[][]int{ Combination1, Combination2 }
	for i, sample := range samples {
		result := sample(nums, 3)
		if len(result) != 10 {
			t.Errorf("[%d] invalid result: %v\n", i, result)
		} else {
			t.Logf("[%d] result: %v\n", i, result)
		}
	}
}

func TestFact(t *testing.T) {
	if f := Fact(5); f != 120 {
		t.Errorf("error: %v", f)
	}
}

func TestCombinationCount(t *testing.T) {
	if c := CombinationCount(5, 3); c != 10 {
		t.Errorf("error: %v", c)
	}
}

func TestNextIndex(t *testing.T) {
	max := 1 << 5
	n := (1 << 3) - 1
	t.Logf("%b", n)
	for n < max {
		n = NextIndex(n)
		t.Logf("%b", n)
	}
}

func BenchmarkCombination_5to3(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5}
	b.Run("Combination1", func(b *testing.B){
		for i := 0; i < b.N; i++ {
			Combination1(nums, 3)
		}
	})
	b.Run("Combination2", func(b *testing.B){
		for i := 0; i < b.N; i++ {
			Combination2(nums, 3)
		}
	})
}

func BenchmarkCombination_10to3(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.Run("Combination1", func(b *testing.B){
		for i := 0; i < b.N; i++ {
			Combination1(nums, 3)
		}
	})
	b.Run("Combination2", func(b *testing.B){
		for i := 0; i < b.N; i++ {
			Combination2(nums, 3)
		}
	})
}