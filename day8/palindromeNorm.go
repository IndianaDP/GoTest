package main

import (
	"fmt"
	"strings"
	"unicode"
)

func PalindromeNormalized(s string) (bool, error) {
	s = strings.ToLower(strings.TrimSpace(s))

	var norm []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			norm = append(norm, r)
		}
	}

	n := len(norm)

	if n == 0 {
		return false, fmt.Errorf("empty string")
	}

	if n == 1 {
		return false, fmt.Errorf("single character string")
	}

	for i := 0; i < n/2; i++ {

		if norm[i] != norm[n-i-1] {
			return false, nil
		}
	}
	return true, nil
}
