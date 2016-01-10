package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Problem 9. ")
	start := time.Now()

	answer := 0

	a := 1
	b := 1
	c := 1

	if !isMatch(a, b, c) {

	}


	fmt.Printf("-->Answer: The answer is %d\n", answer)
	fmt.Printf("took=%s\n\n", time.Since(start))

}

func isMatch(a int, b int, c int) bool {
	return a + b + c == 1000
}
