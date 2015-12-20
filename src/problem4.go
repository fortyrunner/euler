package main


import (
    "fmt"
    "time"
    "euler"
)


func main() {

    fmt.Println("Problem 4. Find the largest palindrome made from the product of two 3-digit numbers")
    start := time.Now()

    answer := 0;
    for i := 999; i >= 100; i-- {
        for j := 999; j >= 100; j-- {
            product := i * j

            if euler.IsPalindrome(product) {
                if product > answer {
                    answer = product
                }
            }
        }
    }

    fmt.Printf("-->Answer: The maximum palindrome of two 3 digit numbers is %d\n", answer)
    fmt.Printf("took=%s\n\n", time.Since(start))
}
