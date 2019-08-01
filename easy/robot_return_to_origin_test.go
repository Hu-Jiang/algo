package easy

import "testing"

func TestJudgeCircle(t *testing.T) {
	tests := []struct {
		moves string
		want  bool
	}{
		{"", true},
		{"UD", true},
		{"LL", false},
	}

	for i, tt := range tests {
		got := judgeCircle(tt.moves)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
