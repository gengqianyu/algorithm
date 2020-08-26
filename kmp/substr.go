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
