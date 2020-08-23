package dac

import (
	"fmt"
)

//使用分治算法解决汉诺塔的移动问题
func HanoiTower(n int, a, b, c string) {
	//如果只有一个盘
	if n == 1 {
		fmt.Printf("第1个盘从%s->%s\r\n", a, c)
	} else {
		//	如果有n>=2的情况，把所有的盘看成是两个盘，盘1：最下面的一个盘，盘2:除去盘1上面的所有盘

		//先盘2从A移动到B，移动过程借助C
		HanoiTower(n-1, a, c, b)
		//然后把盘1从A移动到C
		fmt.Printf("第%d个盘从%s->%s\r\n", n, a, c)
		//最后把盘2从B移动到C，移动过程借助A
		HanoiTower(n-1, b, a, c)
	}
}
