package goselect

import (
	"fmt"
	"time"
)

// select 多通道选择 应用

func multiChannel(ch1 <-chan string,ch2 chan string,ch3 chan string) {
	msg2 := "hello world"
	select {
	case msg1 := <-ch1:
		fmt.Printf("read msg from chan1 %d\n",msg1)
	case ch2<-msg2:
	case msg3,ok := <-ch3:
		fmt.Printf("read msg and chan is closed %d,%d\n",msg3,ok)
	}
}

func MultiChannelMain() {

	ch1,ch2,ch3 := make(chan string),make(chan string),make(chan string)

	// 往 ch1,ch3 生产消息
	go func() {
		for {
			time.Sleep(time.Second)
			ch1 <-"message one"
			ch3 <-"message three"
		}
	}()

	// 起个协程执行 消费 消息 操作
	go func(){
		for {
			time.Sleep(time.Second)
			multiChannel(ch1,ch2,ch3)
		}
	}()

	// 接收 ch2 的 消息
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Printf("read msg from chan2 %d \n",<-ch2)
		}
	}()

	// 阻塞主协程
	<-(chan struct{})(nil)
}
