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
		{6, 7, 8},
		{3000000, 400000, 700000},
	}
	for _, tt := range tests {
		if actual := triangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); got %d, expected %d",
				tt.a, tt.b, actual, tt.c)
		}
	}
}
