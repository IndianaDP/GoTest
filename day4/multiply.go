package main

import (
	"fmt"
	"math"
)

func Multiply(a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, fmt.Errorf("indeterminate form: 0*0 is not allowed")
	}

	if a < 0 || b < 0 {
		return 0, fmt.Errorf("negative numbers are not allowed")
	}

	if a != 0 && b > math.MaxInt32/a {
		return 0, fmt.Errorf("integer overflow")
	}

	return a * b, nil
}
