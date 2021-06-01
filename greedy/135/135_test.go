package greedy

import "testing"

func Test_candy(t *testing.T) {
	tests := []struct {
		name    string
		ratings []int
		expect  int
		real    int
	}{
		{"case1", []int{1, 0, 2}, 5, 0},
		{"case2", []int{1, 2, 2}, 4, 0},
		{"case3", []int{1, 3, 2, 2, 1}, 7, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.real = candy(tt.ratings); tt.real != tt.expect {
				t.Fatal("real: ", tt.real, "expect: ", tt.expect)
			}
		})
	}
}
