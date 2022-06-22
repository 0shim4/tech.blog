package main

import "fmt"

func main() {
	brothers := []string{"taro", "jiro", "taro", "saburo", "saburo", "shiro"}

	unique := make([]string, 0, len(brothers))

	m := make(map[string]struct{})
	for _, v := range brothers {
		if _, ok := m[v]; ok {
			continue
		}
		unique = append(unique, v)
		m[v] = struct{}{}
	}

	fmt.Println(unique)
	// OutPut: [taro jiro saburo shiro]
}
