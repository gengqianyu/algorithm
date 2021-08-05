package containerring

type Ring struct {
	next, prev *Ring
	Value      interface{}
}

func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

// New creates a ring of n elements
//make vs new
//make的作用是初始化内置的数据结构，也就是，slice ，map，chan；
//new 的作用是根据传入的类型在堆上分配一片内存空间并返回指向这片内存空间的指针。
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}

	r := new(Ring)
	//用p作为当前节点，不断的链表后面添加元素
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	//linking the head and tail of a linked list
	// 最后收尾相连形成环形
	p.next = r
	r.prev = p
	return r
}

func (r *Ring) Next() *Ring {
	if r == nil {
		return r.init()
	}
	return r.next
}

func (r *Ring) Prev() *Ring {
	if r == nil {
		return r.init()
	}
	return r.prev
}

// Link 连接两个环形链表
//链表 r 与链表 s 是不同链表，则在 r 链表的后面链接 s 链表，否则删除相同部分
//这里需要注意的是两个环断开之后才能连接，所以是连接两个节点是双向四处位置
func (r *Ring) Link(s *Ring) *Ring {
	//获取r的下一个节点，以便在r和s连接之后，以便处理之前r的下一个节点的prev
	n := r.Next()
	if s != nil {
		//获取 s 环的上一个节点，以便 r 和 s 连接之后，以便处理
		p := s.Prev()
		//接第一个位置 双向
		r.next = s
		s.prev = r
		//接第二个位置 双向 此时n的prev还指向r呢, p的next还指向s呢
		n.prev = p
		p.next = n
	}
	return n
}

// UnLink 方法从下一个元素开始，移除链表连续n个元素
func (r *Ring) UnLink(n int) *Ring {
	if n < 0 {
		return nil
	}
	//整体思想就是找到r要移除元素的下一节点元素，然后与r进行相连

	return r.Link(r.Move(n + 1))
}

// Move 方法将 r 环的前面 n 个元素移到后面，例：0,1,2,3,4,5 => 3,4,5,0,1,2
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

//Len computes the number of element in ring r.
//it executes in time proportional to the number of elements.
func (r *Ring) Len() int {
	n := 0
	if r == nil {
		return n
	}
	//注意循环是从next开始的所以要提前设置n为1开始增长
	n = 1
	for p := r.Next(); p != r; p = p.next {
		n++
	}
	return n
}

// Do 方法以向前的顺序在环的每个元素上调用函数f。
//如果f改变* r，则Do的行为是不确定的。
func (r *Ring) Do(f func(interface{})) {
	if r != nil {
		f(r.Value)
		for p := r.Next(); p != r; p = p.next {
			f(p.Value)
		}
	}
}
