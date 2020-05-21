package pick

//选择排序
func Sort(s []int) []int {
	//比较轮数 n个数要比较n-1次
	for i := 0; i < len(s)-1; i++ {
		t := i
		//假如第一个数最小
		m := s[i]
		//循环查找n个数中最小的
		for j := i + 1; j < len(s); j++ {
			if m > s[j] {
				m = s[j]
				t = j
			}
		}
		//只要t不等于i就说明最小值发生变化了
		if t != i {
			s[i], s[t] = m, s[i]
		}

	}
	return s
}
