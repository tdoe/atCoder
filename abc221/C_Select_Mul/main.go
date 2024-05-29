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
	A := readInts()
	X := readInt()

	Asum := 0
	for _, a := range A {
		Asum += a
	}

	shou := X / Asum
	answer := N * shou
	Bsum := shou * Asum

	for _, i := range A {
		Bsum += i
		answer++
		if Bsum > X {
			fmt.Println(answer)
			return
		}
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
