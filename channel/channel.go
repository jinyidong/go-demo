package channel

import (
	"fmt"
	"time"
)

func ChannelSelect()  {
	exitChan:=make(chan bool)
	testChan:=make(chan struct{})

	go func() {
		for{
			select {
			case <-testChan:
				fmt.Println("read from chan")
			default:
				fmt.Println("default")
			}
		}
	}()

	time.Sleep(10*time.Second)

	close(testChan)

	<-exitChan
}


//chan两次关闭测试
func DoubleCloseChan()  {
	testChan:=make(chan bool)
	close(testChan)
	close(testChan)
}

//向已关闭的channel中写入
func WriteCloseChan()  {
	testChan:=make(chan bool)
	close(testChan)
	testChan<-false
}

func ReadCloseChanWithNoBuffer()  {
	testChan:=make(chan bool)
	close(testChan)
	temp:=<-testChan
	fmt.Println(temp)
}

func TwoConsumer()  {
	testChan:=make(chan bool)

	go func() {
		fmt.Println("1",<-testChan)
	}()

	go func() {
		fmt.Println("2",<-testChan)
	}()

	testChan<-true

	time.Sleep(10*time.Second)
}



