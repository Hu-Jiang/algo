package hard

import "testing"

func TestMinDeletionSize(t *testing.T) {
	tests := []struct {
		A    []string
		want int
	}{
		{[]string{"babca", "bbazb"}, 3},
		{[]string{"edcba"}, 4},
		{[]string{"ghi", "def", "abc"}, 0},
		{[]string{"aabbaa", "baabab", "aaabab"}, 3},
	}

	for i, tt := range tests {
		got := minDeletionSize(tt.A)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
