package easy

import (
	"reflect"
	"testing"
)

func TestDistributeCandies(t *testing.T) {
	tests := []struct {
		candies   int
		numPeople int
		want      []int
	}{
		{0, 0, nil},
		{7, 4, []int{1, 2, 3, 1}},
		{10, 3, []int{5, 2, 3}},
		{60, 4, []int{15, 18, 15, 12}},
	}

	for i, tt := range tests {
		got := distributeCandies(tt.candies, tt.numPeople)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("%d. got %v, want %v", i, got, tt.want)
		}
	}
}
