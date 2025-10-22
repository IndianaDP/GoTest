package main

import (
	"fmt"
)

func Palindrome(s string) (bool, error) {
	n := len(s)

	if n == 0 {
		return false, fmt.Errorf("empty string")
	}

	if n == 1 {
		return false, fmt.Errorf("single character string")
	}

	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false, nil
		}
	}
	return true, nil
}
