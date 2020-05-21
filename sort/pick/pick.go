//选择排序思想
//有n个数，从第1个数开始找n个数中的最小数，找到以后将这个最小数和第一个数交换，也就是把最小数放到前面。
//下一轮，从第二个数开始，往后依次找最小数，找到以后就和第二个数交换，也就是把除第一个最小数外，找到最小的把它放在第二个位置。
//依次类推
//算法总结：
//1 进行n-1轮查找每轮最小数
//2 进行每轮的两两比较循环找到每轮最小数，和最小数所对应的索引
//3 完成一轮查找，找到最小数与本轮最前面的数进行交换，让小数排在最前面。

package pick

//选择排序
func Sort(s []int) []int {
	//比较轮数 n个数要比较n-1次
	for i := 0; i < len(s)-1; i++ {
		//初始化最小值索引
		t := i
		//假如第一个数最小
		m := s[i]
		//循环查找n个数中最小的
		for j := i + 1; j < len(s); j++ {
			if m > s[j] {
				m = s[j]
				//最小值变化，索引跟着相应变化
				t = j
			}
		}
		//最小值索引t不等于初始化索引i就说明最小值发生变化了，把最小值放在前面，前面个数换到后面
		if t != i {
			s[i], s[t] = m, s[i]
		}
	}
	return s
}
