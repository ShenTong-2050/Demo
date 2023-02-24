package gotest

import (
	"bytes"
	"strings"
	"testing"
)

func combineStr() string {
	var s string
	for i:=0; i<1000; i++ {
		s = s+"a"
	}
	return s
}

func optimizeCombineStr() string {
	// 分配足够的内存,避免中途扩张底层数据
	s := make([]string,1000)

	for i := 0; i<1000; i++ {
		s[i] = "a"
	}

	return strings.Join(s,"")
}

// 切片拼接字符串算法
func Join(s []string, sep string) string {

	// 统计分隔符长度
	n := len(sep) * (len(s)-1)

	// 统计所有待拼接字符长度
	for i := range s {
		n = n+len(s[i])
	}
	
	// 一次分配所需长度的数组空间
	distribute := make([]byte,n)

	// 拷贝数据
	bp := copy(distribute,s[0])

	for _,v := range s {
		bp += copy(distribute[bp:],sep)
		bp += copy(distribute[bp:],v)
	}

	return string(distribute)
}

// 通过缓冲区优化拼接字符串【核心是提前分配内存,避免在运行过程中扩容底层数组】
func bytesBuffer() string {
	var b bytes.Buffer

	b.Grow(1000)

	for i:=0; i<1000; i++ {
		b.WriteString("a")	// 将字符串 "a" 写入缓冲区
	}

	return b.String()
}

func BenchmarkCombineStr(t *testing.B) {
	for i:=0; i<t.N; i++ {
		combineStr()
	}
}

func BenchmarkOptimizeCombineStr(t *testing.B) {
	for i:=0; i<t.N; i++ {
		optimizeCombineStr()
	}
}

func BenchmarkBytesBuffer(t *testing.B) {
	for i:=0; i<t.N; i++ {
		bytesBuffer()
	}
}

