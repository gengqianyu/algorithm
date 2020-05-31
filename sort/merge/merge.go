//归并排序(merge-sort)是利用归并的思想实现的排序方法，
//该算法采用经典的分治策略，(分治方法将问题分成一些小的问题然后递归求解，而治的阶段则将分的阶段得到的各答案，"合并"在一起，即分而治之)
package merge

//多路channel 两两合并 如果参数是n个连续一样的值，可以用这种 参数... 的形式，接收到一个切片中
//merge(mergen(c1,c2,c3,c4),mergen(c5,c6,c7,c8))
//merge(merge(mergen(c1,c2),mergen(c3,c4)),merge(mergen(c5,c6),mergen(c7,c8)))
//merge(merge(merge(mergen(c1),mergen(c2)),merge(mergen(c3),mergen(c4))),merge(merge(mergen(c5),mergen(c6)),merge(mergen(c7),mergen(c8))))
//mergen(c1)=>c1
//merge(merge(merge(c1,c2),merge(c3,c4)),merge(merge(c5,c6),merge(c7,c8)))
//merge(c1,c2)=>out1
//merge(merge(out1,out2),merge(out3,out4))
//merge(out1,out2)=>out12
//merge(out12,out34)
//merge(out12,out34)=>out1234
//上一层的merge中goroutine 向out发送数据阻塞等待，下一层merge中goroutine in接收数据 ，上层的out就是下层的in
func MergeN(ins ...<-chan int) <-chan int {
	if len(ins) == 1 {
		return ins[0]
	}
	m := len(ins) / 2
	return Merge(MergeN(ins[:m]...), MergeN(ins[m:]...))
}

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		//使用非阻塞的方式从两个channel里接收数据，
		// v1：表示接收到的数据，未接收到数据时，v1为通道类型的零值。
		// ok1：表示是否接收到数据
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		//如果能接到数据就执行循环体
		for ok1 || ok2 {
			//如果从in2没接收到数据，或者从in1接收到数据，从nv1<v2 就将V1放入 out
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				//接下一组数
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		//如果没有数据了发送了就关闭通道
		close(out)

	}()
	return out
}

func Sort(s []int) []int {
	//分组
	var ins []<-chan int
	//这里需注意如果用make定义,下面循环得用ins[i]赋值，如果再用append相当于增加了ins的长度
	//ins:=make([]<-chan int,len(s))
	for _, e := range s {
		ins = append(ins, MakeIn(e))
	}
	//合并
	out := MergeN(ins...)

	i := 0
	for e := range out {
		s[i] = e
		i++
	}
	return s
}

func MakeIn(i int) <-chan int {
	out := make(chan int)
	go func(i int) {
		out <- i
		close(out)
	}(i)
	return out
}
