package main

import (
	"fmt"
)

func Divide(a, b int) (int, error) {
	if a == 0 && b == 0 {
		return 0, fmt.Errorf("indeterminate form: 0/0 is not allowed")
	}

	if b == 0 {
		return 0, fmt.Errorf("devision by zero is not allowed")
	}

	if a < 0 || b < 0 {
		return 0, fmt.Errorf("negative numbers are not allowed")
	}

	return a / b, nil
}
