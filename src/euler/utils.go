package euler

import (
    "fmt"
    "time"
    "strconv"
)

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

func Ctoi(b byte) int {
    return int(b)-int('0')

}

func LoadsOfPrimes() {

    value := 10000000

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
