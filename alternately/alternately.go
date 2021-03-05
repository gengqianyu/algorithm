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
			if i == 0 {
				fmt.Println(string('X'))
			} else {
				fmt.Println(string(<-x))
			}
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
			fmt.Println(string(<-y))
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
			fmt.Println(string(<-z))
			//预防最后一次发送
			if i == n-1 {
				break
			}
			x <- 'X'
			i++
		}
	}()
	w.Wait()
}
