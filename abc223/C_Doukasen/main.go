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
	A := make([]float64, N)
	B := make([]float64, N)
	var fullTime float64
	for i := 0; i < N; i++ {
		var a, b int
		scanInts(&a, &b)
		A[i] = float64(a)
		B[i] = float64(b)
		fullTime += A[i] / B[i]
	}

	time := fullTime / 2

	var dist float64
	for i := 0; i < N; i++ {
		nt := A[i] / B[i]
		if nt > time {
			dist += B[i] * time
			break
		}
		time -= A[i] / B[i]
		dist += A[i]
	}

	fmt.Println(dist)
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
