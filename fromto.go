package main

import (
	"fmt"
	"strconv"
)

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "invalid\n"
	}
	result := ""
	step := 1
	if from > to {
		step = -1
	}
	for i := from; ; i += step {
		if i < 10 {
			result += "0" + strconv.Itoa(i)
		} else {
			result += strconv.Itoa(i)
		}
		if i == to {
			break
		}
		result += ", "
	}
	return result + "\n"
}

func main() {
	fmt.Println(FromTo(1, 10))
	fmt.Println(FromTo(10, 1))
	fmt.Println(FromTo(10, 10))
	fmt.Println(FromTo(100, 10))
}
