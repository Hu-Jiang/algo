package easy

import (
	"testing"
)

func TestDayOfYear(t *testing.T) {
	tests := []struct {
		data string
		want int
	}{
		{"abc", 0},
		{"abc-d-e", 0},
		{"2019", 0},
		{"2019-01-09", 9},
		{"2019-02-10", 41},
		{"2003-03-01", 60},
		{"2004-03-01", 61},
	}

	for i, tt := range tests {
		got := dayOfYear(tt.data)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
