package utils

import "sort"

func SortStrings(data []string) {
	sort.Sort(sort.StringSlice(data))
}

func SortInts(data []int) {
	sort.Sort(sort.IntSlice(data))
}
