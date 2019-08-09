package utils

import "strings"

func BinSearch(data []int, n int, begin int, end int) bool {
	if begin > end || end > begin {
		return false
	} else {
		var mid int = (begin + end) / 2

		if n == data[mid] {
			return true
		} else if n < data[mid] {
			return BinSearch(data, n, begin, mid - 1)
		} else {
			return BinSearch(data, n, mid + 1, end)
		}
	}
}

func BinSearchString(data []string, n string, begin int, end int) bool {
	if begin > end {
		return false
	} else {
		var mid int = (begin + end) / 2

		if n == data[mid] {
			return true
		} else if strings.Compare(n, data[mid]) < 0 {
			return BinSearchString(data, n, begin, mid - 1)
		} else {
			return BinSearchString(data, n, mid + 1, end)
		}
	}
}
