package main

import (
	"fmt"
	"sync"
)

func main() {

	//WaitGroup 等待一组 goroutine 完成。
	//
	//主 goroutine 调用 Add 来设置要等待的 goroutine 的数量。
	//
	//然后每个 goroutine 运行并在完成时调用 Done。
	//
	//与此同时，Wait 可以用来阻塞当前 goroutine 直到所有的 goroutine 完成。
	//
	//重点* WaitGroup 不能在第一次使用后复制 *。执行一下三行会发生错误
	//wgp := sync.WaitGroup{}
	//yawg := wgp
	//fmt.Println(wgp, yawg)

	var wg sync.WaitGroup

	// 定义 pipeline 用于三个 goroutine 之间进行通信
	x := make(chan rune)
	y := make(chan rune)
	z := make(chan rune)
	//w := CreateWorker(&wg)
	o := consumer(&wg)
	wg.Add(3) //更新计数器

	n := 10

	go func() {
		//注意要先通知 WaitGroup， goroutine 完成任务,然后再关闭通道
		//注意 defer 执行顺序是先进后出，因此执行顺序是 先 wg.Done() 后 close(x)
		defer close(x)
		defer wg.Done()
		i := 0
		for {
			if i == n {
				break
			}
			if i == 0 {
				//w.pipeline <- 'X'
				o <- 'X'
			} else {
				//w.pipeline <- <-x
				o <- <-x
			}
			y <- 'Y'
			i++
		}
	}()

	go func() {
		defer close(y)
		defer wg.Done()
		i := 0
		for {
			if i == n {
				break
			}
			//fmt.Println(string(<-y))
			//w.pipeline <- <-y
			o <- <-y
			z <- 'Z'
			i++
		}
	}()

	go func() {
		defer close(z)
		defer wg.Done()
		i := 0
		for {
			//fmt.Println(string(<-z))
			//w.pipeline <- <-z
			o <- <-z
			//预防最后一次发送
			if i == n-1 {
				close(o)
				break
			}
			x <- 'X'
			i++
		}
	}()
	//sync.WaitGroup.Wait 会在计数器大于 0 并且不存在等待的 Goroutine 时，调用 runtime.sync_runtime_Semacquire 让调用 wg.Wait函数的 goroutine 陷入睡眠。
	//main goroutine 在这里等待，当调用计数器归零，即所有任务都执行完成时，才会通过 sync.runtime_Semrelease 唤醒处于等待状态的 main Goroutine。
	wg.Wait()

	//sync.WaitGroup 必须在 sync.WaitGroup.Wait 方法返回之后才能被重新使用；
	//sync.WaitGroup.Done 只是对 sync.WaitGroup.Add 方法的简单封装，我们可以向 sync.WaitGroup.Add 方法传入任意负数（需要保证计数器非负）快速将计数器归零以唤醒等待的 Goroutine；
	//可以同时有多个 Goroutine 等待当前 sync.WaitGroup 计数器的归零，这些 Goroutine 会被同时唤醒；也就是可以有多个 goroutine 同时执行 sync.WaitGroup.wait() 方法
}

type worker struct {
	pipeline chan rune
	done     func() //函数式编程
}

func CreateWorker(group *sync.WaitGroup) worker {

	w := worker{

		pipeline: make(chan rune),

		done: func() {
			group.Done()
		},
	}

	go func(w worker) {
		defer close(w.pipeline)
		defer w.done()
		for v := range w.pipeline {
			fmt.Println(string(v))
		}

	}(w)

	return w
}

func consumer(group *sync.WaitGroup) chan rune {
	o := make(chan rune)
	group.Add(1)
	go func() {
		defer group.Done()
		//注意 range 一个 channel，会一直阻塞当前 goroutine 协程，
		//如果在其他协程中调用了 close(ch),那么此协程就会跳出 for range o 。这也就是 for range 的特别之处。
		for v := range o {
			fmt.Println(string(v))
		}
	}()

	return o
}
