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

func Powerset2(nums []int) [][]int {
	length := int(math.Pow(2, float64(len(nums))))
	result := make([][]int, length)

	index := 0
	result[index] = []int{}
	index++

	for _, n := range nums {
		max := index
		for i := 0; i < max; i++ {
			result[index] = copyAndAppend(result[i], n)
			index++
		}
	}

	return result
}

func copyAndAppend(nums []int, n int) []int {
	dst := make([]int, len(nums)+1)
	copy(dst, nums)
	dst[len(nums)] = n
	return dst
}

func Powerset3(nums []int, f func([]int)) {
	if len(nums) == 0 {
		f([]int{})
	}

	length := int(math.Pow(2, float64(len(nums))))

	for i := 0; i < length; i++ {
		bi := i
		s := []int{}
		for _, n := range nums {
			if bi % 2 != 0 {
				s = append(s, n)
			}
			bi = bi / 2

		}
		f(s)
	}
}