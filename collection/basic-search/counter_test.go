package basicsearch

import (
	"fmt"
	"testing"
)

func TestCounter(t *testing.T) {
	arr := []int{8, 1, 1, 3, 2, 5, 1, 2, 1, 1}
	fmt.Println(Count(arr, 9))
	fmt.Println(Count(arr, 2))
	fmt.Println(Count(arr, 8))
	fmt.Println(Count(arr, 1))
}

func TestFindIndex(t *testing.T) {
	arr := []int{8, 1, 1, 3, 2, 5, 1, 2, 1, 1}
	fmt.Println(FindIndex(arr, 2))
	fmt.Println(FindIndex(arr, 5))
	fmt.Println(FindIndex(arr, 9))
}
