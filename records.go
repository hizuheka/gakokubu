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
						return false
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

func createRecords() Records {
	return make([]Record, 0, 3000)
}
