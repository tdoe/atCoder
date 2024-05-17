package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var X, Y int
	fmt.Scanf("%d.%d", &X, &Y)
	if Y <= 2 {
		fmt.Printf("%d-", X)
	} else if Y <= 6 {
		fmt.Printf("%d", X)
	} else {
		fmt.Printf("%d+", X)
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
