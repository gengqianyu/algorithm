package queue

import (
	"sync"
)

//custom error
type Error struct {
	message string
}

//implement error interface
func (q Error) Error() string {
	return q.message
}

type MaxSize int

type Queue struct {
	MaxSize   MaxSize       //最大容量 -1表示不限制
	Count     int           //元素个数
	Look      sync.Mutex    //互斥锁
	Container []interface{} //队列容器
}

func (q *Queue) init(m MaxSize) *Queue {
	q.MaxSize = m
	return q
}

func New(m MaxSize) *Queue {
	return new(Queue).init(m)
}

//push
func (q *Queue) Push(e interface{}) error {
	q.Look.Lock()
	defer q.Look.Unlock()

	if q.Full() {
		return Error{message: "queue overflow"}
	}

	q.Container = append(q.Container, e)
	q.Count++
	return nil
}

//pop up
func (q *Queue) Pop() (interface{}, bool) {
	q.Look.Lock()
	defer q.Look.Unlock()

	if q.Empty() {
		return nil, false
	}

	head := q.Container[0]
	q.Container = q.Container[1:]
	q.Count--

	return head, true
}

//is empty
func (q *Queue) Empty() bool {
	if q.Count > 0 {
		return false
	}
	return true
}

//is full
//-1 表示不限制
func (q *Queue) Full() bool {
	if q.MaxSize == -1 {
		return true
	}
	if len(q.Container) <= int(q.MaxSize) {
		return false
	}
	return true
}

//是否存在某个元素
func (q *Queue) Exists() bool {
	return true
}

// get counter
func (q *Queue) GetCount() int {
	return q.Count
}
