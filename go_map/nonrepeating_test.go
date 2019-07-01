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
		//{"bbbbbb",0},
		//{"bbbdddsa",2},
		//{"dadffdsdfad",5},
	}

	for _,tt := range tests {
		if ans := LengthOfNonRepeatingSubStr(tt.s);ans != tt.ans{
			t.Errorf("LengthOfNonRepeatingSubStr(%s);"+
				" got:%d; expected:%d",
				tt.s, ans, tt.ans)
		}
	}
}
//性能测试
func BenchmarkLengthOfNonRepeatingSubStr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s+s
	}
	b.Logf("len(s) = %d",len(s))
	ans := 8
	b.ResetTimer()//重置计算时间
	for i := 0; i < b.N; i++ {
		if actual := LengthOfNonRepeatingSubStr(s); actual != ans {
			b.Errorf("LengthOfNonRepeatingSubStr(%s);"+
				" got:%d; expected:%d",
				s, actual, ans)
		}
	}
}