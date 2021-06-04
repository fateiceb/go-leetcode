package greedy

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	tests := []struct {
		name   string
		gas    []int
		cost   []int
		expect int
		real   int
	}{
		{"case1", []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3, 0},
		{"case2", []int{2, 3, 4}, []int{3, 4, 3}, -1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.real = canCompleteCircuit(tt.gas, tt.cost); tt.real != tt.expect {
				t.Fatal("expect:", tt.expect, "real:", tt.real)
			}
		},
		)
	}
}
