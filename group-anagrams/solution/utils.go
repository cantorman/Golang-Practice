package solution

import (
	"slices"
	"sort"
)

func EqualUnordered[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	cnt := make(map[T]int)
	for _, v := range a {
		cnt[v]++
	}
	for _, v := range b {
		if cnt[v]--; cnt[v] < 0 {
			return false
		}
	}
	return true
}

func Values() {
	panic("Not implemented")
}

func OrderMyMap(myValues [][]string) {
	// order each subslice by the strings in it
	for _, list := range myValues {
		slices.Sort(list)
	}
	// next order each item in the unsorted slice of slices by comparing each first string
	sort.Slice(myValues, func(i, j int) bool {
		return myValues[i][0] < myValues[j][0]
	})
}
