package euler

import "testing"

func TestIsPalindrome(t *testing.T) {

        if !IsPalindrome(0) {
                t.Error("0 should be a palindrome of itself")
        }

        if !IsPalindrome(1) {
                t.Error("1 should be a palindrome of itself")
        }

        if !IsPalindrome(11) {
                t.Error("11 should be a palindrome of itself")
        }

        if !IsPalindrome(111) {
                t.Error("111 should be a palindrome of itself")
        }

        if !IsPalindrome(121) {
                t.Error("121 should be a palindrome of itself")
        }

        if IsPalindrome(43) {
                t.Error("43 should be a palindrome of itself")
        }

}

func TestIsPrime(t *testing.T) {

        if IsPrime(1) {
                t.Error("1 is not a prime")
        }

        if !IsPrime(2) {
                t.Error("2 IS a prime")
        }

        if !IsPrime(3) {
                t.Error("3 IS a prime")
        }

        if !IsPrime(101) {
                t.Error("101 IS a prime")
        }

        if IsPrime(1717172727272) {
                t.Error("1717172727272 is not a prime")
        }

}
