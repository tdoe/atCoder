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
	var H, W, N int
	scanInts(&H, &W, &N)

	ALine := make([]int, N)
	BLine := make([]int, N)
	for i := 0; i < N; i++ {
		scanInts(&ALine[i], &BLine[i])
	}

	ALine = compress(ALine)
	BLine = compress(BLine)

	for i := 0; i < N; i++ {
		fmt.Println(ALine[i], BLine[i])
	}
}

// 座標圧縮
func compress(ints []int) []int {
	mp := make(map[int]int)
	keys := make([]int, len(ints))
	for i, v := range ints {
		mp[v] = 0
		keys[i] = v
	}

	sort.Ints(keys)

	val := 1
	for _, k := range keys {
		mp[k] = val
		val++
	}

	for i := 0; i < len(ints); i++ {
		ints[i] = mp[ints[i]]
	}

	return ints
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
