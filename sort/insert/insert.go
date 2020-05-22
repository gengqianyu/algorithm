//插入排序的基本思想
//把n个待排序的元素看成为一个有序队列和一个无序队列，
//开始时有序队列里只包含一个元素，无序队列中包含有n-1个元素，
//排序过程中每次从无序队列中取出一个元素，和有序队列中元素进行比较，
//将它插入到有序表中的适当位置
package insert

import "log"

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
			if e < p[i] {
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
		log.Println(p)

	}

	return p
}
