package hard

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{nil, []int{2, 3}, 2.5},
		{nil, []int{1}, 1.0},
	}

	for i, tt := range tests {
		got := findMedianSortedArrays(tt.nums1, tt.nums2)
		if got != tt.want {
			t.Fatalf("case %d: got: %v; want %v", i, got, tt.want)
		}
	}
}
