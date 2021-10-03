package thread

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	data := make(chan int)
	canQuit := make(chan bool)
	go func() {
		for d := range data {
			fmt.Println(d)
		}
		canQuit <- true
	}()
	data <- 1
	data <- 2
	data <- 3
	close(data)
	<-canQuit

	data = make(chan int, 3)
	go func() {
		for {
			if d, ok := <-data; ok {
				fmt.Println(d)
			} else {
				break
			}
		}
		canQuit <- true
	}()
	data <- 1
	data <- 2
	data <- 3
	data <- 4
	close(data)
	<-canQuit
}

// 超时控制
func TestChannelTimeOut(t *testing.T) {
	done := do()
	select {
	case <-done:
	case <-time.After(3 * time.Second):
		fmt.Println("3 seconds later,done")
	}
}

func do() <-chan struct{} {
	done := make(chan struct{}, 1)
	go func() {
		time.Sleep(10 * time.Second)
		done <- struct{}{}
	}()
	return done
}

// 获取最快的计算结果
func TestGettingFast(t *testing.T) {
	ret := make(chan string, 3)
	go call1(ret)
	go call2(ret)
	go call3(ret)

	fmt.Println(<-ret)
}

func call1(ret chan string) {
	time.Sleep(3 * time.Second)
	ret <- "call1"
}

func call2(ret chan string) {
	time.Sleep(2 * time.Second)
	ret <- "call2"
}

func call3(ret chan string) {
	time.Sleep(4 * time.Second)
	ret <- "call3"
}

// close 广播
func TestCloseBoardCast(t *testing.T) {
	c := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func() {
			<-c
			fmt.Println("copy that")
		}()
	}
	time.Sleep(3 * time.Second)
	close(c)
}
