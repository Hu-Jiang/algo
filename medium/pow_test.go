package medium

import (
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	tests := []struct {
		x    float64
		n    int
		want float64
	}{
		{
			x:    2,
			n:    10,
			want: 1024,
		},
		{
			x:    2.1,
			n:    3,
			want: math.Pow(2.1, 3),
		},
		{
			x:    2,
			n:    -2,
			want: math.Pow(2, -2),
		},
	}

	for i, test := range tests {
		if myPow(test.x, test.n) != test.want {
			t.Fatalf("test myPow %d, want: %v, got: %v", i+1, test.want, myPow(test.x, test.n))
		}
		if myPow_v1(test.x, test.n) != test.want {
			t.Fatalf("test iterPow %d, want: %v, got: %v", i+1, test.want, myPow_v1(test.x, test.n))
		}
	}
}

func BenchmarkMyPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myPow(10, i)
	}
}

func BenchmarkSysPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Pow(10, float64(i))
	}
}

func BenchmarkMyIterPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myPow_v1(10, i)
	}
}
