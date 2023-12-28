package goselect

import (
	"fmt"
	"time"
)

// 非阻塞通道 通过 default 分支实现

func noBlockChan(ch chan string) {
	select {
	case msg := <-ch:
		fmt.Println("read msg from chan",msg)
	default:
		fmt.Println("chan is nil")
	}
}

func NonBlockChan() {
	go func() {
		for {
			time.Sleep(time.Second)
			noBlockChan(make(chan string))
		}
	}()


	<-(chan struct{})(nil)
}