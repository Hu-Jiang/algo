package hard

import (
	"reflect"
	"testing"
)

func TestSmallestSufficientTeam(t *testing.T) {
	tests := []struct {
		reqSkills []string
		people    [][]string
		want      []int
	}{
		{nil, nil, []int{}},
		{
			reqSkills: []string{"java", "nodejs", "reactjs"},
			people: [][]string{
				[]string{"java"},
				[]string{"nodejs"},
				[]string{"nodejs", "reactjs"},
			},
			want: []int{0, 2},
		},
		{
			reqSkills: []string{"gvp", "jgpzzicdvgxlfix", "kqcrfwerywbwi", "jzukdzrfgvdbrunw", "k"},
			people: [][]string{
				[]string{},
				[]string{},
				[]string{},
				[]string{},
				[]string{"jgpzzicdvgxlfix"},
				[]string{"jgpzzicdvgxlfix", "k"},
				[]string{"jgpzzicdvgxlfix", "kqcrfwerywbwi"},
				[]string{"gvp"},
				[]string{"jzukdzrfgvdbrunw"},
				[]string{"gvp", "kqcrfwerywbwi"},
			},
			want: []int{5, 8, 9},
		},
	}

	for i, tt := range tests {
		got := smallestSufficientTeam(tt.reqSkills, tt.people)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("#%d. got %v, want %v", i, got, tt.want)
		}
	}
}
