package medium

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{
			s:    "",
			want: nil,
		},
		{
			s:    "111",
			want: nil,
		},
		{
			s:    "0000",
			want: []string{"0.0.0.0"},
		},
		{
			s:    "00000",
			want: nil,
		},
		{
			s:    "11011",
			want: []string{"1.1.0.11", "1.10.1.1", "11.0.1.1"},
		},
		{
			s:    "25525511135",
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
	}

	for i, tt := range tests {
		got := restoreIpAddresses(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
