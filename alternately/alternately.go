package main

import (
	"fmt"
	"sync"
)

func main() {

	//WaitGroup等待一组goroutine完成。
	//
	//主goroutine调用Add来设置要等待的goroutine的数量。
	//
	//然后每个goroutine运行并在完成时调用Done。
	//
	//与此同时，Wait可以用来阻塞直到所有的goroutine完成。
	//
	//重点 *WaitGroup不能在第一次使用后复制*。
	var wg sync.WaitGroup

	// 定义 pipeline 用于三个 goroutine 之间进行通信
	x := make(chan rune)
	y := make(chan rune)
	z := make(chan rune)
	//w := CreateWorker(&wg)
	o := consumer(&wg)
	wg.Add(3)

	n := 10

	go func() {
		//注意要先通知 WaitGroup goroutine 完成任务,然后再关闭通道
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
	wg.Wait()
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
		//如果在其他协程中调用了close(ch),那么此协程就会跳出 for range o 。这也就是 for range 的特别之处
		for v := range o {
			fmt.Println(string(v))
		}
	}()

	return o
}
