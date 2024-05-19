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
	var N, Y int
	scanInts(&N, &Y)

	for x := 0; x <= N; x++ {
		for y := 0; y <= N-x; y++ {
			z := N - x - y
			if 10000*x+5000*y+1000*z == Y {
				fmt.Println(x, y, z)
				return
			}
		}
	}
	fmt.Println("-1 -1 -1")
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
