//插入排序的基本思想
//把n个待排序的元素看成为一个有序队列和一个无序队列，
//开始时有序队列里只包含一个元素，无序队列中包含有n-1个元素，
//排序过程中每次从无序队列中取出一个元素，和有序队列中元素进行比较，
//将它插入到有序表中的适当位置
package insert

//{3, 0, 7, 8, 9, 4, 1, 2, 6, 5},
//效率不如下面的优化版本，而且还占内存
func Sort(s []int) []int {
	//初始化有序列表
	var p []int

	for _, e := range s {
		//第一次循环直接在有序队列中插入第一个值
		if len(p) == 0 {
			p = append(p, e)
			continue
		}
		//默认插入位置是1
		t := 1
		//找插入位置
		for i := 0; i < len(p); i++ {
			//找到插入位置就跳出去插入,因为t是插入位置所以是索引+1
			if e <= p[i] {
				t = i + 1
				break
			}
			t++
		}
		//log.Fatal(t)

		// p[b:] 拿到slice中从b位置以后的element，不包含b。p[:b] 拿到b位置以前的element，包含b ，
		// p[a:b] a到b所有element ，同理不包含a，包含b 也就是"后包前不包"
		//注意：保存后部剩余元素，必须新建一个临时切片, b=p[t-1:]不行，因为b和p公用底层数组会发生错误
		b := append([]int{}, p[t-1:]...)
		p = append(append(p[:t-1], e), b...)
		//log.Println(p)
	}

	return p
}

//插入排序简化版本
//时间复杂度O(n²)
//D:\go\algorithm\sort\insert>go test --bench=".*" --benchmem -v
//goos: windows
//goarch: amd64
//pkg: algorithm/sort/insert
//BenchmarkSort-8           149996              7228 ns/op               0 B/op          0 al
//locs/op
//PASS
//ok      algorithm/sort/insert   1.679s

//{3, 0, 7, 8, 9, 4, 1, 2, 6, 5},
func Sort2(s []int) []int {
	//一共需要比较len(s)-1次
	for i := 1; i < len(s); i++ {

		t := s[i]  //定义待插入的数 从索引为1的第二个数
		o := i - 1 //定义要插入的位置  无条件默认插入位置为索引为0的第一个位置

		//给t找到一个要插入的位置
		for o >= 0 && t < s[o] {
			s[o+1] = s[o]
			o--
		}
		//因为for循环中最后o--是希望在o的前一个位置插入，如果条件不成立，是要插在o这个位置的所以得加回来
		s[o+1] = t
	}
	return s
}
