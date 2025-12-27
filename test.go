package main

import (
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	strFields := strings.Fields(in)
	var snums []int

	for _, str := range strFields {
		intConv, err := strconv.Atoi(str)
		if err == nil {
			snums = append(snums, intConv)
		}
	}

	return "throw towel"
}
