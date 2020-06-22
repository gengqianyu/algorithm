package hash

import (
	"log"
	"reflect"
	"strconv"
	"testing"
)

type Personnel struct {
	number int
	name   string
	age    uint8
}

func TestTable_Put(t *testing.T) {
	marshalTests := []Personnel{
		Personnel{number: 123, name: "gengqianyu", age: 32},
		Personnel{number: 456, name: "zhangyanlin", age: 33},
		Personnel{number: 789, name: "gengzinan", age: 2},
	}
	hashTable := New()
	for _, tt := range marshalTests {
		t.Run(tt.name, func(t *testing.T) {
			hashTable.Put(strconv.Itoa(tt.number), tt)
		})
	}

	p, ok := hashTable.get(strconv.Itoa(789)).(Personnel)

	if !ok {
		t.Error("Type error the get Personnel!")
	}

	if reflect.DeepEqual(p, marshalTests[2]) {
		log.Println(p.number, "=>", p.name)
	} else {
		t.Errorf("expected the name:%s,actual:%s", (marshalTests[2]).name, p.name)
	}

	//e := hashTable.Delete("789").(Personnel)
	//log.Println(e)
	//a := hashTable.Delete("789")
	//log.Println(a)
	//log.Println(hashTable.Size())

	iter := hashTable.KeySet().iterator(hashTable.table)
	log.Println("----------------------------------------")
	for iter.HasNext() {
		key := iter.Next().(string)
		p := hashTable.get(key).(Personnel)
		log.Println(p.number, "=>", p.name)
	}

	log.Println(len(hashTable.table))
}
