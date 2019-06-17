package medium

import "testing"

func TestMultiply(t *testing.T) {
	tests := []struct {
		num1 string
		num2 string
		want string
	}{
		{"", "", ""},
		{"", "123", ""},
		{"123", "0", "0"},
		{"2", "3", "6"},
		{"12", "8", "96"},
		{"123", "456", "56088"},
	}

	for i, tt := range tests {
		got := multiply(tt.num1, tt.num2)
		if got != tt.want {
			t.Fatalf("%d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
