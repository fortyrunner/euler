package main

import (
"fmt"
)



func main(){
  problem1()
  problem2()
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

  fmt.Printf("-->Answer - Sum of 3/5 multiples less than %d is %d\n\n", upper_bound, sum )

}
func problem2() {
  fmt.Println("Project Euler - Problem 2. By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.")

  current := 1
  next := 2
  sum := 2

  for next < 4000000{
    //fmt.Printf("%d ", current)

    i := next
    next += current
    current = i

    if next %2 == 0{
      sum += next
    }

  }

  fmt.Printf("-->Answer sum of even values is %d\n\n", sum)

}