package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	N := readInt()
	X, Y := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		scanInts(&X[i], &Y[i])
	}

	ans := 0

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				x1, y1 := X[i], Y[i]
				x2, y2 := X[j], Y[j]
				x3, y3 := X[k], Y[k]
				if (x2-x1)*(y3-y1) != (x3-x1)*(y2-y1) {
					ans++
				}
			}
		}
	}

	fmt.Println(ans)
}

var stdin = bufio.NewReader(os.Stdin)

func readString() string {
	str, _ := stdin.ReadString('\n')
	return strings.TrimSpace(str)
}

func readStrings() []string {
	return strings.Split(readString(), " ")
}

func readInt() int {
	n, _ := strconv.Atoi(readString())
	return n
}

func readInts() []int {
	strs := readStrings()
	ints := make([]int, len(strs))
	for i, str := range strs {
		ints[i], _ = strconv.Atoi(str)
	}
	return ints
}

func scanInts(args ...*int) {
	strs := readStrings()
	for i, str := range strs {
		*args[i], _ = strconv.Atoi(str)
	}
}
