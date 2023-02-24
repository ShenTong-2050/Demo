package bench

import (
	"testing"
	"sync"
	"fmt"
)

var m sync.Mutex

func call() {
	m.Lock()
	m.Unlock()
}

func deferCall() {
	m.Lock()
	defer m.Unlock()
}

func BenchmarkCall(t *testing.B) {
	for i:=0; i<t.N; i++ {
		call()
	}
}

func BenchmarkDeferCall(t *testing.B) {
	for i:=0; i<t.N; i++ {
		deferCall()
	}
}

// 运行分配内存
func BenchmarkAlloc(t *testing.B) {
	for i:=0; i<t.N; i++ {
		fmt.Sprintf("%d",i)
	}
}
