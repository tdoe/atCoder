package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var N, M, ans int
	scanIntVariables(&N, &M)
	A := readInts()
	B := readInts()
	sort.Ints(A)
	sort.Ints(B)
	ans = math.MaxInt
	var ai, bi int
	for ai < N && bi < M {
		ans = int(math.Min(float64(ans), math.Abs(float64(A[ai]-B[bi]))))
		if A[ai] < B[bi] {
			ai++
		} else {
			bi++
		}
	}
	fmt.Println(ans)
}

var stdin = bufio.NewReader(os.Stdin)

func readString() string {
	line, err := stdin.ReadString('\n')
	if err != nil {
		panic("Failed to read stdin: " + err.Error())
	}
	return strings.TrimSpace(line)
}

func readStrings() []string {
	str := readString()
	return strings.Split(str, " ")
}

func readInt() int {
	str := readString()
	num, _ := strconv.Atoi(str)
	return num
}

func readInts() []int {
	strs := readStrings()
	ints := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		ints[i], _ = strconv.Atoi(strs[i])
	}
	return ints
}

func scanIntVariables(args ...*int) {
	var scanned int
	for len(args) > scanned {
		inputSlice := readInts()
		n := len(inputSlice)
		if n == 0 || n > len(args)-scanned {
			fmt.Printf("n %d len(args) %d scanned.\n", n, len(args))
			fmt.Print("--- Something wrong in scanIntVariables ---")
			return
		}
		for i := 0; i < n; i++ {
			*args[i+scanned] = inputSlice[i]
		}
		scanned += n
	}
}

func readStringInt() (string, int) {
	str := readString()
	sp := strings.Split(str, " ")
	i, _ := strconv.Atoi(sp[1])
	return sp[0], i
}

func map2string(m map[string]struct{}, isSort bool) []string {
	var results []string
	for key := range m {
		results = append(results, key)
	}
	if isSort {
		sort.Strings(results)
	}
	return results
}
