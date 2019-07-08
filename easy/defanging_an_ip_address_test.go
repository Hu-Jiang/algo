package easy

import "testing"

func TestDefangIPaddr(t *testing.T) {
	tests := []struct {
		address string
		want    string
	}{
		{"1.1.1.1", "1[.]1[.]1[.]1"},
		{"255.100.50.0", "255[.]100[.]50[.]0"},
	}

	for i, tt := range tests {
		got := defangIPaddr(tt.address)
		if got != tt.want {
			t.Fatalf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
