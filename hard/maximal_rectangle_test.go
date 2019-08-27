package hard

import "testing"

func TestMaximalRectangle(t *testing.T) {
	tests := []struct {
		matrix [][]byte
		want   int
	}{
		{
			matrix: nil,
			want:   0,
		},
		{
			matrix: [][]byte{
				[]byte{'1', '0', '1', '1'},
			},
			want: 2,
		},
		{
			matrix: [][]byte{
				[]byte{'1', '0', '1', '0', '0'},
				[]byte{'1', '0', '1', '1', '1'},
				[]byte{'1', '1', '1', '1', '1'},
				[]byte{'1', '0', '0', '1', '0'},
			},
			want: 6,
		},
	}

	for i, tt := range tests {
		got := maximalRectangle(tt.matrix)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
