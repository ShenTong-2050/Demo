package goroutine

import (
	"fmt"
	"runtime"
)

func GoschedMain() {

	// 指定分配一个线程执行任务
	runtime.GOMAXPROCS(1)

	exit := make(chan struct{})


	go func() {

		defer close(exit)

		go func() {				// 任务 b 放在当前位置是为了确保 a 优先执行
			fmt.Println("this is b")
		}()

		for i := 0; i<4; i++ {

			fmt.Println("this is a",i)

			if i == 1 {
				runtime.Gosched()	// 让出当前线程，调度执行 b
			}
		}
	}()

	// 阻塞直到关闭
	<-exit
}
