package removeelement

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		in1  []int
		in2  int
		want int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for _, tt := range tests {
		got := removeElement(tt.in1, tt.in2)

		if got != tt.want {
			t.Fatalf("in1: %v, in2: %v, got: %v, want: %v", tt.in1, tt.in2, got, tt.want)
		}
	}
}
