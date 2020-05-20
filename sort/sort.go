package main

import (
	"algorithm/sort/bubble"
	"fmt"
)

func main() {
	s := []int{3, 5, 7, 8, 9, 4, 1, 2, 6, 0}
	//冒泡排序
	bs := bubble.BubbleSort(s)
	for _, e := range bs {
		fmt.Println(e)
	}
}
