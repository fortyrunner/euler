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
  fmt.Println("Project Euler - Problem 1. Find the sum of all the multiples of 3 or 5 below 1000")
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

  fmt.Println("Project Euler - Problem 2. By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.")
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

  fmt.Println("Project Euler - Problem 3. What is the largest prime factor of the number 600851475143")
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

  fmt.Println("Project Euler - Problem 4. Find the largest palindrome made from the product of two 3-digit numbers")
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

  fmt.Println("Project Euler - Problem 5. What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?")
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

