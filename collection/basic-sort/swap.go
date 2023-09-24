package basicsort

func Swap(a, b *int) {
	*a, *b = *b, *a
}

// 선택정렬!!!
// 앞으로 나아가는 point의 위치를 잘 파악하자
func Sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				Swap(&arr[i], &arr[j])
			}
		}
	}
}

func CheckSort(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
