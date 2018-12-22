package goroutinesync

import (
	"fmt"
	"sync"
)
//go中协程同步有哪些方式
//1、sync.waitgroup
//2、channel

func ChannelSync()  {
	ch:=make(chan struct{})

	count:=2
	go func() {
		fmt.Println("gotoutine 1")
		ch<- struct{}{}
	}()

	go func() {
		fmt.Println("goroutine 2")
		ch<- struct{}{}
	}()

	for range ch{
		count--
		if count==0 {
			close(ch)
		}
	}

}

func WaitGroupSync()  {
	var wg sync.WaitGroup
	wg.Add(2)
	go wgPrint("a",&wg)
	go wgPrint("b",&wg)
	wg.Wait()
	fmt.Println("print:end")
}

func wgPrint(i string,wg *sync.WaitGroup)  {
	fmt.Println("print:"+i)
	wg.Done()
}


