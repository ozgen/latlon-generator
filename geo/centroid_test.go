package geo

import (
	"math"
	"testing"
)

func floatsEqual(a, b float64) bool {
	const epsilon = 1e-6
	return math.Abs(a-b) < epsilon
}

func TestCalculatePolygonCentroid(t *testing.T) {
	tests := []struct {
		name   string
		input  LinearRing
		expect Point
	}{
		{
			name: "square centered at origin",
			input: LinearRing{
				{-1, -1}, {1, -1}, {1, 1}, {-1, 1}, {-1, -1},
			},
			expect: Point{0, 0},
		},
		{
			name: "rectangle at offset",
			input: LinearRing{
				{2, 2}, {4, 2}, {4, 4}, {2, 4}, {2, 2},
			},
			expect: Point{3, 3},
		},
		{
			name: "triangle",
			input: LinearRing{
				{0, 0}, {4, 0}, {2, 4}, {0, 0},
			},
			expect: Point{2, 1.333333},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CalculatePolygonCentroid(tc.input)
			if !floatsEqual(got[0], tc.expect[0]) || !floatsEqual(got[1], tc.expect[1]) {
				t.Errorf("expected %v, got %v", tc.expect, got)
			}
		})
	}
}
