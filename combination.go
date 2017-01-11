package main

func Combination1(nums []int, n int) [][]int {
	ps := Powerset2(nums)
	result := make([][]int, CombinationCount(len(nums), n))
	index := 0
	for _, s := range ps {
		if len(s) == n {
			result[index] = s
			index++
		}
	}
	return result
}

func Combination2(nums []int, n int) [][]int {
	size := len(nums)
	result := make([][]int, CombinationCount(len(nums), n))
	bi := (1 << uint(n)) - 1
	max := 1 << uint(size)
  index := 0
	for bi < max {
		bi = NextIndex(bi)
		flags := bi
		s := []int{}
		for _, n := range nums {
			if flags % 2 != 0 {
				s = append(s, n)
			}
			flags /= 2
		}
		result[index] = s
		index++
	}
	return result
}

func CombinationCount(n, m int) int {
	return Fact(n) / ( Fact(n - m) * Fact(m) )
}

func Fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * Fact(n - 1)
	}
}

func NextIndex(n int) int {
	smallest := n & -n
	ripple := n + smallest
	newSmallest := ripple & -ripple
	ones := ((newSmallest / smallest) >> 1) - 1
	return ripple | ones
}

