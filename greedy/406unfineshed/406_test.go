package greedy

import "testing"

func Test_reconstructQueue(t *testing.T) {
	tests := []struct {
		name   string
		people [][]int
		expect [][]int
		real   [][]int
	}{
		{
			name:   "case1",
			people: [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
			expect: [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}},
			real:   [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 1}},
		},
		{
			name:   "case2",
			people: [][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}},
			expect: [][]int{{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0}},
			real:   [][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Fatal(reconstructQueue(tt.people))
		})
	}
}
