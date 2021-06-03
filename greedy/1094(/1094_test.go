package greedy

import "testing"

func Test_carPooling(t *testing.T) {
	tests := []struct {
		name     string
		trips    [][]int
		capacity int
		expect   bool
		real     bool
	}{
		{"case1", [][]int{{2, 1, 5}, {3, 3, 7}}, 4, false, true},
		{"case2", [][]int{{2, 1, 5}, {3, 3, 7}}, 5, true, false},
		{"case3", [][]int{{2, 1, 5}, {3, 5, 7}}, 3, true, false},
		{"case4", [][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}, 11, true, false},
		{"case5", [][]int{{10, 5, 7}, {10, 3, 4}, {7, 1, 8}, {6, 3, 4}}, 24, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.real = carPooling(tt.trips, tt.capacity); tt.expect != tt.real {
				t.Fatal("expect:", tt.expect, "real:", tt.real)
			}
		})
	}
}
