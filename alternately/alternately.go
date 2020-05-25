package main

import (
	"fmt"
	"sync"
)

func main() {
	w := sync.WaitGroup{}
	x := make(chan string)
	y := make(chan string)
	z := make(chan string)
	w.Add(3)
	n := 10
	go func() {
		i := 0
		for {
			if i == n {
				break
			}
			e := <-x
			fmt.Println(e)
			y <- "Y"
			i++
		}
		w.Done()
		//close(x) done了就不用close了
	}()
	go func(y chan string) {
		i := 0
		for {
			if i == n {
				break
			}
			e := <-y
			fmt.Println(e)
			z <- "Z"
			i++
		}
		w.Done()

	}(y)
	go func() {
		i := 0
		for {
			if i == n {
				break
			}
			e := <-z
			fmt.Println(e)
			//预防最后一次发送
			if i < n-1 {
				x <- "X"
			}
			i++
		}
		w.Done()
	}()
	x <- "X"
	w.Wait()
}
