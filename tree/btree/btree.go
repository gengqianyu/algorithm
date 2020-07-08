//binary tree
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

func (n *Node) Traverse(l int, c *[]interface{}, order uint8) {

	if n == nil {
		return
	}

	if order == PRE {
		*c = append(*c, n.value)
	}
	n.left.Traverse(l, c, order)
	if order == IN {
		*c = append(*c, n.value)
	}
	n.right.Traverse(l, c, order)
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

// delete node
func (n *Node) Delete(fieldName string, fieldValue interface{}, isDelete *bool) {
	if n == nil {
		return
	}

	switch reflect.ValueOf(n.Value()).FieldByName(fieldName).Kind() {

	case reflect.Int:
		//从n的左子节点查找满足条件的node删除
		if n.left != nil && reflect.ValueOf(n.left.Value()).FieldByName(fieldName).Int() == reflect.ValueOf(fieldValue).Int() {
			n.left = nil
			*isDelete = true
			return
		}
		//从n的右子节点查找满足条件的node删除
		if n.right != nil && reflect.ValueOf(n.right.Value()).FieldByName(fieldName).Int() == reflect.ValueOf(fieldValue).Int() {
			n.right = nil
			*isDelete = true
			return
		}
	case reflect.String:
		if n.left != nil && reflect.ValueOf(n.left.Value()).FieldByName(fieldName).String() == fieldValue {
			n.left = nil
			*isDelete = true
			return
		}
		if n.right != nil && reflect.ValueOf(n.right.Value()).FieldByName(fieldName).String() == fieldValue {
			n.right = nil
			*isDelete = true
			return
		}
	}

	n.Left().Delete(fieldName, fieldValue, isDelete)
	n.Right().Delete(fieldName, fieldValue, isDelete)
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
	b.len = 7
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
func (b *Btree) Select(fieldName string, fieldValue interface{}, order uint8) (interface{}, error) {
	out := b.Traverse(order)

	for node := range out {
		reflectValue := reflect.ValueOf(node.Value())
		reflectType := reflect.TypeOf(node.Value())

		if !b.checkField(reflectType, fieldName) {
			return nil, errors.New("invalid field:" + fieldName)
		}

		switch reflectValue.FieldByName(fieldName).Kind() {

		case reflect.Int:
			//把比较条件都转成int64的再进行比较
			if reflect.DeepEqual(reflectValue.FieldByName(fieldName).Int(), reflect.ValueOf(fieldValue).Int()) {
				fmt.Println("deep ok")
				return node.Value(), nil
			}
		case reflect.String:
			if reflect.DeepEqual(reflectValue.FieldByName(fieldName).String(), fieldValue) {
				fmt.Println("deep ok")
				return node.Value(), nil
			}
		}
	}
	return nil, errors.New("the corresponding value was not found")
}

//未使用反射 使用函数式
func (b *Btree) SelectById(id interface{}, f TypeAssertionFunc, order uint8) (interface{}, error) {
	out := b.Traverse(order)

	for node := range out {
		if f(node.Value(), id) {
			fmt.Println("closure ok")
			return node.Value(), nil
		}
	}
	return nil, errors.New("the field value was not found")
}

type TypeAssertionFunc func(findValue interface{}, fieldValue interface{}) bool

func (b *Btree) checkField(rt reflect.Type, fieldName string) bool {
	if _, ok := rt.FieldByName(fieldName); ok {
		return true
	}
	return false
}

func (b *Btree) GetAll() []interface{} {
	c := new([]interface{})
	b.root.Traverse(b.len, c, 0)
	return *c
}

func (b *Btree) Delete(fieldName string, fieldValue interface{}) (bool, error) {
	if b.Len() == 0 {
		return false, errors.New("btree no records")
	}
	if b.checkField(reflect.TypeOf(b.root.Value()), fieldName) == false {
		return false, errors.New("invalid field:" + fieldName)
	}
	//判断是否删除的是根节点
	reflectValue := reflect.ValueOf(b.Root().Value())

	switch reflectValue.FieldByName(fieldName).Kind() {
	case reflect.Int:
		//把比较条件都转成int64的再进行比较
		if reflect.DeepEqual(reflectValue.FieldByName(fieldName).Int(), reflect.ValueOf(fieldValue).Int()) {
			b.root = nil
			b.len--
			return true, nil
		}
	case reflect.String:
		if reflect.DeepEqual(reflectValue.FieldByName(fieldName).String(), fieldValue) {
			b.root = nil
			b.len--
			return true, nil
		}
	}

	var isDelete bool

	b.Root().Delete(fieldName, fieldValue, &isDelete)

	if isDelete {
		b.len--
		return true, nil
	} else {
		return false, errors.New("no record found")
	}

}
