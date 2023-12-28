package goselect

import (
	"fmt"
	"time"
)

func TimeOutChan(ch chan string) {
	select {
	/*case <-time.Tick(time.Second):
		fmt.Println(time.Now())*/
	case msg,ok := <-ch:
		fmt.Printf("read msg from chan: %s, %v\n",msg, ok)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout...")
	}
}

func TimeOutChanMain() {

	go func() {
		for {
			TimeOutChan(make(chan string,2))
		}
	}()

	// 阻塞主协程
	<-(chan struct{})(nil)
}
