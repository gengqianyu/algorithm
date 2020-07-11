// sorting with large top heap
package htree

type HTree struct {
	s []int
}

func (h *HTree) init(s []int) *HTree {
	h.s = s
	return h
}

func New() *HTree {
	return new(HTree)
}

// func of heap sort
func (h *HTree) Sort() []int {
	//将无序切片序列构建成一个堆，根据升序降序需求，选择大顶堆或小顶堆
	//			4
	//		6		8
	//	  5	  9
	//从第一个非叶子节点开始构建(局部)大顶堆
	//n个节点 n/2-1的索引就是第一个非叶子节点6，第二次循环就是i等于0根节点4作为父节点构建，完成以后就是一个(全局)大顶堆
	for i := (len(h.s) >> 1) - 1; i >= 0; i-- {
		h.heap(i, len(h.s))
	}

	//交换
	for j := len(h.s) - 1; j > 0; j-- {
		//将大顶堆顶的元素与末尾元素交换，将最大元素沉到切片的末尾
		h.s[j], h.s[0] = h.s[0], h.s[j]
		//重新调整使，其满足大顶堆定义，然后继续交换堆顶元素与当前末尾元素，反复执行调整+交换步骤，直到整个序列有序
		h.heap(0, j)
	}
	return h.s
}

// create a large top heap
// i表示当前父节点，
// l表示对多少个节点进行调整
func (h *HTree) heap(i, l int) {
	//取出当非叶子节点(父节点)的值，保存在临时变量
	t := h.s[i]

	//遍历当前节点的子树 形成一个大顶堆

	//(i << 1) + 1 就是i的左子节点
	for j := (i << 1) + 1; j < l; j = (j << 1) + 1 {
		//当i的左子节点小于右子节点
		//也就是说在左右子节点中找到一个最大的节点
		if j+1 < l && h.s[j] < h.s[j+1] {
			// j指向右子节点
			j++
		}
		//如果子节点中的最大值比父节点大，那么父、子节点交换位置
		//如果子节点中的最大的等于小于父节点，就不需要调整了直接退出。调整次序是从左到右，由下到上，如果有这情况说明调整过了已经，不是太好理解
		if h.s[j] <= t {
			break
		}
		h.s[i] = h.s[j] //父节点值等于子节点中的最大值

		// 这里的子节点中最大值的j位置的值，不能马上换成父节点i位置的值，如果子节点中j位置，下面还有子树，那就得把j当作父节点，继续去调整子树成为大顶堆
		i = j //将 j位置作为父节点，继续调整子树成为大顶堆
	}
	//调整完毕以后将最开始的父节点放入，调整后的节点那个位置，也就是最后j的位置，上面用i=j了
	h.s[i] = t
}
