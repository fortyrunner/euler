package main

import (
  "fmt"
  "math"
)



func main() {
  problem1()
  problem2()
  problem3()
}


func problem1() {
  fmt.Println("Project Euler - Problem 1. Find the sum of all the multiples of 3 or 5 below 1000")

  upper_bound := 1000
  sum := 0
  for i := 0; i < upper_bound; i++ {

    if i % 5 == 0 {
      sum += i
    } else if i % 3 == 0 {
      sum += i
    }
  }

  fmt.Printf("-->Answer - Sum of 3/5 multiples less than %d is %d\n\n", upper_bound, sum)

}


func problem2() {
  fmt.Println("Project Euler - Problem 2. By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.")

  current := 1
  next := 2
  sum := 2

  for next < 4000000 {
    //fmt.Printf("%d ", current)

    i := next
    next += current
    current = i

    if next % 2 == 0 {
      sum += next
    }

  }

  fmt.Printf("-->Answer sum of even values is %d\n\n", sum)

}

func problem3() {

  fmt.Println("Project Euler - Problem 3. What is the largest prime factor of the number 600851475143")

  number := 600851475143

  upper_bound := int(math.Sqrt(float64(number))) + 1

  fmt.Printf("The upper bound for factors is %d\n", upper_bound)

  for i := upper_bound; i > 2; i-- {

    if isPrime(i) {

      if number % i == 0 {
        fmt.Printf("-->Answer Largest prime factor of %d is %d\n\n", number, i)

        break
      }
    }
  }
}

func isPrime(n int) bool {

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