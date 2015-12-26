package euler

import "testing"

func TestIsPalindrome(t *testing.T) {

    if !IsPalindrome(0){
        t.Error("0 should be a palindrome of itself")
    }

    if !IsPalindrome(1){
        t.Error("1 should be a palindrome of itself")
    }

    if !IsPalindrome(11){
        t.Error("11 should be a palindrome of itself")
    }

    if !IsPalindrome(111){
        t.Error("111 should be a palindrome of itself")
    }

    if !IsPalindrome(121){
        t.Error("121 should be a palindrome of itself")
    }

    if IsPalindrome(43){
        t.Error("43 should be a palindrome of itself")
    }




}