package linkedlist

import "testing"

func Test_buildList(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := buildList()
			t.Fatal(node)
		})
	}
}
func Test_addNode(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := buildList()
			list.addNode(2)
			t.Fatal(list.Next.Val)
		})
	}
}
func Test_getKthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		k    int
	}{
		{"case1", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := buildList()
			list.addNode(2)
			list.addNode(3)
			list.addNode(4)
			list.addNode(5)
			result := getKthFromEnd(list, tt.k)
			t.Fatal(result)
		})
	}
}
