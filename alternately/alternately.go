package main

import (
	"fmt"
	"sync"
)

func main() {
	w := sync.WaitGroup{}
	x := make(chan rune)
	y := make(chan rune)
	z := make(chan rune)
	w.Add(3)
	n := 10
	go func() {
		defer w.Done()
		defer close(x)
		i := 0
		for {
			if i == n {
				break
			}
			e := <-x
			fmt.Println(string(e))
			y <- 'Y'
			i++
		}
	}()
	go func() {
		defer w.Done()
		defer close(y)
		i := 0
		for {
			if i == n {
				break
			}
			e := <-y
			fmt.Println(string(e))
			z <- 'Z'
			i++
		}
	}()
	go func() {
		defer w.Done()
		defer close(z)
		i := 0
		for {
			if i == n {
				break
			}
			e := <-z
			fmt.Println(string(e))
			//预防最后一次发送
			if i == n-1 {
				break
			}
			x <- 'X'
			i++
		}
	}()
	x <- 'X'
	w.Wait()
}
