package basicsearch

import (
	"fmt"
	"sort"
)

func Compression(input string) {
	m := make(map[string]int)

	for _, value := range input {
		value := string(value)
		m[value] += 1
	}

	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	res := ""
	for _, value := range keys {
		res += fmt.Sprintf("%s%d", value, m[value])
	}
	fmt.Println(res)
}
