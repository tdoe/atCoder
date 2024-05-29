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
	N := readString()
	var tmp, ans int

	for i := 0; i < (1 << len(N)); i++ {
		var A, B string
		for j := 0; j < len(N); j++ {
			if i>>j&1 == 1 {
				A += N[j : j+1]
			} else {
				B += N[j : j+1]
			}
		}

		fmt.Println(A, B)
		if A == "" || B == "" {
			continue
		}

		A = reverseStr(sortStr(A))
		B = reverseStr(sortStr(B))

		if A[0] == '0' || B[0] == '0' {
			continue
		}

		tmpA, _ := strconv.Atoi(A)
		tmpB, _ := strconv.Atoi(B)
		tmp = tmpA * tmpB

		if tmp >= ans {
			ans = tmp
		}
	}

	fmt.Println(ans)
}

func sortStr(w string) string {
	s := strings.Split(w, "")  //文字列を1文字ずつに分解する
	sort.Strings(s)            //文字列をソートする
	return strings.Join(s, "") //ソートされた文字列を結合してreturn
}

// 文字列を反転させる関数
func reverseStr(s string) string {
	runes := []rune(s) //文字列をルーンのスライスに変換
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] //文字を反転させる
	}
	return string(runes) //反転された文字列をreturn
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
