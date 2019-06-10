package easy

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"", true},
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for i, tt := range tests {
		got := isValid(tt.s)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
