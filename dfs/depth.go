/**
深度优先算法走迷宫
*/
package dfs

import (
	"algorithm/container/stack"
	"errors"
)

// defined point struct
type point struct {
	i, j int
}

// point move
func (p point) move(r point) point {

	return point{
		p.i + r.i,
		p.j + r.j,
	}
}

// 定位移动点在地图中的位置。
func (p point) position(m [][]int) int {
	//判断上下是否越界
	if p.i < 0 || p.i >= len(m) {
		return -1
	}
	//判断左右是否越界
	if p.j < 0 || p.j >= len((m)[p.i]) {
		return -1
	}
	return m[p.i][p.j]
}

//四个方向
var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

// 深度优先算法
// m 地图 s开始点 e终点
func DepthWalk(m [][]int, s point, e point) ([][]int, error) {
	//记录走过的位置
	steps := make([][]int, len(m))
	for i := range steps {
		steps[i] = make([]int, len(m[i]))
	}
	//定义一个栈
	Stack := stack.New()

	Stack.Push(stack.NewElement(s))
	for Stack.Len() > 0 {
		//取出栈当前元素
		element := Stack.Pop()
		//type assertion
		cur, ok := element.Value().(point)
		if !ok {
			return nil, errors.New("type error")
		}
		//到达终点break
		if cur == e {
			break
		}

		for _, dir := range dirs {
			//按照一个方向移动一步
			next := cur.move(dir)
			val := next.position(m)
			//如果不能走或者撞墙，进行下一个方向
			if val == -1 || val == 1 {
				continue
			}

			//查看这个位置是否走过
			val = next.position(steps)
			if val != 0 {
				continue
			}
			//如果走回原地 进行下一个方向
			if next == s {
				continue
			}
			//标识走过位置和步数
			steps[next.i][next.j] = cur.position(steps) + 1
			//将next入栈
			Stack.Push(stack.NewElement(next))
		}
	}
	return steps, nil

}
