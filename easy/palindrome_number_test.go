package easy

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{0, true},
	}

	for i, tt := range tests {
		got := isPalindrome(tt.input)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
