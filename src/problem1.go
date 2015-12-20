package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Problem 1. Find the sum of all the multiples of 3 or 5 below 1000")
    start := time.Now()
    upperBound := 1000
    answer := 0
    for i := 0; i < upperBound; i++ {

        if i % 5 == 0 {
            answer += i
        } else if i % 3 == 0 {
            answer += i
        }
    }

    fmt.Printf("-->Answer - Sum of 3/5 multiples less than %d is %d\n", upperBound, answer)
    fmt.Printf("took=%s\n\n", time.Since(start))
}
