package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/width"
)

// レコードを表す構造体
type Record struct {
	region       Region // 住所範囲
	gakuKubun    string // 学校区分
	gakuCode     string // 学校コード
	sakujoFlag   string // 削除フラグ
	updateYMD    string // 更新日
	jichiCode    string // 自治体コード
	updateYMDHMS string // 更新日
}

func (re Record) ToString() string {
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s", re.region.ToString(), re.gakuKubun, re.gakuCode, re.sakujoFlag, re.updateYMD, re.jichiCode, re.updateYMDHMS)
}

// 学校区分が小学校の場合 true を、そうでない場合 false を返す
func (re Record) IsShogaku() bool {
	return re.gakuKubun == "1"
}

// 1行の入力を Record 構造体に変換する関数
func createRecord(line string) (Record, error) {
	items := strings.Split(line, ",")
	if len(items) == 16 {
		return Record{}, fmt.Errorf("入力ファイルの形式が誤っています。想定している項目数は 16 です。(len(line)=%d", len(line))
	}

	// 住所範囲
	r := Region{
		Start: Address{
			MachiCode: items[0],
			BanCode:   width.Fold.String(items[1]),
			EdaCode:   width.Fold.String(items[2]),
			KoedaCode: width.Fold.String(items[3]),
			Eda3Code:  width.Fold.String(items[4]),
		},
		End: Address{
			MachiCode: items[5],
			BanCode:   width.Fold.String(items[6]),
			EdaCode:   width.Fold.String(items[7]),
			KoedaCode: width.Fold.String(items[8]),
			Eda3Code:  width.Fold.String(items[9]),
		},
	}

	record := Record{
		region:       r,
		gakuKubun:    items[10],
		gakuCode:     items[11],
		sakujoFlag:   items[12],
		updateYMD:    items[13],
		jichiCode:    items[14],
		updateYMDHMS: items[15],
	}

	return record, nil
}
