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
	var N, M int
	scanInts(&N, &M)

	var A [][]string
	for i := 0; i < N*2; i++ {
		gcp := make([]string, M)
		hands := []rune(readString())
		for i, hand := range hands {
			gcp[i] = string(hand)
		}

		A = append(A, gcp)
	}

	var wins [][]int
	for i := 0; i < N*2; i++ {
		wins = append(wins, []int{0, i})
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			pA := wins[2*j][1]
			pB := wins[2*j+1][1]

			AHand := A[pA][i]
			BHand := A[pB][i]

			APoint, BPoint := match(AHand, BHand)

			wins[2*j][0] -= APoint
			wins[2*j+1][0] -= BPoint
		}

		sort.Slice(wins, func(i, j int) bool {
			if wins[i][0] == wins[j][0] {
				return wins[i][1] < wins[j][1]
			}
			return wins[i][0] < wins[j][0]
		})
	}

	for _, win := range wins {
		fmt.Println(win[1] + 1)
	}
}

func match(AHand, BHand string) (int, int) {
	if AHand == "G" && BHand == "C" {
		return 1, 0
	} else if AHand == "G" && BHand == "P" {
		return 0, 1
	} else if AHand == "C" && BHand == "P" {
		return 1, 0
	} else if AHand == "C" && BHand == "G" {
		return 0, 1
	} else if AHand == "P" && BHand == "G" {
		return 1, 0
	} else if AHand == "P" && BHand == "C" {
		return 0, 1
	} else {
		return 0, 0
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
