package dyanmic

import (
	"fmt"
)

//使用动态规划解决背包问题
func Knapsack() {
	//	物品的重量
	w := []int{1, 4, 3}
	//	物品的价值
	p := []int{1500, 3000, 2000}

	//	二维数组v表示背包，v[i][j] 表示在前i个商品中，能装入容量为j的背包中，物品的最大价值
	var v [4][5]int
	// 二维数组s 记录商品放入背包的情况
	var s [4][5]int

	//动态规划处理，背包问题,不处理背包容量为0，和没有物品的情况
	//v[i][j]表示一个装入策略
	for i := 1; i < len(v); i++ {
		for j := 1; j < len(v[i]); j++ {
			//当前物品的重量大于总容量j时，就直接将上一个商品的装入策略，放入当前装入策略中
			if w[i-1] > j {
				v[i][j] = v[i-1][j]
				continue
			}
			//当前物品重量小于等于总容量j时，就从上一个商品的装入策略，和当前商品价值+剩余空间能装入商品的价值，两者中选一个最大的
			//v[i][j] = int(math.Max(float64(v[i-1][j]), float64(p[i-1]+v[i-1][j-w[i-1]])))
			if v[i-1][j] < p[i-1]+v[i-1][j-w[i-1]] {
				v[i][j] = p[i-1] + v[i-1][j-w[i-1]]
				//记录装入状态
				s[i][j] = 1
			} else {
				v[i][j] = v[i-1][j]
			}
		}
	}

	//打印背包
	for _, c := range v {
		for _, l := range c {
			fmt.Print(l)
		}
		fmt.Println()
	}

	//输出我们最后放入的商品
	c := len(s) - 1
	l := len(s[0]) - 1

	for c > 0 && l > 0 {
		//从最后一个商品开始找
		if s[c][l] == 1 {
			fmt.Printf("第%d个商品放入背包\r\n", c)
			//让l等于剩余容量，得出剩余空间可以放的商品
			l -= w[c-1]
		}
		//倒数第二个商品
		c--
	}
}
