package goselect

import (
	"fmt"
	"time"
)

// select 多通道选择 应用

func multiChannel(ch1 <-chan string,ch2 chan string,ch3 chan string) {
	msg2 := "msg2"
	select {
	case msg1 := <-ch1:
		fmt.Printf("received msg from ch1 %v \n",msg1)
	case ch2 <-msg2:
	case msg3,ok := <-ch3 :
		fmt.Printf("received msg from ch3 %v,ok:%v\n",msg3,ok)
	default:
		fmt.Printf("Unknown message\n")
	}
}

func MultiChannelMain() {

	ch1,ch2,ch3 := make(chan string),make(chan string),make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			ch1<-"message one"
			ch3<-"message three"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			multiChannel(ch1,ch2,ch3)
		}
	}()

	// 读取 chan2 的 消息
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Printf("read msg from chan2 %v \n",<-ch2)
		}
	}()

	// 阻塞
	<-(chan struct{})(nil)
}
