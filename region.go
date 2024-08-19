package main

import "fmt"

// 住所範囲を表す構造体
type Region struct {
	Start Address // 開始住所
	End   Address // 終了住所
}

func (r Region) ToString() string {
	return fmt.Sprintf("%s,%s", r.Start.ToString(), r.End.ToString())
}
