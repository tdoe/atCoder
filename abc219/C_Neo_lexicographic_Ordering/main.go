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
	X := readString()
	N := readInt()
	S := make([]string, N)
	for i := 0; i < N; i++ {
		S[i] = readString()
	}

	al := "abcdefghijklmnopqrstuvwxyz"
	orgMap := make(map[string]string)
	for i, x := range X {
		orgMap[string(x)] = string(al[i])
	}

	convMap := make(map[string]string)
	for key, val := range orgMap {
		convMap[val] = key
	}

	NewS := make([]string, N)
	for i := 0; i < N; i++ {
		for _, n := range S[i] {
			NewS[i] += orgMap[string(n)]
		}
	}

	sort.Strings(NewS)

	for _, s := range NewS {
		answer := ""
		for _, c := range s {
			answer += convMap[string(c)]
		}
		fmt.Println(answer)
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
