package gotest

import "testing"

var Result int

// 测试内联性能优化
func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func BenchmarkMax(t *testing.B) {
	var r int
	for i:=0; i<t.N; i++ {
		if -1 > i {
			r = -1
		} else {
			r = i
		}
	}
	Result = r
}

