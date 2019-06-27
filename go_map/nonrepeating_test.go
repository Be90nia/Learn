package main

import "testing"

func TestLengthOfNonRepeatingSubStr(t *testing.T) {
	tests := []struct{
		s string
		ans int
	}{
		{"abcabcbb",3},
		{"pwwkew",3},
		{"",0},
		{"bbbbbb",0},
		{"bbbdddsa",2},
		{"dadffdsdfad",5},
	}

	for _,tt := range tests {
		if ans := LengthOfNonRepeatingSubStr(tt.s);ans != tt.ans{
			t.Errorf("LengthOfNonRepeatingSubStr(%s);"+
				" got:%d; expected:%d",
				tt.s, ans, tt.ans)
		}
	}
}