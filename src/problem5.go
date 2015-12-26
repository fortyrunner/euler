package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Problem 5. What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?")
	start := time.Now()

	// Upper bound is product of all values <= 20
	upper_bound := 1
	for i := 1; i <= 20; i++ {
		upper_bound = upper_bound * i
	}

	guess := 20
	answer := 0
	for guess < upper_bound {

		// We don't need to do values below 11 because they are all divisors
		// of the values above 10

		if guess%20 == 0 && guess%19 == 0 && guess%18 == 0 &&
			guess%17 == 0 && guess%16 == 0 && guess%15 == 0 &&
			guess%14 == 0 && guess%13 == 0 && guess%12 == 0 &&
			guess%11 == 0 {
			answer = guess
			break
		}

		guess++
	}

	fmt.Printf("-->Answer: The smallest number is %d\n", answer)
	fmt.Printf("took=%s\n\n", time.Since(start))
}
