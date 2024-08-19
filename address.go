package main

import (
	"fmt"

	"golang.org/x/text/width"
)

// 住所を表す構造体
type Address struct {
	MachiCode string // 町コード
	BanCode   string // 番地コード
	EdaCode   string // 枝番コード
	KoedaCode string // 小枝番コード
	Eda3Code  string // 枝番３コード
}

func (a Address) ToString() string {
	return fmt.Sprintf("%s,%s,%s,%s,%s", a.MachiCode, width.Widen.String(a.BanCode), width.Widen.String(a.EdaCode), width.Widen.String(a.KoedaCode), width.Widen.String(a.Eda3Code))
}
