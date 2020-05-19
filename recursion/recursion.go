package main

import (
	"errors"
	"fmt"
)

/**
递归需要遵守的重要规则
1执行方法时，创建一个新的受保护的独立空间（栈空间）
2方法的局部变量是独立的，不会相互影响。比如n变量
3如果方法使用了引用类型变量（比如数组指针），就会共享该引用类型数据
4递归必须向退出递归条件逼近，否则就是无限递归，最后栈溢出，
5当一个方法执行完毕，或遇到return就会返回，遵守谁调用就将结果返回给谁，同时当方法执行完毕或者返回时，该方法也执行完毕
*/

func main() {
	test(5)
	fmt.Println(factorial(2))
	fmt.Println(fib(6)) // 1 1 2 3 5 8
	fmt.Println(pinBall(1, 100.00, 0.00))
	fmt.Println(add(5))
	r := []string{
		"宋江",
		"卢俊义",
		"吴用",
		"公孙胜",
		"林冲",
	}
	r, _ = joseph(r, 2, 3)
	fmt.Println(r)

}

//output
//n= 2
//n= 3
//n= 4
//n= 5
//注意调用规则第一条，每次调用test函数时会去栈上分配一个独立的内存空间，像一个压栈过程。将新指令压入前一个test函数指令之上
//独立空间中包含所有test函数指令，是独立的。所以自然包含打印指令。
//所以打印的是四个数不是一个。又因为栈的执行顺序永远都是从栈顶开始。所以打印指令输出是逆序的
func test(n int) {

	if n > 2 {
		test(n - 1)
	}

	fmt.Println("n=", n)
}

//阶乘的运算过程
//第一步 factorial(3-1)*3 => factorial(2)*3
//note: factorial(2) return factorial(2-1)*2
//第二部 factorial(2-1)*2*3 => factorial(1)*2*3
//第三部 factorial(1)直接返回1 表达式成为 1*2*3
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return factorial(n-1) * n
}

//斐波那契序列
//第一步 fib(4-1)+fib(4-2) =>fib(3)+fib(2)
//note: fib(2)返回1
//第二步 fib(3-1)+fib(3-2)+1=>fib(2)+fib(1)+1
//第三步 fib(2)+fib(1)+1=>1+1+1
// output 3
func fib(n int) int {
	if n == 0 {
		return 0

	}
	if n < 3 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

//小球跌落算法 100米高空落下每次反弹一半。10的距离
func pinBall(n int, h float64, l float64) float64 {
	if n <= 10 {
		return pinBall(n+1, h/2, l+(h+h/2))
	} else {
		return l
	}
}

//0到100加法
func add(n int) int {
	if n == 1 {
		return 1
	}
	return n + add(n-1)
}

//约瑟夫 后期可以尝试递归可能会简单
//r 人数切片，s开始位置，j步数
func joseph(r []string, s, j int) ([]string, error) {
	if len(r) < s {
		return nil, errors.New("s 溢出")
	}
	if len(r) == 1 {
		return r, nil
	}
	//重置slice的数序，重置开始位置
	if s != 1 {
		r = append(r[s-1:], r[:s-1]...)
		fmt.Println(r)
	}
	// 注意切片的切分于合并一定是要以位置取做切割不要用索引 i为切片初始位置1
	i := 1
	// c为步数也是从第一步开始
	c := 1
	for {
		if len(r) == 1 {
			break
		}
		if i > len(r) {
			i = 1
		}
		if c == j {
			//删除slice中的第i个元素
			r = append(r[i:], r[:i-1]...)
			fmt.Println(r)
			break
		}
		i++
		c++
	}
	return joseph(r, 1, j)
}
