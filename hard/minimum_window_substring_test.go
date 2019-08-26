package hard

import (
	"testing"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want string
	}{
		{
			s:    "",
			t:    "",
			want: "",
		},
		{
			s:    "",
			t:    "ABC",
			want: "",
		},
		{
			s:    "ABC",
			t:    "",
			want: "",
		},
		{
			s:    "AAAA",
			t:    "BB",
			want: "",
		},
		{
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
	}

	for i, tt := range tests {
		got := minWindow(tt.s, tt.t)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
