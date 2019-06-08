package easy

import (
	"testing"
)

func TestReverseInteger(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{123, 321},
		{-123, -321},
		{1234567899, 0},
		{-1234567899, 0},
		{0, 0},
	}

	for i, tt := range tests {
		got := reverse(tt.input)
		if got != tt.want {
			t.Fatalf("%d: got: %d; want: %d", i, got, tt.want)
		}
	}
}
