package permula

import (
	"bufio"
	"strconv"
)

// ScanToSlicebyInt get bufio.Scanner split white space.
// return slice list.
func ScanToSlicebyInt(sc *bufio.Scanner, count int) []int {
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
