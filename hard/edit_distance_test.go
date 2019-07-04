package hard

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct {
		word1, word2 string
		want         int
	}{
		{"", "", 0},
		{"a", "ab", 1},
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"dinitrophenylhydrazine", "acetylphenylhydrazine", 6},
	}

	for i, tt := range tests {
		got := minDistance(tt.word1, tt.word2)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}

func BenchmarkMinDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestMinDistance(&testing.T{})
	}
}
