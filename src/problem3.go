package main

import (
    "fmt"
    "time"
    "math"
    "euler"
)

func main() {

    fmt.Println("Problem 3. What is the largest prime factor of the number 600851475143")
    start := time.Now()

    answer := 600851475143

    upperBound := int(math.Sqrt(float64(answer))) + 1

    fmt.Printf("The upper bound for factors is %d\n", upperBound)

    for i := upperBound; i > 2; i-- {

        if euler.IsPrime(i) {

            if answer % i == 0 {
                fmt.Printf("-->Answer: Largest prime factor of %d is %d\n", answer, i)

                break
            }
        }
    }

    fmt.Printf("took=%s\n\n", time.Since(start))
}
