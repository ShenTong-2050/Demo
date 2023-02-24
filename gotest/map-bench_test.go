package gotest

import "testing"

func lessMem() {
	m := make(map[int]int)
	for i:=0; i<1000; i++ {
		m[i] = i
	}
}

func fullMem() {
	m := make(map[int]int , 1024)
	for i:=0; i<1000; i++ {
		m[i] = i
	}
}

func BenchmarkLessMem(t *testing.B) {
	for i:=0; i<t.N; i++ {
		lessMem()
	}
}

func BenchmarkFullMem(t *testing.B) {
	for i:=0; i<t.N; i++ {
		fullMem()
	}
}