package main

import (
	"fmt"
)

func main() {
	var X string
	_, _ = fmt.Scanf("%s", &X)

	same := true
	step := true
	for i := 0; i < 3; i++ {
		if X[i] != X[i+1] {
			same = false
		}
		a := X[i] - '0'
		b := X[i+1] - '0'
		if (a+1)%10 != b {
			step = false
		}
	}

	if same || step {
		fmt.Println("Weak")
	} else {
		fmt.Println("Strong")
	}
}
