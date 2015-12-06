package utils

import "strconv"

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

