package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCases struct {
	name     string
	input    string
	expected bool
	wantErr  bool
}

func TestPalindrome(t *testing.T) {
	tests := []testCases{
		{"empty string", "", false, true},
		{"single character string", "a", false, true},
		{"palindrome with even lenght", "abba", true, false},
		{"palindrome with odd lenght", "aba", true, false},
		{"not a palindrome", "abc", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Palindrome(tt.input)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestPalindromeNormalized(t *testing.T) {
	tests := []testCases{
		{"empty string", "", false, true},
		{"single character string", "a", false, true},
		{"palindrome with even lenght", "abba", true, false},
		{"palindrome with odd lenght", "aba", true, false},
		{"not a palindrome", "abc", false, false},
		{"palindrome 1", "A man a plan, a canal: Panama!", true, false},
		{"palindrome 2", "No 'x' in Nixon", true, false},
		{"palindrome 3", "Never odd, or even.", true, false},
		{"palindrome 4", "Was it a car or a cat I saw?", true, false},
		{"palindrome 5", "Hello, world!", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := PalindromeNormalized(tt.input)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
