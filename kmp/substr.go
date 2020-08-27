package kmp

//求一个字符串最长不重复子串
func MaxLenSubstr(s string) (int, string) {
	//defined max length of substr
	maxLen := 0
	// defined start index of substr
	startIndex := 0
	// defined the substring
	var substr string

	//defined a map,used to store the last position of the character
	lastContainer := make(map[rune]int)

	for index, char := range []rune(s) {
		//	get an index of character from lastContainer
		lastIndex, ok := lastContainer[char]
		// if last index exist,and last index is greater than or equal start index ,readjust the staring index of the substring
		if ok && lastIndex >= startIndex {
			startIndex = lastIndex + 1
		}

		//	calculate the length of the substring
		// if the substring length greater than max length
		// readjust the max length of the substring
		if substrLen := (index + 1) - startIndex; substrLen > maxLen {
			maxLen = substrLen
			substr = string([]rune(s)[startIndex : index+1])
		}

		lastContainer[char] = index
	}
	return maxLen, substr
}

//暴力匹配算法
func IndexStrings(s, substr string) int {
	//将s和substr转成字符切片
	strings := []rune(s)
	substring := []rune(substr)
	//定义索引指针 分别指向 strings 和substring的起始位置
	var i, j int

	for i < len(strings) && j < len(substring) {
		//匹配到第一个就，就进行下一次匹配
		if strings[i] == substring[j] {
			i++
			j++
			continue
		}
		//没有匹配到，字符串索引i往前回退j-1个位置，子串索引j回退到初始位置
		i = i - (j - 1)
		j = 0
	}
	//如果找到子串，返回子串，首次出现的索引
	if j != 0 {
		return i - j
	}
	return -1
}

//kmp搜索算法
func KMP(s, substr string) int {
	//获取子串的部分匹配表
	p := PartialMach(substr)

	ss := []rune(s)
	chars := []rune(substr)

	for i, j := 0, 0; i < len(ss); i++ {
		//匹配不到的情况,就一直调整j的索引位置
		//kmp核心算法
		for j > 0 && ss[i] != chars[j] {
			j = p[j-1]
		}
		//匹配到，接着匹配下一个字符
		if ss[i] == chars[j] {
			j++
		}

		//找到返回索引
		if j == len(chars) {
			return i - j
		}
	}
	return -1
}

//获取字符串的部分匹配表 前缀和后缀 共同元素的个数
func PartialMach(s string) []int {
	//将字符串转化为切片
	ss := []rune(s)
	//创建一个p切片 保存字符串的部分匹配值
	p := make([]int, len(s))
	//如果只有一个字符，这个字符的部分匹配值就是0
	p[0] = 0
	//从第二个字符开始 i=1 i为字符索引
	//j代表字符的部分匹配值
	for i, j := 1, 0; i < len(ss); i++ {

		//当ss[i] != ss[j]，需要从p[j-1]获取新的匹配值
		//字符匹配值大于0，
		for j > 0 && ss[i] != ss[j] {
			//kmp的核心，匹配值计算公式
			j = p[j-1]
		}

		//满足条件部分匹配值加1
		if ss[i] == ss[j] {
			j++
		}

		p[i] = j
	}

	return p
}
