package mathutil

import (
	"strings"
	"testing"
)

// Benchmark Multiply
func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(123, 456) // пример чисел
	}
}

// Benchmark Divide с суб-бенчмарками
func BenchmarkDivide(b *testing.B) {
	cases := []struct {
		name string
		a, b int
	}{
		{"small numbers", 10, 2},
		{"large numbers", 1_000_000, 2},
		{"division by zero", 10, 0},
	}

	for _, tt := range cases {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Divide(tt.a, tt.b)
			}
		})
	}
}

// Benchmark поиска через цикл
func BenchmarkSearchLoop(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		found := false
		for _, v := range data {
			if v == target {
				found = true
				break
			}
		}
		_ = found
	}
}

// Benchmark поиска через strings.Contains
func BenchmarkSearchContains(b *testing.B) {
	data := "1,2,3,4,5,6,7,8,9,10"
	target := "7"

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = strings.Contains(data, target)
	}
}
