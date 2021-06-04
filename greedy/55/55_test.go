package greedy

import "testing"

func Test_canJump(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		expect bool
		real   bool
	}{
		{"case1", []int{2, 3, 1, 1, 4}, true, false},
		{"case2", []int{3, 2, 1, 0, 4}, false, true},
		{"case2", []int{2, 5, 0, 0, }, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.real = canJump(tt.nums); tt.expect != tt.real {
				t.Fatal("expect:", tt.expect, "real:", tt.real)
			}
		})
	}
}
