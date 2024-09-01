package main

import (
	"sort"
)

type Records []Record

func (r Records) Sort() {
	sort.Slice(r, func(i, j int) bool {
		if r[i].region.Start.MachiCode == r[j].region.Start.MachiCode {
			if r[i].region.Start.BanCode == r[j].region.Start.BanCode {
				if r[i].region.Start.EdaCode == r[j].region.Start.EdaCode {
					if r[i].region.Start.KoedaCode == r[j].region.Start.KoedaCode {
						if r[i].gakuKubun == r[j].gakuKubun {
							return false
						} else if r[i].gakuKubun < r[j].gakuKubun {
							return true
						}
					} else if r[i].region.Start.KoedaCode < r[j].region.Start.KoedaCode {
						return true
					}
				} else if r[i].region.Start.EdaCode < r[j].region.Start.EdaCode {
					return true
				}
			} else if r[i].region.Start.BanCode < r[j].region.Start.BanCode {
				return true
			}
		} else if r[i].region.Start.MachiCode < r[j].region.Start.MachiCode {
			return true
		}
		return false
	})
}

func (r *Records) FillDummyRecords(ymd string) {
	var addRecords Records
	var preRecord Record
	for i, record := range *r {
		if i == 0 { // 最初のレコードの場合
			if re := record.region.StartRegion(); re != nil {
				addRecord := createDummyRecord(record.gakuKubun, record.jichiCode, ymd)
				addRecord.region = *re
				addRecords = append(addRecords, addRecord)
			}
		} else { // 2行目以降のレコードの場合
			if fmr := record.FindMissingRegions(preRecord); fmr != nil {
				for _, region := range fmr {
					addRecord := createDummyRecord(record.gakuKubun, region.Start.JichiCode, ymd)
					re := region
					addRecord.region = re
					addRecords = append(addRecords, addRecord)
				}
			}

			// 最後のレコードの場合
			if i == len(*r)-1 {
				if re := record.region.EndRegion(); re != nil {
					addRecord := createDummyRecord(record.gakuKubun, record.jichiCode, ymd)
					addRecord.region = *re
					addRecords = append(addRecords, addRecord)
				}
			}
		}

		preRecord = record
	}
	*r = append(*r, addRecords...)
}

func createRecords() Records {
	return make([]Record, 0, 3000)
}
