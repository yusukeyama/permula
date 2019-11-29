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
func Combination(list int, target int) int {
	if list == 0 || target == 0 {
		return 0
	}
	answer := 0
	listArray := make([]int, target)
	for i := 0; i < target; i++ {
		listArray = append(listArray, list - i)
	}
	tmp := 0
	for i, v := range listArray {
		if i == 0 {
			tmp = v
		}
		tmp = tmp * v
	}
	targetArray := make([]int, target)
	for i := 0; i < target; i++ {
		targetArray = append(targetArray, target - 1)
	}
	stmp := 0
	for i, v := range targetArray {
		if i == 0 {
			stmp = v
		}
		stmp = stmp * v
	}
	answer = tmp / stmp
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