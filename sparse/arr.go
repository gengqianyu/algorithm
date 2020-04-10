package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var chessArr1 [11][11]int
	chessArr1[1][2] = 1
	chessArr1[2][3] = 2
	chessArr1[4][5] = 2
	//打印棋盘
	for _, row := range chessArr1 {
		for _, item := range row {
			fmt.Printf("%d\t", item)
		}
		fmt.Println()
	}

	// 二维转稀疏，稀疏转文件
	err := ChessToSparse(chessArr1)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("----------------------------------------------------")

	// 文件转稀疏，稀疏转二维
	chessArr2, err := SparseToChess()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, row := range chessArr2 {
		for _, item := range row {
			fmt.Printf("%d\t", item)
		}
		fmt.Println()
	}

}

const filename = "map.data"

// 二维转稀疏,稀疏转文件
func ChessToSparse(chessArr [11][11]int) error {
	// 统计非零个数
	sum := 0
	//遍历棋盘读非零数num
	for _, row := range chessArr {
		for _, item := range row {
			if item != 0 {
				sum++
			}
			fmt.Printf("%d\t", item)
		}
		fmt.Println()
	}

	//初始化稀疏数组
	var sparseArr [][3]int
	sparseArr = append(sparseArr, [3]int{11, 11, sum})
	//二维数组转稀疏切片
	for rIndex, row := range chessArr {
		for lIndex, item := range row {
			if item != 0 {
				sparseArr = append(sparseArr, [3]int{rIndex + 1, lIndex + 1, item})
			}
		}
	}
	//创建文件
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	//创建一个buffer writer
	w := bufio.NewWriter(file)

	for _, row := range sparseArr {
		//写文件
		fmt.Fprintf(w, "%d\t%d\t%d\n", row[0], row[1], row[2])
	}
	// 将缓冲区的内容，刷到文件中，重置缓冲
	defer w.Flush()

	return nil
}

func SparseToChess() ([11][11]int, error) {
	//文件转稀疏数组
	var chessArr [11][11]int

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return chessArr, err
	}

	//初始化缓冲区
	r := bufio.NewReader(file)

	// 定义稀疏数组
	var sparseArr [][3]int

	// 读文件初始化稀疏切片
	var a, b, c int
	for {
		// 扫描一行
		n, err := fmt.Fscanf(r, "%d\t%d\t%d\n", &a, &b, &c)

		if (n == 0) && (err != nil) {
			log.Printf("%T\n", err)
			log.Println(err.Error())
			break
		}
		fmt.Println(a, b, c)
		sparseArr = append(sparseArr, [3]int{a, b, c})
	}
	// 去除第一行 说明信息
	sparseArr = sparseArr[1:]

	//  还原二维数组
	for _, row := range sparseArr {
		chessArr[row[0]-1][row[1]-1] = row[2]
	}
	return chessArr, nil
}
