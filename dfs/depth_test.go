package dfs

import (
	"fmt"
	"os"
	"testing"
)

func TestDepthWalk(t *testing.T) {
	marshalTests := []struct {
		name string
		src  string
	}{
		{name: "maze", src: "./maze.txt"},
	}

	for _, tt := range marshalTests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := ReadMaze(tt.src)
			if err != nil {
				panic(err)
			}
			//ShowMaze(m)

			//起始/终点位置
			s, e := point{0, 0}, point{len(m) - 1, len(m[0]) - 1}

			steps, err := DepthWalk(m, s, e)
			ShowMaze(steps)
		})
	}
}

func ShowMaze(m [][]int) {
	for i := range m {
		for j := range m[i] {
			fmt.Printf("%3d", m[i][j])
		}
		fmt.Println()
	}
}

func ReadMaze(src string) ([][]int, error) {
	//打开一个文件
	f, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// add buffer caches 这里还不能加buffer 加了读入地图有错误
	//reader := bufio.NewReader(f)

	var row, col int
	// 从reader中按照format格式扫描一组数据，并赋值给row，col，所以要用地址
	fmt.Fscanf(f, "%d %d", &row, &col)
	//fmt.Println(row, col)
	m := make([][]int, row)

	for i := range m {
		//创建二维数组
		m[i] = make([]int, col)
		for j := range m[i] {
			//在reader中按照 format格式扫描元素，并赋值
			fmt.Fscanf(f, "%d", &m[i][j])
		}
	}
	return m, nil
}
