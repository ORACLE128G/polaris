package main

import (
	"math"
	"testing"
)

func triangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}
func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
	}
	for _, tt := range tests {
		if actual := triangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); got %d, expected %d",
				tt.a, tt.b, actual, tt.c)
		}
	}
}


func BenchmarkTriangle(b *testing.B) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
	}
	for i := 0; i < b.N; i++ {
		if actual := triangle(tests[0].a, tests[0].b); actual != tests[0].c {
			b.Errorf("calcTriangle(%d, %d); got %d, expected %d",
				tests[i].a, tests[i].b, actual, tests[i].c)
		}
	}
}