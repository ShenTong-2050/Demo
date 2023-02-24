package gotest

import "testing"

// 对比数组与切片拷贝内存占用的情况
func testArray() [1024]int {
	var x [1024]int
	for i:=0; i<len(x); i++ {
		x[i] = i
	}
	return x
}

func testSlice() []int {
	var s []int = make([]int,1024)
	for i:=0; i<len(s); i++ {
		s[i] = i
	}
	return s
}

func BenchmarkArray(t *testing.B) {
	for i:=0; i<t.N; i++ {
		testArray()
	}
}

func BenchmarkSlice(t *testing.B) {
	for i:=0; i<t.N; i++ {
		testSlice()
	}
}