package btree

import (
	"errors"
	"fmt"
	"reflect"
)

// tree node
type Node struct {
	value       interface{}
	left, right *Node
}

func (n *Node) Value() interface{} {
	return n.value
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) Traverse(order uint8, c container) {
	if n == nil {
		return
	}
	if order == PRE {
		*c = append(*c, n.value)
	}
	n.left.Traverse(order, c)
	if order == IN {
		*c = append(*c, n.value)
	}
	n.right.Traverse(order, c)
	if order == POST {
		*c = append(*c, n.value)
	}
}

// traverse binary tree with closure func
func (n *Node) TraverseByClosure(order uint8, f NodeFunc) {
	if n == nil {
		return
	}

	if order == PRE {
		f(n)
	}

	n.left.TraverseByClosure(order, f)

	if order == IN {
		f(n)
	}

	n.right.TraverseByClosure(order, f)

	if order == POST {
		f(n)
	}

}

type NodeFunc func(node *Node)

// binary tree
type Btree struct {
	root *Node // root node
	len  int   //number of nodes
	high int   //layers
}

func New() *Btree {
	return new(Btree).init()
}

func (b *Btree) init() *Btree {
	b.root = new(Node)
	b.len = 0
	b.high = 0
	return b
}

//numbers of nodes in binary tree
func (b *Btree) Len() int {
	return b.len
}

func (b *Btree) insert(n *Node) {

}

// layers of Btree
func (b *Btree) High() int {
	return b.high
}

func (b *Btree) Root() *Node {
	return b.root
}

// defied enum of traverse order
const (
	PRE  = iota //前序
	IN          //中序
	POST        //后序
)

// traverse binary tree
func (b *Btree) Traverse(order uint8) <-chan *Node {
	out := make(chan *Node)

	o := order
	go func(o uint8) {
		b.root.TraverseByClosure(o, func(n *Node) {
			out <- n
		})
		close(out)
	}(o)

	return out
}

//使用reflect 反射
func (b *Btree) Select(fieldName string, fieldValue interface{}, order uint8) (*Node, error) {
	out := b.Traverse(order)

	for node := range out {
		reflectNodeValue := reflect.ValueOf(node.Value())
		reflectNodeType := reflect.TypeOf(node.Value())
		//structField, _ := reflectNodeType.FieldByName("id")
		//fmt.Printf("%#v,%#v \r\n", structField.Name, structField)
		//fmt.Printf("%T,%#v \r\n", reflectNodeType.String(), reflectNodeType.String())
		//fmt.Printf("%T,%#v \r\n", reflectNodeType.Name(), reflectNodeType.Name())

		if !b.checkField(reflectNodeType, fieldName) {
			return nil, errors.New("invalid field:" + fieldName)
		}
		//fmt.Printf("%T,%v \r\n", reflectNodeValue.FieldByName(fieldName).Int(), reflectNodeValue.FieldByName(fieldName).Int())
		//fmt.Printf("%T,%#v \r\n", fieldValue, fieldValue)
		//fmt.Println(reflect.DeepEqual(reflectNodeValue.FieldByName(fieldName), fieldValue))

		switch reflectNodeValue.FieldByName(fieldName).Kind() {

		case reflect.Int:
			//把比较条件都转成int64的再进行比较
			if reflect.DeepEqual(reflectNodeValue.FieldByName(fieldName).Int(), reflect.ValueOf(fieldValue).Int()) {
				fmt.Println("deep ok")
				return node, nil
			}
		case reflect.String:
			if reflect.DeepEqual(reflectNodeValue.FieldByName(fieldName).String(), fieldValue) {
				fmt.Println("deep ok")
				return node, nil
			}
		}
	}
	return nil, errors.New("the corresponding value was not found")
}

//未使用反射
func (b *Btree) SelectById(id interface{}, f TypeAssertionFunc, order uint8) (*Node, error) {
	out := b.Traverse(order)

	for node := range out {
		if f(node.Value(), id) {
			fmt.Println("closure ok")
			return node, nil
		}
	}
	return nil, errors.New("the field value was not found")
}

type container *[]interface{}

type TypeAssertionFunc func(findValue interface{}, fieldValue interface{}) bool

func (b *Btree) checkField(rt reflect.Type, fieldName string) bool {
	if _, ok := rt.FieldByName(fieldName); ok {
		return true
	}
	return false
}
