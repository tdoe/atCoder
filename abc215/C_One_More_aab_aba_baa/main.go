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
	inputs := readStrings()
	S := inputs[0]
	K, _ := strconv.Atoi(inputs[1])

	words := make(map[string]bool)
	permute([]rune(S), 0, words)

	answers := make([]string, 0, len(words))
	for word := range words {
		answers = append(answers, word)
	}
	sort.Strings(answers)
	fmt.Println(answers[K-1])
}

func swap(runes []rune, i, j int) {
	runes[i], runes[j] = runes[j], runes[i]
}

func permute(runes []rune, l int, words map[string]bool) {
	if l == len(runes)-1 {
		words[string(runes)] = true
		return
	}
	for i := l; i < len(runes); i++ {
		swap(runes, l, i)
		permute(runes, l+1, words)
		swap(runes, l, i)
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
