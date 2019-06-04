package medium

import (
	"reflect"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		paths []string
		want  [][]string
	}{
		{
			nil,
			nil,
		},
		{
			[]string{"root/a 1.txt(abcd)"},
			nil,
		},
		{
			[]string{"root/a 1.txt(abcd)", "root/b 2.txt(abcd) 3.txt(efgh)"},
			[][]string{[]string{"root/a/1.txt", "root/b/2.txt"}},
		},
		{
			[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"},
			[][]string{[]string{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"}, []string{"root/a/1.txt", "root/c/3.txt"}},
		},
	}

	for i, tt := range tests {
		got := findDuplicate(tt.paths)
		if !disorderEqual(got, tt.want) {
			t.Fatalf("test %d: got: %v; want: %v", i, got, tt.want)
		}
	}
}

func disorderEqual(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for _, v := range a {
		if !contain(b, v) {
			return false
		}
	}

	for _, v := range b {
		if !contain(a, v) {
			return false
		}
	}

	return true
}

func contain(arr [][]string, str []string) bool {
	for _, v := range arr {
		if reflect.DeepEqual(v, str) {
			return true
		}
	}

	return false
}
