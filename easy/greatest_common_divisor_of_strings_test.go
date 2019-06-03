package easy

import "testing"

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		str1, str2, want string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"LEET", "COOD", ""},
	}

	for i, tt := range tests {
		got := gcdOfStrings(tt.str1, tt.str2)
		if got != tt.want {
			t.Fatalf("test %d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
