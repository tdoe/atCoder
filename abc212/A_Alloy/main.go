package main

import "fmt"

func main() {
	var A, B int
	_, _ = fmt.Scanf("%d %d", &A, &B)

	if A == 0 {
		fmt.Println("Silver")
	} else if B == 0 {
		fmt.Println("Gold")
	} else {
		fmt.Println("Alloy")
	}
}
