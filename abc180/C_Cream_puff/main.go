package main

import (
	"fmt"
	"sort"
)

func main() {
	N := readInt()
	var ans []int
	for i := 1; i*i <= N; i++ {
		if N%i == 0 {
			ans = append(ans, i)
			if i != N/i {
				ans = append(ans, N/i)
			}
		}
	}

	sort.Ints(ans)

	for _, a := range ans {
		fmt.Println(a)
	}
}

func readInt() int {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	return n
}
