package hard

import (
	"testing"
)

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		s1, s2, s3 string
		want       bool
	}{
		{
			s1:   "",
			s2:   "",
			s3:   "",
			want: true,
		},
		{
			s1:   "abc",
			s2:   "",
			s3:   "abc",
			want: true,
		},
		{
			s1:   "abc",
			s2:   "abc",
			s3:   "abc",
			want: false,
		},
		{
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbcbcac",
			want: true,
		},
		{
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbbaccc",
			want: false,
		},
	}

	for i, tt := range tests {
		got := isInterleave(tt.s1, tt.s2, tt.s3)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
