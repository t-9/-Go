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

// Fいろは歌 はいろは歌文字列を作成します。
func Fいろは歌() *T文字 {
	return &T文字{
		中身: "色はにほへど　散りぬるを 我が世たれぞ　常ならむ 有為の奥山　今日越えて 浅き夢見じ　酔ひもせず",
	}
}

// F表示する は文字を表示します。
func (字 *T文字) F表示する() {
	fmt.Println(字.中身)
}
