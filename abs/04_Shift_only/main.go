package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/abs/tasks/abc081_b
func main() {
	_ = readInt()
	Digits := readInts()
	ans := div(Digits, 0)
	fmt.Println(ans)
}

func div(digits []int, cnt int) int {
	ints := make([]int, len(digits))
	for i, a := range digits {
		if a%2 == 0 {
			ints[i] = a / 2
		} else {
			return cnt
		}
	}
	return div(ints, cnt+1)
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
