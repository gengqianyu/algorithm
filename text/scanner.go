package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	// src是我们想要标记化的输入。
	src := []byte("cos(x) + 1i*sin(x) // Euler")

	// 初始化扫描仪。
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions相对于fset
	file := fset.AddFile("", fset.Base(), len(src)) // 注册输入“file”
	s.Init(file, src, nil /* 没有错误处理程序 */, scanner.ScanComments)

	// 重复调用Scan会产生输入中找到的标记序列。
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
