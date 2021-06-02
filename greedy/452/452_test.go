package greedy

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		expect int
		real   int
	}{
		{"case1", [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2, 0},
		{"case2", [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4, 0},
		{"case3", [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.real = findMinArrowShots(tt.points); tt.expect != tt.real {
				t.Fatal("expect: ", tt.expect, "real: ", tt.real)
			}

		})
	}
}
