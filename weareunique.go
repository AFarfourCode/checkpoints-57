package main

import "fmt"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	map1 := make(map[rune]bool)
	map2 := make(map[rune]bool)

	for _, ch := range str1 {
		map1[ch] = true
	}

	for _, ch := range str2 {
		map2[ch] = true
	}
	unique := make(map[rune]bool)

	for ch := range map1 {
		if !map2[ch] {
			unique[ch] = true
		}
	}

	for ch := range map2 {
		if !map1[ch] {
			unique[ch] = true
		}
	}
	return len(unique)
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
