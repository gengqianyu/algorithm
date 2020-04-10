/**
关于堆栈自己的理解不一定对但是有一定道理
程序执行过程中内存使用情况
内存会被分为三块区域
代码/常量区：用于存放程序代码和常量
栈区：用与存放编译后的机器指令，取址执行（弹栈）
堆区：用于存放程序中定义的变量，对象。
编译的过程就是将程序的代码逻辑，编译成机器指令。
静态语言为什么执行速度快，是因为运行时，直接将机器指令加载到内存，省去了编译过程。
*/

package stack

//defined a element struct
type Element struct {
	value interface{}
	prev  *Element
}

// Prev get the prev element
func (e *Element) Prev() *Element {
	if p := e.prev; p != nil {
		return p
	}
	return nil
}

func (e *Element) Value() interface{} {
	return e.value
}

// defined an stack
type Stack struct {
	top *Element
	len int
}

// init the stack
func (s *Stack) Init() *Stack {
	s.top = nil
	s.len = 0
	return s
}

func New() *Stack {
	return new(Stack).Init()
}

func (s *Stack) Pop() *Element {
	if s.Len() == 0 {
		return nil
	}
	t := s.top
	s.top = t.prev
	s.len--
	return t
}

func (s *Stack) Push(e *Element) {
	e.prev = s.top
	s.top = e
	s.len++
}

func (s *Stack) Peek() *Element {
	return s.top
}

func (s *Stack) Len() int {
	return s.len
}
