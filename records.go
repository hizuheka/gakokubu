package main


import (
        "sort"
)

type Records []Record

func (rs Records) Sort() {
	sort.SliceStable(a, func(i, j int) bool {
		return a[i].CreatedAt.Before(a[j].CreatedAt)
	})
}