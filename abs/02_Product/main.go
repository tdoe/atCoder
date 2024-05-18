package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://atcoder.jp/contests/abs/tasks/abc086_a
func main() {
	var a, b int
	scanInts(&a, &b)

	mul := a * b
	if mul%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
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
