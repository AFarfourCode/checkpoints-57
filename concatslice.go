package main

import "fmt"

func ConcatSlice(slice1, slice2 []int) []int {
	result := append([]int{}, slice1...)

	result = append(result, slice2...)

	return result
}

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}

/*
package main

import "fmt"

func ConcatSlice(slice1, slice2 []int) []int {
	// نعرف مصفوفة جديدة طولها مجموع طول الـ 2 slice
	result := make([]int, len(slice1)+len(slice2))

	// ننسخ slice1 في النصف الأول
	copy(result, slice1)

	// ننسخ slice2 في النصف الثاني
	copy(result[len(slice1):], slice2)

	return result
}

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))         // [1 2 3 4 5 6]
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))      // [4 5 6 7 8 9]
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))               // [1 2 3]
}
*/
