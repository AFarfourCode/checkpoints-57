package main

import "fmt"

func Slice(a []string, nbrs ...int) []string {
	length := len(a)

	if len(nbrs) == 0 {
		return nil
	}

	start := nbrs[0]
	end := length

	if len(nbrs) > 1 {
		end = nbrs[1]
	}

	// handle negative values
	if start < 0 {
		start = length + start
	}
	if end < 0 {
		end = length + end
	}

	// invalid cases
	if start < 0 || end > length || start >= end {
		return nil
	}

	return a[start:end]
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}

	fmt.Println(Slice(a, 1))
	// ["algorithm" "ascii" "package" "golang"]

	fmt.Println(Slice(a, 2, 4))
	// ["ascii" "package"]

	fmt.Println(Slice(a, -3))
	// ["ascii" "package" "golang"]

	fmt.Println(Slice(a, -2, -1))
	// ["package"]

	fmt.Println(Slice(a, 2, 0))
	// nil
}
