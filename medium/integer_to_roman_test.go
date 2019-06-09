package medium

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for i, tt := range tests {
		got := intToRoman(tt.num)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want %v", i, got, tt.want)
		}
	}
}
