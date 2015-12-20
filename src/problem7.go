package main

import (
    "fmt"
    "time"
    "euler"
)

func main(){

    fmt.Println("Problem 7. What is the 10001st Prime?")
    start := time.Now()

    // Skip the number 2 so that we can use an increment of 2 and halve the set
    // I know that this is brute force but..
    // 1. It takes 7.1ms on my machine.
    // 2. I would never calculate these primes each time if I needed them. I would pre-calculate a set up front

    possibility := 3;
    answer := 3
    count := 1

    for count < 10001 {
        if euler.IsPrime(possibility) {
            count = count + 1;
            answer = possibility
        }
        possibility = possibility + 2
    }


    fmt.Printf("-->Answer: The 10,001st prime is %d\n", answer)
    fmt.Printf("took=%s\n\n", time.Since(start))

}
