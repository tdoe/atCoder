package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	N := readInt()

	t := make([]int, N+1)
	x := make([]int, N+1)
	y := make([]int, N+1)

	for i := 1; i < N+1; i++ {
		scanInts(&t[i], &x[i], &y[i])
	}

	for i := 1; i < N+1; i++ {
		currentPos := abs(x[i-1] + y[i-1])
		nextPos := abs(x[i] + y[i])
		posDiff := abs(nextPos - currentPos)
		tDiff := t[i] - t[i-1]
		if posDiff > tDiff {
			fmt.Println("No")
			return
		} else if (tDiff-posDiff)%2 == 1 {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")

}

var stdin = bufio.NewReader(os.Stdin)

func readString() string {
	str, _ := stdin.ReadString('\n')
	return strings.TrimSpace(str)
}

func readStrings() []string {
	return strings.Split(readString(), " ")
}

func scanInts(args ...*int) {
	strs := readStrings()
	for i, str := range strs {
		*args[i], _ = strconv.Atoi(str)
	}
}

func readInts() []int {
	strs := readStrings()
	ints := make([]int, len(strs))
	for i, str := range strs {
		ints[i], _ = strconv.Atoi(str)
	}
	return ints
}

func readInt() int {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	return n
}

func descSort(a []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
