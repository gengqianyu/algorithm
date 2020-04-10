package containerlist

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"testing"
)

type Hero struct {
	name     string // 名字
	nickname string // 昵称
	no       int    //排位
}

func (h *Hero) Name() string {
	return h.name
}

func (h *Hero) Nickname() string {
	return h.nickname
}

func (h *Hero) No() int {
	return h.no
}

func TestList(t *testing.T) {
	var tests []*Hero
	tests = append(tests, &Hero{
		name:     "宋江",
		nickname: "及时雨",
		no:       1,
	})
	tests = append(tests, &Hero{
		name:     "卢俊义",
		nickname: "玉麒麟",
		no:       2,
	})

	// new a list
	l := New()

	for _, tt := range tests {
		// insert  e after l.root
		l.PushFront(tt)
	}

	if l.Len() != 2 {
		t.Errorf("expected:%v,actual:%v", 2, l.Len())
	}

	f := l.Front()

	//类型断言的意义，type assertion 可以将interface{} type，转回Hero Struct type
	if g, ok := f.Value.(*Hero); ok {
		log.Printf("%T,%v", g, g.Name())
	}

	b := l.Back()
	if v, ok := b.Value.(*Hero); ok {
		log.Printf("%T,%v", v, v.Name())
	}

	c := l.InsertAfter(&Hero{name: "吴用", nickname: "智多星", no: 3}, b)
	if v, ok := c.Value.(*Hero); ok {
		log.Printf("%T,%v", v, v.Name())
	}

	l.InsertBefore(&Hero{name: "公孙胜", nickname: "入云龙", no: 4}, f)

	fe := l.Front()
	if v, ok := fe.Value.(*Hero); !ok {
		t.Errorf("type error:expected:%s,actual:%T", "*Hero", v)
	}

	if f := l.Front(); f.Value.(*Hero).name != "公孙胜" {
		t.Errorf("expected:%s,actual:%s", "公孙胜", f.Value.(*Hero).name)
	}

	// 查找单链表中的倒数第2个节点 上面list的顺序是 公孙胜，卢俊义，松江，吴用
	ie, err := FindLastIndexNode(l, 2)
	if err != nil {
		t.Error(err.Error())
	}

	v, ok := ie.Value.(*Hero)
	if !ok {
		t.Errorf("type error:expected %s,actual%T", "*Hero", v)
	}
	if v.Name() != "宋江" {
		t.Errorf("expected:%s,actual:%s", "宋江", v.Name())
	}

	//反转list
	l = ReverseList(l)
	for n := l.Front(); n != nil; n = n.Next() {
		v := n.Value.(*Hero)
		log.Println(v.Name())
	}

	//倒序打印
	ReversePrint(l)
}

//排序测试
func ExampleList_InsertAfter() {
	tests := []*Hero{
		&Hero{
			name:     "吴用",
			nickname: "智多星",
			no:       3,
		},
		&Hero{
			name:     "宋江",
			nickname: "及时雨",
			no:       1,
		},
		&Hero{
			name:     "公孙胜",
			nickname: "入云龙",
			no:       4,
		},
		&Hero{
			name:     "卢俊义",
			nickname: "玉麒麟",
			no:       2,
		},
		&Hero{
			name:     "林冲",
			nickname: "豹子头",
			no:       5,
		},
	}
	// new an list l
	l := New()
	for _, tt := range tests {
		//add element to front of list l ,if l is empty
		if l.Len() == 0 {
			l.PushFront(tt)
			continue
		}

		//初始化辅助节点
		t := l.Front()
		//按顺序插入 注意用从root节点开始。
		for n := &l.root; n != nil; n = n.Next() {
			//如果后一个位置是root 就在root之前插入，也就是插入到最后
			if n.next == &l.root {
				t = &l.root
				break
			}
			// 通过当前辅助节点，查找下一个节点，下一个节点的值比要插入节点的值大，则当前辅助节点就是要找的位置
			v, ok := n.Next().Value.(*Hero)
			if !ok {
				log.Fatal("类型错误")
			}
			//if v.No() == tt.No() {
			//	log.Fatal("节点已存在")
			//}

			//找到链表中要插入的位置 1234567 只要插入位置的后一个位置的no比要插入的element no大就在它之前插入
			if tt.No() < v.No() {
				t = n.Next()
				break
			}

		}
		// 在标记位插入element
		l.InsertBefore(tt, t)
	}

	//log.Printf("l len:%d", l.Len())

	for n := l.Front(); n != nil; n = n.Next() {
		if v, ok := n.Value.(*Hero); ok {
			fmt.Printf("%s\n", v.Name())
		}

	}

	// Output:
	// 宋江
	// 卢俊义
	// 吴用
	// 公孙胜
	// 林冲
}

// 查找单链表中的倒数第k个节点
func FindLastIndexNode(l *List, i int) (*Element, error) {
	if l.Len() == 0 || i == 0 || i > l.Len() {
		return nil, errors.New("not find element of index is " + strconv.Itoa(i))
	}
	c := 0
	var e *Element
	for n := l.Front(); true; n = n.Next() {
		//倒数第i个节点，就是第l.Len()-i 个节点
		if c == (l.Len() - i) {
			e = n
			break
		}
		c++
	}
	return e, nil
}

//翻转List
func ReverseList(l *List) *List {
	rev := New().Init()
	for n := l.Front(); n != nil; n = n.Next() {
		// type assertion
		v := n.Value.(*Hero)
		//如果没有节点直接添加
		if rev.Len() == 0 {
			rev.PushFront(v)
			continue
		}

		//找到第一个节点 在第一个之前添加
		b := rev.Front()
		rev.InsertBefore(v, b)
	}
	// 修改原始list
	l.root.next = rev.root.next
	l.root.prev = rev.root.prev
	//回收rev
	defer func() { rev = nil }()
	return l
}

// 可以利用栈数据结构，将各个节点压入到栈中
func ReversePrint(l *List) {
	var s []*Element

	for n := l.Front(); n != nil; n = n.Next() {
		s = append(s, n)
	}

	sl := len(s) - 1
	for i := range s {
		e := s[sl-i]
		c := e.Value.(*Hero)
		log.Println(c.Name())
	}
}
