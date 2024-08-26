package main

import (
	"fmt"
	"strconv"

	"golang.org/x/text/width"
)

// 住所を表す構造体
type Address struct {
	JichiCode string // 自治体コード
	MachiCode string // 町コード
	BanCode   string // 番地コード
	EdaCode   string // 枝番コード
	KoedaCode string // 小枝番コード
}

func (a Address) ToString() string {
	return fmt.Sprintf("%s,%s,%s,%s", a.MachiCode, width.Widen.String(a.BanCode), width.Widen.String(a.EdaCode), width.Widen.String(a.KoedaCode))
}

func (a Address) Next() *Address {
	// ALL 9の場合、nil を返す
	if a.BanCode == "99999" && a.EdaCode == "99999" && a.KoedaCode == "99999" {
		return nil
	}

	// 番地～枝番３までを連結する
	s := a.BanCode + a.EdaCode + a.KoedaCode
	// 数値に変換する
	l, _ := strconv.ParseUint(s, 10, 64)
	// 1を足す
	nextL := l + 1
	// 15桁の文字列に戻す
	nextS := fmt.Sprintf("%015d", nextL)
	// next Addressの組み立て
	return &Address{
		JichiCode: a.JichiCode,
		MachiCode: a.MachiCode,
		BanCode:   nextS[0:5],
		EdaCode:   nextS[5:10],
		KoedaCode: nextS[10:15],
	}
}

func (a Address) Previous() *Address {
	// ALL 0の場合、nil を返す
	if a.BanCode == "00000" && a.EdaCode == "00000" && a.KoedaCode == "00000" {
		return nil
	}

	// 番地～枝番３までを連結する
	s := a.BanCode + a.EdaCode + a.KoedaCode
	// 数値に変換する
	l, _ := strconv.ParseUint(s, 10, 64)
	// 1を足す
	preL := l - 1
	// 15桁の文字列に戻す
	preS := fmt.Sprintf("%015d", preL)
	// next Addressの組み立て
	return &Address{
		JichiCode: a.JichiCode,
		MachiCode: a.MachiCode,
		BanCode:   preS[0:5],
		EdaCode:   preS[5:10],
		KoedaCode: preS[10:15],
	}
}
