package solution

import (
	"slices"
)

func GroupAnagrams(strs []string) (answer [][]string) {

	groups := anagramgroups{}
	for _, str := range strs {
		key := getKey(str)
		if hasKey(groups, key) {
			groups[key] = append(groups[key], str)
		} else {
			groups[key] = []string{str}
		}
	}

	for _, group := range groups {
		answer = append(answer, group)
	}

	return
}

func getKey(s string) (k mapkeytype) {
	rs := []rune(s)
	slices.Sort(rs)
	k = mapkeytype(string(rs))
	return
}

func hasKey(g anagramgroups, key mapkeytype) bool {
	_, ok := g[key]
	return ok
}
