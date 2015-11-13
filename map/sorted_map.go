package main

// sort a map's keys in descending order of its values.

import (
	"fmt"
	"sort"
)

type sortedMap struct {
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]int) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}

func main() {
	maps := make(map[string]int, 10)
	maps["aa"] = 3
	maps["bb"] = 2
	maps["cc"] = 1
	maps["dd"] = 6
	maps["ee"] = 7
	maps["ff"] = 5
	maps["gg"] = 3

	str := sortedKeys(maps)
	fmt.Println(str)
}
