package easy

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{0, true},
		{16, true},
		{14, false},
	}

	for i, tt := range tests {
		got := isPerfectSquare(tt.num)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
