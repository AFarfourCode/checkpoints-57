package main

import "fmt"

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := ""
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			result += fmt.Sprintf("%d%c", count, s[i-1])
			count = 1
		}
	}
	result += fmt.Sprintf("%d%c", count, s[len(s)-1])
	return result
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
