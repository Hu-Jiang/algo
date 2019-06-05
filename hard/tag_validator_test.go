package hard

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		code string
		want bool
	}{
		{"<DIV>This is the first line <![CDATA[<div>]]></DIV>", true},
		{"<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>", true},
		{"<A>  <B> </A>   </B>", false},
		{"<DIV>  div tag is not closed  <DIV>", false},
		{"<DIV>  unmatched <  </DIV>", false},
		{"<DIV> closed tags with invalid tag name  <b>123</b> </DIV>", false},
		{"<DIV> unmatched tags with invalid tag name  </1234567890> and <CDATA[[]]>  </DIV>", false},
		{"<DIV>  unmatched start tag <B>  and unmatched end tag </C>  </DIV>", false},
		{"<![CDATA[wahaha]]]><![CDATA[]> wahaha]]>", false},
		{"<A></A><B></B>", false},
		{"<A></A>A>", false},
		{"<A></A><A></A>", false},
		{"<A></A>>", false},
		{"", false},
	}

	for i, tt := range tests {
		got := isValid(tt.code)
		if got != tt.want {
			t.Fatalf("case %d: got: %v; want: %v", i, got, tt.want)
		}
	}
}
