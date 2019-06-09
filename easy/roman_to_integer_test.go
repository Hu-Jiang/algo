package easy

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MMMXLV", 3045},
	}

	for i, tt := range tests {
		got := romanToInt(tt.s)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
