package basicsort

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	a, b := 1, 2
	fmt.Printf("a: %d, b: %d\n", a, b)
	Swap(&a, &b)
	fmt.Printf("a: %d, b: %d\n", a, b)
}

func TestSort(t *testing.T) {
	arr := [2]int{9, 3}
	if arr[0] > arr[1] {
		Swap(&arr[0], &arr[1])
	}
	fmt.Println(arr)
}

func TestSort2(t *testing.T) {
	arr := make([][]int, 0, 10)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			element := []int{j, i}
			fmt.Printf("%d %d %v\n", j, i, j <= i)
			if j > i {
				Swap(&element[0], &element[1])
			}
			arr = append(arr, element)
		}
	}
}

func TestSort3(t *testing.T) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				arr := make([]int, 0, 3)
				arr = append(arr, i, j, k)
				fmt.Println("before", CheckSort(arr))
				Sort(arr)
				fmt.Println("after", CheckSort(arr))
			}
		}
	}
}