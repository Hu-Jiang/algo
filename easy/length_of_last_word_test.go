package easy

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"", 0},
		{"    ", 0},
		{"  aaa   ", 3},
		{"abcdef", 6},
		{"Hello World", 5},
	}

	for i, tt := range tests {
		got := lengthOfLastWord(tt.s)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
