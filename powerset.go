package main

import "math"

func Powerset1(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{ []int{} }
	}

	length := int(math.Pow(2, float64(len(nums))))
	result := make([][]int, length)

	for i := 0; i < length; i++ {
		bi := i
		s := []int{}
		for _, n := range nums {
			if bi % 2 != 0 {
				s = append(s, n)
			}
			bi = bi / 2

		}
		result[i] = s
	}

	return result
}
