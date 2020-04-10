package main

import (
	"algorithm/queue"
	"fmt"
	"log"
	"runtime"
)

func main() {
	//初始化队列
	q := queue.NewQueue(100)
	//创建10个goroutine send element to queue
	for i := 0; i < 10; i++ {
		go producer(i, q)
	}
	//创建了一个consumer 消费队列
	out := consumer(q)

	for {
		// 非阻塞调度
		select {

		case ele := <-out:
			fmt.Println(ele)
		}
	}
}

func producer(i int, q *queue.Queue) {
	for {

		err := q.Push(i)

		if err != nil {
			log.Println(err.Error())
			runtime.Gosched()
			continue
			// 坑爹的break 不要用
		}

	}
}

func consumer(q *queue.Queue) <-chan interface{} {
	out := make(chan interface{})

	go func(q *queue.Queue) {

		for {
			// 弹出一个element
			ele, err := q.Pop()
			// 日志
			log.Println("count:", q.GetCount())

			if err != nil {
				log.Println(err.Error())
				runtime.Gosched()
				continue
			}
			//add element to in channel
			out <- ele
		}
	}(q)

	return out
}
