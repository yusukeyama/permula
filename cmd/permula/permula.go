package permula

import (
	"bufio"
	"strconv"
)

// ScanToSlicebyInt get bufio.Scanner split white space.
// return slice list.
func ScanToSlicebyInt(sc *bufio.Scanner, count int) []int {
	if sc == nil || count == 0 {
		return nil
	}
	var list []int
	for i := 0; i < count; i++ {
		sc.Scan()
		i, e := strconv.Atoi(sc.Text())
		if e != nil {
			panic(e)
		}
		list = append(list, i)
	}
	return list
}

// Combination is sort by combination.
func Combination(n int, r int) int {
	if n == 0 || r == 0 {
		return 0
	}
	if a := n - r; a < 0 {
		return 0
	} 
	danswer := 1
	for i := n; i > 1; i-- {
		danswer *= i
	}
	manswer := 1
	for j := r; j > 1; j-- {
		manswer *= j
	}
	answer := danswer / manswer
	return answer
}

// Permutation is sort by permutation.
func Permutation(n int, r int) int {
	if n == 0 || r == 0 {
		return 0
	}
	if n < 0 || r < 0 {
		return 0
	}
	underCount := n - r
	if underCount < 0 {
		return 0
	}
	// denominator
	danswer := 1
	for i := underCount; i > 1; i-- {
		danswer *= i
	}
	// molecule
	manswer := 1
	for j := n; j > 1; j-- {
		manswer *= j
	}
	answer := manswer / danswer
	return answer
}