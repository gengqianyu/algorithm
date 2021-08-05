package containerring

import (
	"fmt"
	"testing"
)

func TestRing_Move(t *testing.T) {

}

func TestRing_Do(t *testing.T) {
	tests := []string{
		"宋江",
		"卢俊义",
		"吴用",
		"公孙胜",
		"林冲",
	}
	//t.Log(len(tests))
	r := New(len(tests))

	for i := 0; i < len(tests); i++ {
		r.Value = tests[i]
		r = r.Next() //正好循环一圈r又指向了第一个节点
	}
	//在函数中调用子程序，以下为例，首先会将r.do在内存中的指针，push到stack中。
	//等joseph执行return返回时，再将stack中记录的r.do在内存中指针pop出来，接着执行下面的指令。
	//解决约瑟夫问题
	r = joseph(r, 2, 3)

	r.Do(func(v interface{}) {
		fmt.Println(v.(string))
	})

}

//r是环形链表
//s表示起始位置
//j表示数到几
func joseph(r *Ring, s, j int) *Ring {
	if r.Len() < 2 {
		return r
	}
	r = r.Move(s - 1)

	i := 1 //标识数到的数
	for {
		if r == r.Next() {
			break
		}
		if i == j {
			r = r.Prev()
			r.UnLink(1) //注意这里是从下一个元素开始删除后面n个
			i = 1
			r = r.Next()
			continue
		}
		r = r.Next()
		i++
	}
	return r
}
