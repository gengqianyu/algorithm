//基数排序数据桶排序的扩展
//将整数位按位数切割成不同的数字，然后按每个位数分别比较
//将切片s的每一个元素e，从个位开始到最高位，将位数和桶内编号相等的元素放到一个桶内，再将桶里的元素重新放到s中，
//位数=桶的编号，桶的编号是从0到9，s中的最大数m有多少位，就进行多少轮以上操作。
package radix

import "strconv"

func Sort(s []int) []int {
	//find max number
	m := 0
	for _, e := range s {
		if e > m {
			m = e
		}
	}
	// create a barrel container
	b := make([][]int, 10)
	//m的位数
	m = len(strconv.Itoa(m))
	//除数
	n := 1
	//定义一个s指针
	var p int
	//放入取出的轮数
	for j := 0; j < m; j++ {
		// 往桶里添加元素
		for i := 0; i < len(s); i++ {
			//取模获取一个个/白/千/万...位数
			d := (s[i] / n) % 10
			b[d] = append(b[d], s[i])
		}
		//重置每一轮s的指针
		p = 0
		//将桶里的数放回s
		for i, e := range b {
			for _, v := range e {
				s[p] = v
				p++
			}
			//重置桶
			b[i] = []int{}
		}
		//重置除数
		n *= 10
	}
	return s
}
