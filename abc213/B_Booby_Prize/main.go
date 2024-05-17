package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Booby struct {
	num   int
	index int
}

func main() {
	var N int
	scanInts(&N)

	ints := readInts()
	maxInt := Booby{
		num:   0,
		index: 0,
	}
	boobyInt := Booby{
		num:   0,
		index: 0,
	}
	for i, a := range ints {
		if a > maxInt.num {
			boobyInt = maxInt
			maxInt.num = a
			maxInt.index = i + 1
		}
	}
	fmt.Println(boobyInt.index)
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
