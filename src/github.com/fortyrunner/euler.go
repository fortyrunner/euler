package main

import (
    "fmt"
    "math"
    "strconv"

    "time"
)


func main() {

    //loadsaprimes()

    problem1()
    problem2()
    problem3()
    problem4()
    problem5()
    problem6();
    problem7();

}

func loadsaprimes() {

    value := 100000000

    start := time.Now()

    count := 1
    for i := 0; i < value; i++ {
        if IsPrime(i) {
            //fmt.Printf("%d %d\n", i, count)
            count++
        }
    }

    fmt.Printf("%s\n", time.Since(start))
}



func problem1() {
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



func problem2() {

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

        if next % 2 == 0 {
            answer += next
        }

    }

    fmt.Printf("-->Answer sum of even values is %d\n", answer)
    fmt.Printf("took=%s\n\n", time.Since(start))
}

func problem3() {

    fmt.Println("Problem 3. What is the largest prime factor of the number 600851475143")
    start := time.Now()

    answer := 600851475143

    upperBound := int(math.Sqrt(float64(answer))) + 1

    fmt.Printf("The upper bound for factors is %d\n", upperBound)

    for i := upperBound; i > 2; i-- {

        if IsPrime(i) {

            if answer % i == 0 {
                fmt.Printf("-->Answer: Largest prime factor of %d is %d\n", answer, i)

                break
            }
        }
    }

    fmt.Printf("took=%s\n\n", time.Since(start))
}

func problem4() {

    fmt.Println("Problem 4. Find the largest palindrome made from the product of two 3-digit numbers")
    start := time.Now()

    answer := 0;
    for i := 999; i >= 100; i-- {
        for j := 999; j >= 100; j-- {
            product := i * j

            if IsPalindrome(product) {
                if product > answer {
                    answer = product
                }
            }
        }
    }

    fmt.Printf("-->Answer: The maximum palindrome of two 3 digit numbers is %d\n", answer)
    fmt.Printf("took=%s\n\n", time.Since(start))
}

func problem5() {

    fmt.Println("Problem 5. What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?")
    start := time.Now()


    // Upper bound is product of all values <= 20
    upper_bound := 1;
    for i := 1; i <= 20; i++ {
        upper_bound = upper_bound * i
    }

    guess := 20
    answer := 0
    for guess < upper_bound {

        // We don't need to do values below 11 because they are all divisors
        // of the values above 10

        if guess % 20 == 0 && guess % 19 == 0 && guess % 18 == 0 &&
        guess % 17 == 0 && guess % 16 == 0 && guess % 15 == 0 &&
        guess % 14 == 0 && guess % 13 == 0 && guess % 12 == 0 &&
        guess % 11 == 0 {
            answer = guess;
            break;
        }

        guess++
    }

    fmt.Printf("-->Answer: The smallest number is %d\n", answer)
    fmt.Printf("took=%s\n\n", time.Since(start))
}

func problem6() {
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

func problem7() {

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
        if IsPrime(possibility) {
            count = count + 1;
            answer = possibility
        }
        possibility = possibility + 2
    }


    fmt.Printf("-->Answer: The 10,001st prime is %d\n", answer)
    fmt.Printf("took=%s\n\n", time.Since(start))
}


func IsPalindrome(p int) bool {

    s := strconv.Itoa(p)

    midpoint := len(s) / 2

    for i := 0; i < midpoint; i++ {
        if s[i] != s[len(s) - i - 1] {
            return false
        }
    }
    return true
}

func IsPrime(n int) bool {

    // Test the shortcuts first
    if n <= 1 {
        return false
    }

    if n <= 3 {
        return true
    }

    if n % 2 == 0 || n % 3 == 0 {
        return false
    }



    i := 5
    for ( i * i <= n) {
        if n % i == 0 || n % (i + 2) == 0 {
            return false
        }
        i = i + 6
    }
    return true
}

