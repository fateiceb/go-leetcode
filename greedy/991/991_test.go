package gredy

import (
	"log"
	"testing"
)

func Test_brokenCalc(t *testing.T) {
	tests := []struct {
		name   string
		x      int
		y      int
		expect int
		real   int
	}{
		{"case1", 2, 3, 2, 0},
		{"case2", 5, 8, 2, 0},
		{"case3", 3, 10, 3, 0},
		{"case4", 1024, 1, 1023, 0},
		{"case5", 1, 100000000, 35, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.real = brokenCalc(tt.x, tt.y); tt.real != tt.expect {
				t.Fatal(tt.name, "real:", tt.real, " expect:", tt.expect)
			}
		})
	}
}
