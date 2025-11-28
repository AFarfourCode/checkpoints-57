/*
package main

import "fmt"

func CanJump(s []uint) bool {

	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	maxreach := uint(0)

	for i := uint(0); i < uint(len(s)); i++ {
		if i > maxreach {
			return false
		}
		reach := i + s[i]

		if reach > maxreach {
			maxreach = reach
		}
		if maxreach >= uint(len(s)-1) {
			return true
		}
	}
	return true
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1)) // true

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2)) // false

	input3 := []uint{5}
	fmt.Println(CanJump(input3)) // true

	input4 := []uint{7, 1, 1, 1, 1}
	fmt.Println(CanJump(input4)) // true

	input5 := []uint{1, 1, 1, 1, 1}
	fmt.Println(CanJump(input5)) // true

	input6 := []uint{}
	fmt.Println(CanJump(input6)) // false
}
*/

package main

import "fmt"

func CanJump(s []uint) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}

	indx := uint(0)
	for i := len(s); i > 0; i-- {
		steps := s[indx]
		if indx+steps < uint(len(s)-1) {
			indx = indx + steps
		}
		if indx+steps == uint(len(s)-1) {
			return true
		}
		if indx+steps > uint(len(s)-1) {
			break
		}
	}
	return false
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1)) // true

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2)) // false

	input3 := []uint{5}
	fmt.Println(CanJump(input3)) // true

	input4 := []uint{7, 1, 1, 1, 1}
	fmt.Println(CanJump(input4)) // true

	input5 := []uint{1, 1, 1, 1, 1}
	fmt.Println(CanJump(input5)) // true

	input6 := []uint{}
	fmt.Println(CanJump(input6)) // false
}
