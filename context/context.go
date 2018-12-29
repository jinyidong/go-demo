package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//其本质上还是通过channel来进行协程同步，只不过多了上下文来控制context树上相关协程
func SyncWithContext2()  {
		ctx, cancel := context.WithCancel(context.Background())
		go watch(ctx,"【监控1】")
		go watch(ctx,"【监控2】")
		go watch(ctx,"【监控3】")

		time.Sleep(5 * time.Second)
		fmt.Println("可以了，通知监控停止")
		cancel()
		//为了检测监控过是否停止，如果没有监控输出，就表示停止了
		time.Sleep(20 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		fmt.Println(name,"------------------------>watch:",&ctx)
		select {
		case <-ctx.Done()://channel读事件
			fmt.Println(name,"监控退出，停止了...")
			return
		default:
			fmt.Println(name,"goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
//针对goroutine某一关系链(由某一goroutine衍生其他goroutine)的同步操作如何实现？
func SyncWithContext()  {
	ctx,cancel:=context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for{
			select {
			case <-ctx.Done():
				fmt.Println("stop....")
				return
			default:
				fmt.Println("continue...")
				time.Sleep(2*time.Millisecond)
			}
		}

	}(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

//可以通过共享变量与锁的机制实现对一个go程的控制，但这种使用方式在go中并不推荐
//推荐使用channel+select的方式，但这种方式只能实现对单个go程的控制
func SyncWithChannelSelect()  {
	stop:=make(chan bool)

	go func() {
		for{
			select {
			case <-stop:
				fmt.Println("stop....")
				return
			default:
				fmt.Println("continue...")
				time.Sleep(2*time.Millisecond)
			}
		}

	}()

	time.Sleep(20*time.Millisecond)
	fmt.Println("触发停止条件...")
	stop<-true

	time.Sleep(10*time.Millisecond)
}

//适用于多个goroutine协作做一件事，每个goroutine都是完成其中一部分，只有全部完成后才算完成
func SyncWithWaitGroup()  {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		time.Sleep(1*time.Millisecond)
		fmt.Println("goroutine 1")
		wg.Done()
	}()

	go func() {
		time.Sleep(2*time.Millisecond)
		fmt.Println("goroutine 2")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("two routine down")
}
