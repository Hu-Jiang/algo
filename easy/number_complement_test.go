package easy

import "testing"

func TestFindComplement(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{0, 1},
		{1, 0},
		{5, 2},
	}

	for i, tt := range tests {
		got := findComplement(tt.num)
		if got != tt.want {
			t.Fatalf("#%d. got %d, want %d", i, got, tt.want)
		}
	}
}
