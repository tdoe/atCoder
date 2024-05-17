package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tmp := readInts()
	S, T := tmp[0], tmp[1]

	count := 0
	for a := 0; a <= S; a++ {
		for b := 0; b <= S; b++ {
			for c := 0; c <= S; c++ {
				if a+b+c <= S && a*b*c <= T {
					count++
				}
			}
		}
	}

	fmt.Println(count)
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
