package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Problem 2. By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.")
	start := time.Now()

	current := 1
	next := 2
	answer := 2

	for next < 4000000 {
		//fmt.Printf("%d ", current)

		i := next
		next += current
		current = i

		if next%2 == 0 {
			answer += next
		}

	}

	fmt.Printf("-->Answer sum of even values is %d\n", answer)
	fmt.Printf("took=%s\n\n", time.Since(start))
}
