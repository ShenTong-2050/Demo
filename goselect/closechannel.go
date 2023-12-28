package goselect

import (
	"fmt"
	"os"
	"time"
)

// 监听通道关闭操作，执行对应逻辑

func closeChan(ch chan string) {
	select {
	case <-ch:
		fmt.Println("通道已关闭")
		os.Exit(201) // 通道关闭后程序中断执行
	default:
		fmt.Println("通道未关闭")
	}
}

func ClosingMain() {

	ch := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second) // 每个1秒
			closeChan(ch) // 检测通道是否关闭
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 5) // 5秒后执行关闭通道操作
			close(ch)
		}
	}()

	// 阻塞主协程
	<-(chan struct{})(nil)
}
