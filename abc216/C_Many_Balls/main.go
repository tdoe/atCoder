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
	N := readInt()

	answer := ""
	for N > 1 {
		if N%2 == 0 {
			N /= 2
			answer += "B"
		} else {
			N -= 1
			answer += "A"
		}
	}
	answer = "A" + reverseString(answer)

	if len(answer) <= 120 {
		fmt.Println(answer)
	} else {
		fmt.Println("???")
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

func descSort(a []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sortDigits(num int) int {
	digits := []rune(strconv.Itoa(num))
	sort.Slice(digits, func(i, j int) bool {
		return digits[i] < digits[j]
	})
	sortedNum, _ := strconv.Atoi(string(digits))
	return sortedNum
}

func descSortDigits(num int) int {
	digits := []rune(strconv.Itoa(num))
	sort.Slice(digits, func(i, j int) bool {
		return digits[i] > digits[j]
	})
	sortedNum, _ := strconv.Atoi(string(digits))
	return sortedNum
}
