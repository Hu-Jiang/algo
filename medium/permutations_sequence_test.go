package medium

import "testing"

func TestGetPermutation(t *testing.T) {
	tests := []struct {
		n, k int
		want string
	}{
		{3, 3, "213"},
		{4, 9, "2314"},
	}

	for i, tt := range tests {
		got := getPermutation(tt.n, tt.k)
		if got != tt.want {
			t.Fatalf("%d. got: %v; want: %v", i, got, tt.want)
		}
	}
}
