package mathutil

import (
	"fmt"
)

func Subtract(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("negative numbers are not allowed")
	}

	if a == 0 && b == 0 {
		return 0, fmt.Errorf("intermediate form: 0-0 is not allowed")
	}

	return a - b, nil
}
