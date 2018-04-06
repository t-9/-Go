package 日本語

import (
	"fmt"
)

// T文字 は文字構造体です。
type T文字 struct {
	中身 string
}

// F新しい文字 は新しい文字を作成します。
func F新しい文字(字 string) *T文字 {
	return &T文字{
		中身: 字,
	}
}

// F表示する は文字を表示します。
func (字 *T文字) F表示する() {
	fmt.Println(字.中身)
}
