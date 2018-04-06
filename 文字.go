package main

import (
	"fmt"
)

type 文字 struct {
	中身 string
}

func (字 *文字) 表示する() {
	fmt.Println(字.中身)
}
