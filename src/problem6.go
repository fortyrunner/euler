package main

import (
    "fmt"
    "time"
    "math"
)

func main() {
    // Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

    fmt.Println("Problem 6. What is the difference between the sum of the squares of the first 100 natural numbers and the square of the sum ?")
    start := time.Now()

    answer := 0.0

    sum_of_squares := math.Pow(100 / 2 * (100 + 1), 2)

    fmt.Printf("-->Answer: The sum of squares is %10.0f\n", sum_of_squares)

    for i := 1.0; i <= 100.0; i++ {
        answer = answer + i * i
    }

    answer = math.Abs(answer - sum_of_squares);

    fmt.Printf("-->Answer: The difference number is %10.0f\n", answer)
    fmt.Printf("took=%s\n\n", time.Since(start))
}

