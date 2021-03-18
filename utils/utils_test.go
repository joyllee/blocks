package utils

import (
	"sort"
	"testing"
)

func TestSortCompare(t *testing.T) {
	type Class struct {
		Name         string
		MathSorts    int
		ChineseSorts int
	}

	var cla = []Class{
		{
			Name:         "zhangSan",
			MathSorts:    88,
			ChineseSorts: 90,
		},
		{
			Name:         "LiSi",
			MathSorts:    90,
			ChineseSorts: 70,
		},
		{
			Name:         "WangWu",
			MathSorts:    89,
			ChineseSorts: 96,
		},
	}
	t.Log(cla)
	sort.Slice(cla, func(i, j int) bool {
		return SortCompare(cla[i], cla[j], "MathSorts", "desc")
	})
	t.Log(cla)
}
