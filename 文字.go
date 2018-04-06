package main

import (
	"fmt"
)

type 文字 struct {
	中身 string
}

func 新しい文字(字 string) *文字 {
	return &文字{
		中身: 字,
	}
}

func (字 *文字) を表示する() {
	fmt.Println(字.中身)
}
