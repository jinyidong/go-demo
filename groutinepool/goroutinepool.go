package groutinepool

import (
	"fmt"
	"sync"
	"time"
)

type worker struct {
	Func func()
}
func gopool2()  {
	var wg sync.WaitGroup
	channels:=make(chan worker,100)

	//起5个协程处理channels中的任务,worker其实更应该成为task
	for i:=0; i<5;i++  {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for channel:=range channels{
				channel.Func()
			}
		}()
	}

	//向channels中压入task
	for i:=0;i<10000 ;i++  {
		j:=i
		wk := worker{
			Func: func() {
				fmt.Println(j + j)
			},
		}
		channels <- wk
	}

	close(channels)
	wg.Wait()
}



func gopool1()  {
	start:=time.Now()
	wg:=sync.WaitGroup{}
	data:=make(chan int,100)

	for i:=0;i<10 ;i++  {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for _=range data{
				fmt.Println("groutine:",n,i)
			}
		}(i)
	}

	for i:=0;i<10000 ;i++  {
		data<-i
	}

	close(data)
	wg.Wait()
	end:=time.Now()
	fmt.Println(end.Sub(start))
}
