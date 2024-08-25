package main

import "fmt"

// 住所範囲を表す構造体
type Region struct {
	Start Address // 開始住所
	End   Address // 終了住所
}

func (r Region) ToString() string {
	return fmt.Sprintf("%s,０００００,%s,９９９９９", r.Start.ToString(), r.End.ToString())
}

func (r Region) EndRegion() *Region {
	// 終了住所が最大値の場合、nil を返す
	na := r.End.Next()
	if na == nil {
		return nil
	}

	return &Region{
		Start: *na,
		End: Address{
			JichiCode: na.JichiCode,
			MachiCode: na.MachiCode,
			BanCode:   "99999",
			EdaCode:   "99999",
			KoedaCode: "99999",
		},
	}
}

func (r Region) StartRegion() *Region {
	// 開始住所が最小値の場合、nilを返す
	pa := r.Start.Previous()
	if pa == nil {
		return nil
	}

	return &Region{
		Start: Address{
			JichiCode: pa.JichiCode,
			MachiCode: pa.MachiCode,
			BanCode:   "00000",
			EdaCode:   "00000",
			KoedaCode: "00000",
		},
		End: *pa,
	}
}

// 2つの地域が連続しているかをチェックする関数
func (r Region) CheckContinuity(pr Region) bool {
	return *pr.End.Next() == r.Start
}
