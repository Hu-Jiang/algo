package easy

import (
	"testing"
)

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
	}

	for i, tt := range tests {
		got := countAndSay(tt.input)
		if got != tt.want {
			t.Fatalf("%d: got %v; want %v", i, got, tt.want)
		}
	}
}
