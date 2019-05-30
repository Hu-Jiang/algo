package easy

import "testing"

func TestAddDigits(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{38, 2},
		{2, 2},
	}

	for i, tt := range tests {
		got := addDigits(tt.input)
		if got != tt.want {
			t.Fatalf("case %d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
