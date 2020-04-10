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
	MaxSize   MaxSize       //最大容量
	Count     int           //元素个数
	Look      sync.Mutex    //互斥锁
	Container []interface{} //队列容器
}

func NewQueue(m MaxSize) *Queue {
	return &Queue{
		MaxSize: m,
		Count:   0,
	}
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
func (q *Queue) Pop() (interface{}, error) {
	q.Look.Lock()
	defer q.Look.Unlock()

	if q.Empty() {
		return nil, Error{message: "no element"}
	}

	head := q.Container[0]
	q.Container = q.Container[1:]
	q.Count--

	return head, nil
}

//is empty
func (q *Queue) Empty() bool {
	if q.Count > 0 {
		return false
	}
	return true
}

//is full
func (q *Queue) Full() bool {
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
