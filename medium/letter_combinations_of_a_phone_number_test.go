package medium

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		digits string
		want   []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", nil},
	}

	for i, tt := range tests {
		got := letterCombinations(tt.digits)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
