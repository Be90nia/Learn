package main
//BenchmarkLengthOfNonRepeatingSubStr-16    	  200000	     10715 ns/op
var lastOccurred = make([]int,0xffff)

func LengthOfNonRepeatingSubStr(s string) int {
	//BenchmarkLengthOfNonRepeatingSubStr-16    	 2000000	       780 ns/op
	//lastOccurred := make(map[rune]int)
	//start := 0
	//maxLength := 0
	//for i, ch := range []rune(s) {
	//	if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
	//		start = lastI + 1
	//	}
	//	if i-start+1 > maxLength {
	//		maxLength = i - start + 1
	//	}
	//	lastOccurred[ch] = i
	//}
	//return maxLength

	//BenchmarkLengthOfNonRepeatingSubStr-16    	   50000	     25133 ns/op
	//stores last occurred pos + 1
	//0 means not seen.
	//此段代码明显性能消耗过大,查看过后发现主要消耗在垃圾回收这块的问题
	//lastOccurred := make([]int,0xffff)  //将其设置为全局的查看性能消耗

	for i := range lastOccurred{
		lastOccurred[i] = 0
	}
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI >= start {
			start = lastI
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i + 1
	}
	return maxLength
}
