package main

import (
  "fmt"
  "math"
  "strconv"
)




func main() {
  problem1()
  problem2()
  problem3()
  problem4()
}


func problem1() {
  fmt.Println("Project Euler - Problem 1. Find the sum of all the multiples of 3 or 5 below 1000")

  upperBound := 1000
  sum := 0
  for i := 0; i < upperBound; i++ {

    if i % 5 == 0 {
      sum += i
    } else if i % 3 == 0 {
      sum += i
    }
  }

  fmt.Printf("-->Answer - Sum of 3/5 multiples less than %d is %d\n\n", upperBound, sum)

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

  upperBound := int(math.Sqrt(float64(number))) + 1

  fmt.Printf("The upper bound for factors is %d\n", upperBound)

  for i := upperBound; i > 2; i-- {

    if isPrime(i) {

      if number % i == 0 {
        fmt.Printf("-->Answer: Largest prime factor of %d is %d\n\n", number, i)

        break
      }
    }
  }
}

func problem4() {

  fmt.Println("Project Euler - Problem 4. Find the largest palindrome made from the product of two 3-digit numbers")

  max := 0;
  for i := 999; i >= 100; i-- {
    for j := 999; j >= 100; j-- {
      product := i * j

      if isPalindrome(product) {
        if product > max {
          max = product
        }
      }
    }
  }

  fmt.Printf("-->Answer: The maximum palindrome of two 3 digit numbers is %d", max)

}

func isPalindrome(p int) bool {

  s := strconv.Itoa(p)

  midpoint := len(s) / 2

  for i := 0; i < midpoint; i++ {
    if s[i] != s[len(s) - i - 1] {
      return false
    }
  }
  return true
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

