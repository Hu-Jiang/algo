package easy

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a, b, want string
	}{
		{"0", "0", "0"},
		{"1", "0", "1"},
		{"1", "1", "10"},
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"1", "111", "1000"},
		{"100", "110010", "110110"},
	}

	for i, tt := range tests {
		got := addBinary(tt.a, tt.b)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
