package hash

import (
	"hash/fnv"
	"log"
	"reflect"
)

// defined a hash table
type Table struct {
	// array of save key-value
	table []*Entry
	// count hash table number of elements
	count int
	// 阈值，用于判断是否需要调整Hashtable的容量（threshold = 容量*加载因子）
	threshold int
	//加载因子
	loadFactor float64
	// hashtable 被改变次数
	modCount int
	//key 集合
	keySet *KeySet
}

// new a hashtable (factory)
func New() *Table {
	return new(Table).init()
}

//初始化hashtable
func (t *Table) init() *Table {
	//table初始容量16
	t.table = make([]*Entry, 16)
	t.loadFactor = 0.75
	t.threshold = (int)(16 * t.loadFactor)
	return t
}

//add the “key-value” to Hashtable
func (t *Table) Put(key string, value interface{}) interface{} {
	// make hash code
	hash := fnv.New32()
	if _, e := hash.Write([]byte(key)); e != nil {
		log.Fatal(e)
	}
	hashCode := hash.Sum32()
	index := int(hashCode&0x7FFFFFFF) % len(t.table)

	// 若Hashtable中已存在键为key的键值对，则用新的value替换旧的value
	for e := t.table[index]; e != nil; e = e.next {
		if e.hash == hashCode && reflect.DeepEqual(e.key, key) {
			o := e.value
			e.value = value
			return o
		}
	}

	//容量达到阀值，hashtable扩容
	if t.count >= t.threshold {
		t.reHash()
		//重新计算index
		table := t.table
		index = int(hashCode&0x7FFFFFFF) % len(table)
	}
	//select Entry
	e := t.table[index]
	//add key=>value
	t.table[index] = &Entry{
		next:  e,
		hash:  hashCode,
		key:   key,
		value: value,
	}
	t.count++
	return nil
}

//根据key获取元素
func (t *Table) get(key string) interface{} {
	// get hash code
	hash := fnv.New32()
	if _, e := hash.Write([]byte(key)); e != nil {
		log.Fatal(e)
	}
	code := hash.Sum32()
	index := int(code&0x7FFFFFFF) % len(t.table)
	// 找到key对应的Entry(链表)，然后在链表中找出 哈希值和键值 与key都相等的元素
	for e := t.table[index]; e != nil; e = e.next {
		if e.hash == code && reflect.DeepEqual(e.key, key) {
			return e.value
		}
	}
	return nil
}

// delete element from hashtable
func (t *Table) Delete(key string) interface{} {
	// get hash code
	hash := fnv.New32()
	if _, e := hash.Write([]byte(key)); e != nil {
		log.Fatal(e)
	}
	hashCode := hash.Sum32()
	index := int(hashCode&0x7FFFFFFF) % len(t.table)
	// 找到“key对应的Entry(链表)”
	// 然后在链表中找出要删除的节点，并删除该节点。
	var prev *Entry
	prev = nil

	for e := t.table[index]; e != nil; e = e.next {

		if e.hash == hashCode && reflect.DeepEqual(e.key, key) {
			t.modCount++

			if prev != nil {
				//指针操作 用前一个节点指针，指向当前节点的下一个节点 相当于删除了当前节点指针
				prev.next = e.next
			} else {
				t.table[index] = e.next
			}

			t.count--

			oldValue := e.value
			e.value = nil
			return oldValue
		}
		//相当于拷贝指针
		prev = e
	}
	return nil
}

//清空hashtable
func (t *Table) Clear() {
	t.modCount++
	for i, _ := range t.table {
		t.table[i] = nil
	}
	t.count = 0
}

// 调整Hashtable的长度，将长度变成原来的2倍
// (01) 将“旧的Entry数组”赋值给一个临时变量。
// (02) 创建一个“新的Entry数组”，并赋值给“旧的Entry数组”
// (03) 将“Hashtable”中的全部元素依次添加到“新的Entry数组”中
func (t *Table) reHash() {
	//旧hashtable
	oldCapacity := len(t.table)
	oldTable := t.table
	//新hashtable
	newCapacity := oldCapacity * 2
	newTable := make([]*Entry, newCapacity)

	t.modCount++
	//更新阀值
	t.threshold = int(float64(newCapacity) * t.loadFactor)
	// 更改hashtable为新hashtable
	t.table = newTable
	// 将旧hashtable 元素导入到新hashtable中
	for i := oldCapacity - 1; i > 0; i-- {
		for e := oldTable[i]; e != nil; e = e.next {
			//单链表节点赋给临时变量
			temp := e
			// 重新计算
			index := int(e.hash&0x7FFFFFFF) % newCapacity
			//将temp的next指针指向新hashtable对应的entry
			temp.next = newTable[index]
			//将新链表赋给对应的entry 相当于在新的 entry链表头插入了一个元素
			newTable[index] = temp
		}
	}
}

//get hashtable number of elements
func (t *Table) Size() int {
	return t.count
}

//判断hashtable是否为空
func (t *Table) isEmpty() bool {
	return t.count == 0
}

//判断hashtable是否包含"值(value)"
func (t *Table) containsValue(value interface{}) bool {
	for _, entry := range t.table {
		for e := entry; e != nil; e = e.next {
			if reflect.DeepEqual(e.value, value) {
				return true
			}
		}
	}
	return false
}

//判断hashtable是否包含key
func (t *Table) containsKey(key string) bool {
	hash := fnv.New32()
	if _, e := hash.Write([]byte(key)); e != nil {
		log.Fatal(e)
	}
	hashCode := hash.Sum32()
	index := int(hashCode&0x7FFFFFFF) % len(t.table)

	for e := t.table[index]; e != nil; e = e.next {
		if e.hash == hashCode && reflect.DeepEqual(e.key, key) {
			return true
		}
	}
	return false
}

// 返回“所有key”的枚举对象
func (t *Table) Keys() []string {
	return nil
}

// 返回“所有Element”的枚举对象
func (t *Table) Elements() []string {
	return nil
}

func (t *Table) KeySet() *KeySet {
	if t.keySet == nil {
		t.keySet = new(KeySet)
	}
	return t.keySet
}

//获取一个
func getIterator(table []*Entry, t int) *Enumerator {
	return MakeEnumerator(table, t, true)
}

// Hashtable的Key的Set集合。
type KeySet struct {
	count int
}

// 获取集合个数
func (k *KeySet) Siz() int {
	return k.count
}

func (k *KeySet) iterator(table []*Entry) *Enumerator {
	return getIterator(table, KEYS)
}

//定义枚举
const (
	KEYS = iota
	VALUES
	ENTRIES
)

//Enumerator的作用是提供了“通过elements()遍历Hashtable的接口” 和 “通过entrySet()遍历Hashtable的接口”。 实现了“Iterator接口”。
type Enumerator struct {
	//指向hashtable的table
	table []*Entry
	//hashtable的总大小
	index int

	entry *Entry
	//类型
	genre int

	// Enumerator是 “迭代器(Iterator)” 还是 “枚举类(Enumeration)”的标志
	// iterator为true，表示它是迭代器；否则，是枚举类。
	iterator bool
}

//factory of Enumerator
func MakeEnumerator(table []*Entry, genre int, iterator bool) *Enumerator {
	return new(Enumerator).initEnumerator(table, genre, iterator)
}

//初始化 Enumerator
func (i *Enumerator) initEnumerator(table []*Entry, genre int, iterator bool) *Enumerator {
	i.table, i.genre, i.iterator = table, genre, iterator
	i.index = len(i.table)
	return i
}

// 迭代器Iterator的判断是否存在下一个元素
// 实际上，它是调用的hasMoreElements()
func (i *Enumerator) HasNext() bool {
	return i.hasMoreElements()
}

// 遍历table 从切片的末尾向前查找，直到找到不为nil的Entry。
func (i *Enumerator) hasMoreElements() bool {
	//copy i.entry i.index i.table
	e := i.entry
	index := i.index
	table := i.table
	for e == nil && index > 0 {
		index--
		e = table[index]
	}
	i.entry = e
	i.index = index
	return e != nil
}

// 迭代器获取下一个元素
// 实际上，它是调用的nextElement()
func (i *Enumerator) Next() interface{} {
	return i.nextElement()
}

// 获取下一个元素
// 注意：从hasMoreElements() 和nextElement() 可以看出Hashtable的elements()遍历方式
// 首先，从后向前的遍历table切片。table切片的每个节点都是一个单向链表(Entry)。
// 然后，依次向后遍历单向链表Entry。
func (i *Enumerator) nextElement() interface{} {
	et := i.entry
	index := i.index
	table := i.table

	for et == nil && index > 0 {
		et = table[index]
	}
	i.entry = et
	i.index = index

	if et != nil {
		e := i.entry
		i.entry = e.next
		return e.key
	}
	return nil
}

func (i *Enumerator) Current() interface{} {
	return nil
}

func (i *Enumerator) Rewind() {

}

func (i *Enumerator) Key() int {
	return i.index
}

// defined an Entry of a single list
type Entry struct {
	next  *Entry
	hash  uint32
	key   string
	value interface{}
}

func (e *Entry) Key() string {
	return e.key
}

func (e *Entry) Value() interface{} {
	return e.value
}

func (e *Entry) SetValue(v interface{}) {
	e.value = v
}

func (e *Entry) Next() *Entry {
	return e.next
}

func (e *Entry) Hash() uint32 {
	return e.hash
}

//定义迭代器接口
type Iterator interface {
	//验证当前是否还有下一个元素
	HasNext() bool
	//获取当前一个元素
	Current() interface{}
	//下一个元素
	Next() interface{}
	//重置迭代器
	Rewind()
	//迭代器的指针位置
	Key() int
}
