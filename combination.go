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
		bi = Combination2NextIndex(bi)
		flags := bi
		s := []int{}
		for _, n := range nums {
			if flags%2 != 0 {
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
	return Fact(n) / (Fact(n-m) * Fact(m))
}

func Fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * Fact(n-1)
	}
}

func Combination2NextIndex(n int) int {
	// (1) 2 進数で 1 になっている一番下の桁を取得
	smallest := n & -n
	// (2) 1 になっている一番下の桁に 1 を足す
	ripple := n + smallest
	// (3) 新たに 1 になっている一番下の桁を取得
	newSmallest := ripple & -ripple
	// (3) / (1) で (2) の操作により増えた 0 の数（失われた 1 の数）が算出され
	// - 1 することで補填するべき 1 のビット列を取得する
	ones := ((newSmallest / smallest) >> 1) - 1
	// ripple に ones を補填して新しいインデックスとする
	return ripple | ones
}
