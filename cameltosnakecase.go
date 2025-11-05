package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if !(ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z') {
			return s
		}

		if i > 0 && ch >= 'A' && ch <= 'Z' && s[i-1] >= 'A' && s[i-1] <= 'Z' {
			return s
		}

		last := s[len(s)-1]
		if last >= 'A' && last <= 'Z' {
			return s
		}
	}
	result := ""
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if i != 0 {
			if ch >= 'A' && ch <= 'Z' {
				result += "_"
			}
		}
		result += string(ch)

	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))       // Hello_World
	fmt.Println(CamelToSnakeCase("helloWorld"))       // hello_World
	fmt.Println(CamelToSnakeCase("camelCase"))        // camel_Case
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE")) // CAMELtoSnackCASE
	fmt.Println(CamelToSnakeCase("camelToSnakeCase")) // camel_To_Snake_Case
	fmt.Println(CamelToSnakeCase("hey2"))             // hey2
}
