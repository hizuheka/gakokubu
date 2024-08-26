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

func (r Record) ToString() string {
	return fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s", r.region.ToString(), r.gakuKubun, r.gakuCode, r.sakujoFlag, r.updateYMD, r.jichiCode, r.updateYMDHMS)
}

// 学校区分が小学校の場合 true を、そうでない場合 false を返す
func (r Record) IsShogaku() bool {
	return r.gakuKubun == "1"
}

// 2つの地域の間の欠けている地域を見つける関数
func (r Record) FindMissingRegions(pr Record) []Region {
	// 直前のレコードと町コードが異なる場合は、直前のレコード～ALL 9と、ALL 0～自身のレコードを返す
	if pr.region.End.MachiCode != r.region.Start.MachiCode {
		mr := make([]Region, 0)
		if er := pr.region.EndRegion(); er != nil {
			mr = append(mr, *er)
		}
		if sr := pr.region.StartRegion(); sr != nil {
			mr = append(mr, *sr)
		}
		return mr
	}

	// 直前のレコードと続いている場合は、nilを返す
	if r.region.CheckContinuity(pr.region) {
		return nil
	}

	// 直前のレコードと続いていない場合、直前のレコードの次の住所～今のレコードの前の住所を返す
	return []Region{
		{
			Start: *pr.region.End.Next(),
			End:   *r.region.Start.Previous(),
		},
	}
}

// 1行の入力を Record 構造体に変換する関数
func createRecord(line string) (Record, error) {
	items := strings.Split(line, ",")
	if len(items) != 16 {
		return Record{}, fmt.Errorf("入力ファイルの形式が誤っています。想定している項目数は 16 です。(len(line)=%d", len(line))
	}

	// 住所範囲
	r := Region{
		Start: Address{
			JichiCode: items[14],
			MachiCode: items[0],
			BanCode:   width.Fold.String(items[1]),
			EdaCode:   width.Fold.String(items[2]),
			KoedaCode: width.Fold.String(items[3]),
		},
		End: Address{
			JichiCode: items[14],
			MachiCode: items[5],
			BanCode:   width.Fold.String(items[6]),
			EdaCode:   width.Fold.String(items[7]),
			KoedaCode: width.Fold.String(items[8]),
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

func createDummyRecord(kubun, jichiCode string) Record {
	return Record{
		region:       Region{},
		gakuKubun:    kubun,
		gakuCode:     "00999",
		sakujoFlag:   "0",
		updateYMD:    "99999999",
		jichiCode:    jichiCode,
		updateYMDHMS: "99999999999999",
	}
}
