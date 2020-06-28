package tree

type Interface interface {
	Len() int
	High() int
	traverse(uint8) <-chan interface{}
}
