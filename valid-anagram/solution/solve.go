package solution

import "sort"

type charlist []rune

func toCharList(s string) (cl charlist) {
	for _, chr := range s {
		cl = append(cl, chr)
	}
	return
}

func toString(cl charlist) (s string) {
	for _, c := range cl {
		s = s + string(c)
	}
	return
}

func sortCharList(rList charlist) {
	sort.Slice(rList, func(i, j int) bool {
		return rList[i] < rList[j]
	})
}

func IsAnagram(s, t string) bool {
	sList := toCharList(s)
	tList := toCharList(t)
	sortCharList(sList)
	sortCharList(tList)

	return toString(sList) == toString(tList)
}
