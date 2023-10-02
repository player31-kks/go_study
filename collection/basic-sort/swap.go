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

func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIdx] > arr[j] {
				minIdx = j
			}
		}
		Swap(&arr[i], &arr[minIdx])
	}
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		changed := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(&arr[j], &arr[j+1])
				changed = true
			}
		}
		if !changed {
			return
		}
	}
}

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
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
