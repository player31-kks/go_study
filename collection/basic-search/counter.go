package basicsearch

func Count(arr []int, search int) int {
	count := 0
	for _, value := range arr {
		if value == search {
			count += 1
		}
	}
	return count
}

func FindIndex(arr []int, search int) int {
	for idx, value := range arr {
		if value == search {
			return idx
		}
	}
	return -1
}

// 정렬된 배열에서 Count하기
func SortedCount(arr []int, search int) int {
	idx := FindIndex(arr, search)
	if idx > 0 {
		return sortedCountHelper(arr, idx, search)
	}
	return 0
}

func sortedCountHelper(arr []int, idx int, search int) int {
	count := 0
	for i := idx; i < len(arr); i++ {
		if arr[i] != search {
			return count
		}
		count += 1
	}
	return count
}
