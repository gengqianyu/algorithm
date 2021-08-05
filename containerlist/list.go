package containerlist

//Element is an element of a linked list
type Element struct {
	//双向链接的元素链表中的下一个和上一个指针。
	//为了简化实现，内部将链表l实现为环，这样＆l.root既是最后一个链表元素的下一个元素（l.Back（）），也是第一个链表元素的前一个元素（l.Front （））。
	next, prev *Element
	//此节点元素所属的链表
	list *List
	// Value 存储元素的值
	Value interface{}
}

// Next 方法返回下一个链表元素或nil。
func (e *Element) Next() *Element {
	if p := e.next; p != nil && p != &e.list.root {
		return p
	}
	return nil
}

//Prev 返回上一个链表元素或nil
func (e *Element) Prev() *Element {
	if p := e.prev; p != nil && p != &e.list.root {
		return p
	}
	return nil
}

//List表示双链表。
//List的零值是可以使用的空链表。
type List struct {
	// 链表根节点
	root Element
	// 当前链表的元素个数，不包含链表根节点
	len int
}

//初始化和清空链表l
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.root.list = l

	l.len = 0
	return l

}

func New() *List {
	return new(List).Init()
}

// 惰式初始化l
func (l *List) lazyInit() {
	//如果root节点的next是nil就说明没初始化过
	if l.root.next == nil {
		l.Init()
	}
}

// Len returns a number of element of list l
func (l *List) Len() int {
	return l.len
}

// Front returns first element of list l or nil if list is empty
func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Back returns last element of list l or nil if list is empty
func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

// insert inserts e after at
// increments l.len
// returns e
func (l *List) insert(e, at *Element) *Element {
	n := at.next //获取at.next 元素

	at.next = e //让at.next指向新元素e
	e.prev = at // 设置e的上一个元素为at

	e.next = n // 设置e的下一个元素，n
	n.prev = e // 设置n的上一个元素为 新元素e

	e.list = l // 设置e的所属链表

	l.len++
	return e
}

// InsertValue is a convenience wrapper for Insert
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// PushFront inserts an new element e with a value v at the front for list l and returns e.
func (l *List) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// PushBack inserts an new element e with an value v the back for list l and return e.
func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// Remove Removes an e from its list ,decrements l.len , return e
func (l *List) remove(e *Element) *Element {

	e.next.prev = e.prev
	e.prev.next = e.next

	e.next = nil
	e.prev = nil
	e.list = nil

	l.len--
	return e
}

func (l *List) Remove(e *Element) interface{} {

	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}

	//连接e移除位置
	e.next.prev = e.prev
	e.prev.next = e.next
	//获取at next
	n := at.next

	at.next = e
	e.prev = at

	e.next = n
	n.prev = e

	l.len--
	return e
}

// InsertBefore 方法 在 mark 之前插入 element
func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

// 在mark之后插入element
func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

// 移动节点到标记节点之前
func (l *List) MoveBefore(e, mark *Element) {
	if e == mark || e.list != l || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

// 移动节点到标记节点之后
func (l *List) MoveAfter(e, mark *Element) {
	if e == mark || e.list != l || mark.list != l {
		return
	}
	l.move(e, mark)
}

//把节点移动到一个元素位置
func (l *List) MoveToFront(e *Element) {
	if e.list != l || e == &l.root {
		return
	}
	l.move(e, &l.root)
}

// 把节点移动到最后一个元素位置
func (l *List) MoveToBack(e *Element) {
	if e.list != l || e == l.root.prev {
		return
	}
	l.move(e, l.root.prev)
}
