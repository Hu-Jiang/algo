package hard

import "testing"

func TestSubStrHash(t *testing.T) {
	tests := []struct {
		s         string
		power     int
		modulo    int
		k         int
		hashValue int

		want string
	}{
		{
			s:         "leetcode",
			power:     7,
			modulo:    20,
			k:         2,
			hashValue: 0,
			want:      "ee",
		},
		{
			s:         "fbxzaad",
			power:     31,
			modulo:    100,
			k:         3,
			hashValue: 32,
			want:      "fbx",
		},
		{
			s:         "a",
			power:     31,
			modulo:    100,
			k:         1,
			hashValue: 1,
			want:      "a",
		},
		{
			s:         "xmmhdakfursinye",
			power:     96,
			modulo:    45,
			k:         15,
			hashValue: 21,
			want:      "xmmhdakfursinye",
		},
		{
			s:         "xxterzixjqrghqyeketqeynekvqhc",
			power:     15,
			modulo:    94,
			k:         4,
			hashValue: 16,
			want:      "nekv",
		},
		{
			s:         "abc",
			power:     7,
			modulo:    10,
			k:         1,
			hashValue: 1,
			want:      "a",
		},
	}

	for i, tt := range tests {
		got := subStrHash(tt.s, tt.power, tt.modulo, tt.k, tt.hashValue)
		if got != tt.want {
			t.Fatalf("case %d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
