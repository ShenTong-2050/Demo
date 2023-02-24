package goroutine

import "runtime"

func GoExitMain() {

	ch := make(chan struct{})

	go func() {

		defer close(ch)
		defer println("a")
		func(){
			defer func() {
				println("b",recover() == nil)		// 执行,recover 返回 nil
			}()

			func() {
				println("c")
				runtime.Goexit()
				println("c done")
			}()

			println("b done")
		}()
		println("a done")
	}()

	<-ch

	println("main done")
}
