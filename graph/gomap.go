package graph

import (
	"algorithm/queue"
)

type GoMap struct {
	vertices []string    //顶点
	edges    [][]float64 //边 存储图对应的邻接矩阵
	number   int         //边的数量
	marks    []bool      //标记顶点是否已遍历
}

func (m *GoMap) init(n int) *GoMap {
	m.edges = make([][]float64, n)
	for i, _ := range m.edges {
		m.edges[i] = make([]float64, n)
	}
	m.marks = make([]bool, n)
	return m
}

func New(n int) *GoMap {
	return new(GoMap).init(n)
}

func (m *GoMap) Number() int {
	return m.number
}

// add vertex
func (m *GoMap) Add(v ...string) {
	m.vertices = append(m.vertices, v...)
}

//x，y是矩阵索引 v为边的权值
func (m *GoMap) SetEdge(x, y int, v float64) {
	m.edges[x][y] = v
	m.edges[y][x] = v
	m.number++
}

//直接设置
func (m *GoMap) SetEdges(v [][]float64) {
	m.edges = v
}
func (m *GoMap) Edges() [][]float64 {
	return m.edges
}

//获取顶点
func (m *GoMap) GetVertices() []string {
	return m.vertices
}

//得到第一个邻接节点的索引，//遍历矩阵一行中的所有节点，找到邻接
func (m *GoMap) FirstAdjacentIndex(i int) int {
	//以i为起点在矩阵中找第一个关系邻接节点
	for j := 0; j < len(m.vertices); j++ {
		if m.edges[i][j] == 1 {
			return j
		}
	}
	return -1
}

//根据前一个邻接节点的索引，来获取下一个邻接顶点
//i为起始节点，c为i的当前邻接节点
func (m *GoMap) NextAdjacentIndex(i, c int) int {
	//遍历矩阵一行中的所有节点，找到邻接
	for j := c + 1; j < len(m.vertices); j++ {
		if m.edges[i][j] == 1 {
			return j
		}
	}
	return -1
}

//图的深度优先遍历
func (m *GoMap) dfs(f func(string), i int) {
	//发送当前顶点
	//将节点设置成已访问
	f(m.vertices[i])
	m.marks[i] = true

	//查找顶点i的第一个邻接顶点w
	w := m.FirstAdjacentIndex(i)
	//判断是否有邻接顶点
	for w != -1 {
		//如果邻接顶点没有被访问过
		if m.marks[w] == false {
			m.dfs(f, w)
		}
		//如果w顶点已经被访问过
		//试着找w的下一个邻接顶点，如果找到继续，未找到跳出循环 进行下一轮从B开始的dfs
		w = m.NextAdjacentIndex(i, w)
	}
}

//遍历所有顶点，进行dfs
func (m *GoMap) DFS() <-chan string {
	out := make(chan string)
	go func() {
		for i := 0; i < len(m.vertices); i++ {
			//如果顶点访问过就继续往下一个顶点
			if m.marks[i] == true {
				continue
			}
			//从当前顶点开始进行dfs
			m.dfs(func(vertex string) {
				out <- vertex
			}, i)
			close(out)
		}
	}()
	return out
}

func (m *GoMap) bfs(f func(string), i int) {
	f(m.vertices[i])
	m.marks[i] = true

	//队列
	q := queue.New(-1)
	//添加首个元素
	q.Push(i)
	//如果队列不为空
	for q.Empty() == false {
		if h, ok := q.Pop(); ok {
			index := h.(int)
			w := m.FirstAdjacentIndex(index)
			//如果邻接结点存在
			for w != -1 {
				//	判断是否访问过，未访问就处理
				if m.marks[w] == false {
					f(m.vertices[w])
					m.marks[w] = true //标记已访问
					q.Push(w)
				}
				//以index为当前结点，以index的邻接节点w，去找下一个邻接节点，如果存在邻接就处理，一直在矩阵的一行中找不到邻接节点为止
				w = m.NextAdjacentIndex(index, w)
			}
		}

	}

}

func (m *GoMap) BFS() <-chan string {
	o := make(chan string)
	go func() {
		//从第一个顶点开始，处理所有顶点
		for i := 0; i < len(m.vertices); i++ {
			if m.marks[i] == true {
				continue
			}
			m.bfs(func(vertex string) {
				o <- vertex
			}, i)
		}
		close(o)
	}()
	return o
}
