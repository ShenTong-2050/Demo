package goroutine

import (
	"math"
	"runtime"
	"sync"
)

func count()  {
	var x = 0
	for i:=0; i<math.MaxInt32; i++ {
		x += i
	}
	println(x)
}

func SingleCpuCount(n int) {
	for i:=0; i<n; i++ {
		count()
	}
}

func GoroutineCpuCount(n int) {
	var Wg sync.WaitGroup
	Wg.Add(n)
	for i:=0; i<n; i++ {
		go func() {
			count()
		}()
		Wg.Done()
	}
	Wg.Wait()
}

func GoroutineMain() {
	var n = runtime.GOMAXPROCS(0)
	//SingleCpuCount(n)
	GoroutineCpuCount(n)
	//time.Sleep(time.Second)
}
