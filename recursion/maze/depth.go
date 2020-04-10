/**
深度优先算法走迷宫
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	m, err := ReadMaze("./recursion/maze/maze.txt")
	if err != nil {
		panic(err)
	}
	ShowMaze(m)
	//行走路径
	var p []point
	//起始位置
	s := point{0, 0}
	e := point{len(m) - 1, len(m[0]) - 1}
	DepthWalk(m, s, e, &p)
}

func ReadMaze(src string) ([][]int, error) {
	f, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// add buffer caches 这里还不能加buffer 加了读入地图有错误
	//reader := bufio.NewReader(f)
	var row, col int
	// read data from reader 会给row，col赋值，所以要用地址
	fmt.Fscanf(f, "%d %d", &row, &col)
	//fmt.Println(row, col)
	m := make([][]int, row)

	for i := range m {
		m[i] = make([]int, col)
		for j := range m[i] {
			fmt.Fscanf(f, "%d", &m[i][j])
		}
	}
	return m, nil
}

func ShowMaze(m [][]int) {
	for i := range m {
		for j := range m[i] {
			fmt.Printf("%3d", m[i][j])
		}
		fmt.Println()
	}
}

// defined point struct
type point struct {
	i, j int
}

// point move
func (p point) move(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// 定位移动点在地图中的位置。
func (p point) position(m [][]int) (int, bool) {
	//判断上下是否越界
	if p.i < 0 || p.i >= len(m) {
		return 0, false
	}
	//判断左右是否越界
	if p.j < 0 || p.j >= len(m[p.i]) {
		return 0, false
	}
	return m[p.i][p.j], true
}

// 深度优先算法
// m 地图 s开始点 e终点
func DepthWalk(m [][]int, s point, e point, p *[]point) ([]point, error) {
	if s == e {
		return *p, nil
	}

	if  {

	}
	//向下走
	n := s.move(point{1, 0})
	v, ok := n.position(m)
	if ok && v != 1 {
		DepthWalk(m, n, e, p)
	}
	//向左走
	n = s.move(point{0, -1})
	v, ok = n.position(m)
	if ok && v != 1 {
		DepthWalk(m, n, e, p)
	}
	//向上走
	n = s.move(point{-1, 0})
	v, ok = n.position(m)
	if ok && v != 1 {
		DepthWalk(m, n, e, p)
	}
	//向右走
	n = s.move(point{0, 1})
	v, ok = n.position(m)
	if ok && v != 1 {
		DepthWalk(m, n, e, p)
	}

}

func Move()  {
	//向下走
	n := s.move(point{1, 0})
	v, ok := n.position(m)
	if ok && v != 1 {
		DepthWalk(m, n, e, p)
	}
}