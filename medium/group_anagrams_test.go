package medium

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{nil, nil},
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				[]string{"eat", "tea", "ate"},
				[]string{"tan", "nat"},
				[]string{"bat"},
			},
		},
	}

	for i, tt := range tests {
		got := groupAnagrams(tt.strs)
		assertGroupAngrams(t, i, got, tt.want)
	}
}

func assertGroupAngrams(t *testing.T, caseNum int, got, want [][]string) {
	t.Helper()

	if len(got) != len(want) {
		t.Fatalf("%d: got: %v; want: %v", caseNum, got, want)

	}

	for _, v := range got {
		var equal bool
		for _, vv := range want {
			if reflect.DeepEqual(v, vv) {
				equal = true
				break
			}
		}
		if !equal {
			t.Fatalf("%d: got: %v; want: %v", caseNum, got, want)
		}
	}
}
