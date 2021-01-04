//Binary Search
package binary

//s slice for search
//v value for match 匹配
//第一步：找到中间索引
//第二步：比较查找的值和索引对应的值
//第三步：如果查找的值比索引对应值小，改变切片右边界索引为，中间索引减1。
//		如果查找的值比索引对应值大，改变切片左边界索引为，中间索引加1。
//		继续递归二分查找
//第四步：如果查找值等于索引对应值，那么返回找到的索引。
func Search(s []int, l, r, v int) int {
	if l > r {
		return -1
	}

	// defined mid
	m := (l + r) >> 1
	//如果中间值比v大就向左递归
	if s[m] > v {
		//注意这里的m，比完以后，下一轮m就不用比了，
		//如果取的查找的是两边是数，m-1,m+1可以防止就剩最后两个数的时候出现死归，
		return Search(s, l, m-1, v)
	}
	//如果中间值比v小就向右递归
	if s[m] < v {
		return Search(s, m+1, r, v)
	}
	//默认s[m]=v
	return m
}
//非递归形式实现
func Search2(s []int, v int) int {
	l, r := 0, len(s)-1
	//只要l小于r就可以一直找
	for l <= r {
		//找到中间索引
		m := (l + r) >> 1

		//比较
		//向左边查找
		if s[m] > v {
			r = m - 1
			continue
		}
		//向右边查找
		if s[m] < v {
			l = m + 1
			continue
		}

		return m
	}
	return -1
}
