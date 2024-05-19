package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/abs/tasks/abc083_b
func main() {
	var N, A, B int
	scanInts(&N, &A, &B)

	ans := 0
	for i := 1; i <= N; i++ {
		ss := strconv.Itoa(i)
		sum := 0
		for _, s := range ss {
			d := int(s - '0')
			sum += d
		}
		if sum >= A && sum <= B {
			ans += i
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
