package interpolation

func Search(s []int, l, r, v int) int {

	// 未找到，查找的值明显不在s中 直接返回
	if l > r || v < s[l] || v > s[r] {
		return -1
	}

	m := l + (v-s[l])/(s[r]-s[l])*(r-l)
	if v > s[m] { //向右递归
		return Search(s, m+1, r, v)
	}

	if v < s[m] { //向左递归
		return Search(s, l, m-1, v)
	}
	return m
}
