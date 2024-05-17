package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	N := readInt()
	if N == 1 {
		fmt.Println(0)
		return
	}

	for i := 2; i <= N; i++ {
		if int(math.Pow(2, float64(i))) > N {
			fmt.Println(i - 1)
			break
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
